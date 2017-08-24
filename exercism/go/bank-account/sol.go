package account

import "sync"

const testVersion = 1

type Account struct {
	sync.Mutex
	balance int64
	ok      bool
}

func Open(in int64) (res *Account) {
	if in >= 0 {
		res = new(Account)
		res.ok = true
		res.balance = in
	}
	return res
}

func (a *Account) Balance() (balance int64, ok bool) {
	a.Lock()
	defer a.Unlock()
	if a.ok {
		balance = a.balance
		ok = a.ok
		return
	}
	//a.Unlock()
	return
}

func (a *Account) Close() (payout int64, ok bool) {
	a.Lock()
	defer a.Unlock()

	if a.ok {
		ok = a.ok
		payout = a.balance
		a.balance = 0
		a.ok = false
		return
	}
	//a.Unlock()
	return
}

func (a *Account) Deposit(in int64) (newBalance int64, ok bool) {
	a.Lock()
	defer a.Unlock()
	if a.ok && a.balance+in > 0 {
		a.balance = a.balance + in
		newBalance = a.balance
		ok = a.ok
		return
	}
	//a.Unlock()
	return
}
