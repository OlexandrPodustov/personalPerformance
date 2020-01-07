//Package account is about doing things with bank accounts
//Somewhy it hangs on windows env, on linux tests finishes successfully
package account

import "sync"

const testVersion = 1

type Account struct {
	sync.Mutex
	balance int64
	open    bool
}

//Open opens account
func Open(in int64) (res *Account) {
	if in >= 0 {
		res = new(Account)
		res.open = true
		res.balance = in
	}

	return res
}

//Balance retrieves amount on account
func (a *Account) Balance() (balance int64, ok bool) {
	a.Lock()
	defer a.Unlock()
	if a.open {
		balance = a.balance
		ok = a.open
		return
	}

	return
}

//Close closes account
func (a *Account) Close() (payout int64, ok bool) {
	a.Lock()
	defer a.Unlock()

	if a.open {
		ok = a.open
		payout = a.balance
		a.balance = 0
		a.open = false
		return
	}

	return
}

//Deposit deposits account
func (a *Account) Deposit(in int64) (int64, bool) {
	a.Lock()
	defer a.Unlock()
	if !a.open {
		return 0, false
	}
	if a.balance+in < 0 {

		return 0, false
	}
	a.balance += in

	return a.balance, true
}
