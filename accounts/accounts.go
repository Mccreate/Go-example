package accounts

import (
	"errors"
	"fmt"
)

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

// GetBalance is return Account's balance is not a copy.
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

func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// GetOwner return owner is not a copy.
func (a Account) GetOwner() string {
	return a.owner
}

func (a Account) String() string {
	return fmt.Sprint(a.GetOwner(), "'s account. Has: ", a.GetBalance())
}
