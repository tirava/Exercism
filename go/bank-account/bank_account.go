// Package account implements simulating a bank account supporting:
// opening/closing, withdrawals, and deposits of money.
package account

import (
	"sync"
)

// Account is the base account type.
type Account struct {
	active  bool
	balance int64
	mux     *sync.Mutex
}

// Open opens account.
func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	return &Account{true, initialDeposit, &sync.Mutex{}}
}

// Close close account.
func (a *Account) Close() (payout int64, ok bool) {
	a.mux.Lock()
	defer a.mux.Unlock()
	if !a.active {
		return 0, false
	}
	a.active = false
	return a.balance, true
}

// Balance returns account balance.
func (a *Account) Balance() (balance int64, ok bool) {
	if !a.active {
		return 0, false
	}
	return a.balance, true
}

// Deposit deposits of money.
func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	a.mux.Lock()
	defer a.mux.Unlock()
	if !a.active {
		return 0, false
	}
	a.balance += amount
	if a.balance < 0 {
		a.balance = 0
		return a.balance, false
	}
	return a.balance, true
}
