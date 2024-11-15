package bank

import (
	"errors"
)

type Account struct {
	Balance float64
}

// Deposit
func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("deposit amount must be greater than zero")
	}
	a.Balance += amount
	return nil
}

// Withdraw
func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("withdrawal amount must be greater than zero")
	} else if amount > a.Balance {
		return errors.New("insufficient funds")
	}
	a.Balance -= amount
	return nil
}

// GetBalance
func (a *Account) GetBalance() float64 {
	return a.Balance
}
