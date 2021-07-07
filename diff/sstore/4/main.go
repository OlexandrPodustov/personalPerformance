package main

import (
	"sync"
	"time"
)

type connetction struct {
	conn  string
	timer *time.Timer
}

type Pool struct {
	sync.Mutex
	storage          map[storageKey]connections
	idleTimeout      time.Duration
	totalConnections chan connetction // substitute with db connections later
}

type storageKey struct {
	auth string
	db   string
}

type connections struct {
	idleConnections []connetction // substitute with db connections later
}

func InitPool(limit int, idleTimeout time.Duration) *Pool {
	p := Pool{
		storage:          make(map[storageKey]connections),
		idleTimeout:      idleTimeout,
		totalConnections: make(chan connetction, limit),
	}

	return &p
}

type result struct{}

func (p *Pool) Query(auth, db, query string) (result, error) {
	// get idle if available. move it to other state
	c := p.GetIdle(auth, db)
	if c != nil {
		defer p.Put(auth, db)
	}

	// if no - open new if total not exceeded
	if c == nil {
		c = p.GetNew()
	}
	// if no - try to close idle to other db, and open here
	// if no available - then wait

	return result{}, nil
}

func (p *Pool) GetIdle(auth, db string) *connetction {
	p.Mutex.Lock()

	defer p.Mutex.Unlock()

	conns := p.storage[storageKey{auth: auth, db: db}]
	// 	if len(conns) == 0 {
	//
	// 	}

	return &conns.idleConnections[0]
}

func (p *Pool) GetNew() *connetction {
	// p.Mutex.Lock()
	// defer p.Mutex.Unlock()

	return nil
}

func (p *Pool) Get() *connetction {
	// p.Mutex.Lock()
	// defer p.Mutex.Unlock()

	return nil
}

func (p *Pool) Put(sauth, db string) *connetction {
	// p.Mutex.Lock()
	// defer p.Mutex.Unlock()

	return nil
}

// CREATE TABLE T (docId int64, doc blob)
// SELECT docId FROM bt WHERE doc CONTAINS (('foo' & 'bar') | 'baz')

// 100 million documents
// 100ms runtime per query
// Words belong to a fixed size dictionary

// Class Op
// {
// func Next() bool
// func Scan() (row, erorr)
// }
//
// Class OpAnd : Op {
// Op op1;
// Op op2;
// }

type OpAnd struct {
	NextRow *OpAnd
}

type row struct{}

func (o *OpAnd) Next() bool {
	if o.NextRow == nil {
		return false
	}

	return true
}

func (o *OpAnd) Scan() (row, error) {

	return row{}, nil
}

func (o *OpAnd) search() (row, error) {
	// for _, v := o.Bar{
	//
	// }
}

// Bar: [1,3,2]
// [1, 2, 3]
// Foo: [1,5,4,2]
// [1, 2, 4, 5]
//
//
// [1,2]
