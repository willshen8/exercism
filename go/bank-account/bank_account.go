package account

import (
	"sync"
)

type Account struct {
	amount int64
	open   bool
	mutex  *sync.Mutex
}

func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	return &Account{amount: initialDeposit, open: true, mutex: &sync.Mutex{}}
}

func (a *Account) Close() (payout int64, ok bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	if !a.open {
		return 0, false
	}
	payout = a.amount
	a.amount = 0
	a.open = false
	return payout, true
}

func (a *Account) Balance() (balance int64, ok bool) {
	if !a.open {
		return 0, false
	}
	return a.amount, true
}

func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	if !a.open {
		return 0, false
	}
	if a.amount+amount < 0 {
		return a.amount, false
	}
	a.amount += amount
	return a.amount, true
}
