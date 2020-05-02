package account

import (
	"sync"
)

// Account represent a person's bank account
type Account struct {
	amount int64
	open   bool
	mutex  *sync.Mutex
}

// Open creates a new account with some initial money
func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	return &Account{amount: initialDeposit, open: true, mutex: &sync.Mutex{}}
}

// Close shuts down an account by paying out the balance and return a boolean
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

// Balance outputs the current account balance
func (a *Account) Balance() (balance int64, ok bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	if !a.open {
		return 0, false
	}
	return a.amount, true
}

// Deposit allows user to withdraw or deposit money into account
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
