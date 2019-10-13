// Package account implements simulating a bank account supporting:
// opening/closing, withdrawals, and deposits of money.
package account

// Account is the base account type.
type Account struct {
	balance int64
	active  bool
}

// Open opens account.
func Open(initialDeposit int64) *Account {

	return &Account{initialDeposit, true}
}

// Close close account.
func (a *Account) Close() (payout int64, ok bool) {
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

	return 0, false
}
