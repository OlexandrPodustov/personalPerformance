package main

import (
	"sync"

	"github.com/google/uuid"
)

type InMemoryDatabase struct {
	data         map[string]interface{}  // Stores key-value pairs
	transactions map[string]*transaction // Active transactions (key: unique ID)
	lock         *sync.RWMutex           // Ensures thread-safe access
}

type transaction struct {
	changes map[string]interface{} // Changes made within the transaction
	parent  *transaction           // Parent transaction (nil for top-level)
}

func NewInMemoryDatabase() *InMemoryDatabase {
	return &InMemoryDatabase{
		data:         make(map[string]interface{}),
		transactions: make(map[string]*transaction),
		lock:         &sync.RWMutex{},
	}
}

func (db *InMemoryDatabase) get(key string) interface{} {
	db.lock.RLock()
	defer db.lock.RUnlock()

	// Find the latest value considering nested transactions
	currTx := db.currentTransaction()
	for currTx != nil {
		if value, ok := currTx.changes[key]; ok {
			return value
		}
		currTx = currTx.parent
	}

	// Return value from main database if not found in transactions
	if value, ok := db.data[key]; ok {
		return value
	}
	return nil
}

func (db *InMemoryDatabase) set(key string, value interface{}) {
	db.lock.Lock()
	defer db.lock.Unlock()

	// Apply change to current transaction, not directly to main database
	currTx := db.currentTransaction()
	if currTx != nil {
		currTx.changes[key] = value
	} else {
		db.data[key] = value
	}
}

func (db *InMemoryDatabase) delete(key string) {
	db.lock.Lock()
	defer db.lock.Unlock()

	// Mark for deletion in current transaction, avoid direct removal
	currTx := db.currentTransaction()
	if currTx != nil {
		currTx.changes[key] = nil // Use a sentinel value to indicate deletion
	} else {
		delete(db.data, key)
	}
}

func (db *InMemoryDatabase) startTransaction() string {
	db.lock.Lock()
	defer db.lock.Unlock()

	// Generate unique transaction ID
	txId := uuid.New().String()

	// Create a new transaction and link it to the current one if applicable
	parent := db.currentTransaction()
	tx := &transaction{
		changes: make(map[string]interface{}),
		parent:  parent,
	}
	db.transactions[txId] = tx

	return txId
}

func (db *InMemoryDatabase) commit(txId string) {
	db.lock.Lock()
	defer db.lock.Unlock()

	// Ensure transaction exists and is active
	tx, ok := db.transactions[txId]
	if !ok || tx == db.currentTransaction() {
		return // No-op for non-existent or already committed transactions
	}

	// Apply changes from committed transaction to its parent or main database
	if tx.parent != nil {
		tx.parent.changes = mergeMaps(tx.parent.changes, tx.changes)
	} else {
		for key, value := range tx.changes {
			if value == nil { // Sentinel value for deletion
				delete(db.data, key)
			} else {
				db.data[key] = value
			}
		}
	}

	// Remove committed transaction from active transactions
	delete(db.transactions, txId)
}

func (db *InMemoryDatabase) rollback(txId string) {
	db.lock.Lock()
	defer db.lock.Unlock()

	// Ensure transaction exists and is active
	tx, ok := db.transactions[txId]
	if !ok || tx == db.currentTransaction() {
		return // No-op for non-existent or already rolled back transactions
	}

	// Remove transaction changes from its parent or main database (using revertChanges helper)
	if tx.parent != nil {
		revertChanges(tx.parent.changes, tx.changes)
	} else {
		for key := range tx.changes {
			delete(db.data, key) // Directly remove from main database for top-level transactions
		}

		// Remove rolled back transaction from active transactions
		delete(db.transactions, txId)
	}
}

// Helper function to revert changes made within a committed transaction
func revertChanges(parentChanges, txChanges map[string]interface{}) {
	for key, value := range txChanges {
		// If change in transaction was deletion, mark for undo-deletion in parent
		if value == nil {
			parentChanges[key] = true // Use a different sentinel value for undo-deletion
		} else {
			// Otherwise, restore original value from parent if it exists
			if originalValue, ok := parentChanges[key]; ok {
				parentChanges[key] = originalValue
			} else {
				// If no original value in parent, remove the key entirely
				delete(parentChanges, key)
			}
		}
	}
}

// Helper function to merge two maps, handling potential conflicts
func mergeMaps(parent, child map[string]interface{}) map[string]interface{} {
	merged := make(map[string]interface{})
	for key, value := range parent {
		merged[key] = value
	}
	for key, value := range child {
		if existingValue, ok := merged[key]; ok {
			// Handle conflicts (e.g., prefer child's value for updates)
			merged[key] = handleConflict(existingValue, value)
		} else {
			merged[key] = value
		}
	}
	return merged
}

// Placeholder function for handling conflicts during map merging (modify as needed)
func handleConflict(existing, newValue interface{}) interface{} {
	// Example: Prefer the newer value (modify based on your requirements)
	return newValue
}

func (db *InMemoryDatabase) currentTransaction() *transaction {
	db.lock.RLock() // Use Read lock for checking
	defer db.lock.RUnlock()

	// Iterate through the transaction stack to find the active one
	currTx := db.transactions[""] // Start with the top-level transaction (empty key)
	for currTx != nil {
		if _, ok := currTx.changes[""]; ok { // Check if transaction is active (non-empty)
			return currTx
		}
		currTx = currTx.parent
	}

	// No active transaction found
	return nil
}
