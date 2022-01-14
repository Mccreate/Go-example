package accounts

import "errors"

// Account struct
type Account struct {
	owner   string
	balance int
}

// error type var must be started name with err000.
var errNoMoney = errors.New("can't withdraw.")

// NewAccount with *
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit with *
func (a *Account) Deposit(amount int) {
	a.balance += amount

}

// GetBalance is return Account's balance.
func (a Account) GetBalance() int {
	return a.balance
}

// Withdraw can be occurred, Account's balance is enough.
// Unless Account's balance is enough, return error type.
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		// It's error control.
		return errNoMoney
	}
	a.balance -= amount
	// if not error occurred, return nil.
	return nil
}
