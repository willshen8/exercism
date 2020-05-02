package account

import (
	"sync"
)

// Account represent a person's bank account
type Account struct {
	sync.RWMutex
	balance int64
	close   bool
}

// close creates a new account with some initial money
func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	return &Account{balance: initialDeposit}
}

// Close shuts down an account by paying out the balance and return a boolean
func (a *Account) Close() (payout int64, ok bool) {
	a.Lock()
	defer a.Unlock()
	if a.close {
		return 0, false
	}
	payout = a.balance
	a.balance = 0
	a.close = true
	return payout, true
}

// Balance outputs the current account balance
func (a *Account) Balance() (balance int64, ok bool) {
	a.Lock()
	defer a.Unlock()
	if a.close {
		return 0, false
	}
	return a.balance, true
}

// Deposit allows user to withdraw or deposit money into account
func (a *Account) Deposit(balance int64) (newBalance int64, ok bool) {
	a.Lock()
	defer a.Unlock()
	if a.close {
		return 0, false
	}
	if a.balance+balance < 0 {
		return a.balance, false
	}
	a.balance += balance
	return a.balance, true
}
