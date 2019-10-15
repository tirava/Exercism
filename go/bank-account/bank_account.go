// Package account implements simulating a bank account supporting:
// opening/closing, withdrawals, and deposits of money.
package account

import (
	"sync"
)

// Account is the base account type.
type Account struct {
	sync.Mutex
	balance int64
	closed  bool
}

// Open opens account.
func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	return &Account{balance: initialDeposit}
}

// Close close account.
func (a *Account) Close() (payout int64, ok bool) {
	a.Lock()
	defer a.Unlock()
	if a.closed {
		return 0, false
	}
	a.closed = true
	return a.balance, true
}

// Balance returns account balance.
func (a *Account) Balance() (balance int64, ok bool) {
	a.Lock()
	defer a.Unlock()
	if a.closed {
		return 0, false
	}
	return a.balance, true
}

// Deposit deposits of money.
func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	a.Lock()
	defer a.Unlock()
	if a.closed {
		return 0, false
	}
	a.balance += amount
	if a.balance < 0 {
		a.balance = 0
		return a.balance, false
	}
	return a.balance, true
}
