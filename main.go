package main

import (
	"fmt"
	"unit-testing/bank"
)

func main() {
	account := bank.Account{}

	// Deposit money
	if err := account.Deposit(100.50); err != nil {
		fmt.Println("Error:", err)
	}

	// Attempt to withdraw money
	if err := account.Withdraw(50); err != nil {
		fmt.Println("Error:", err)
	}

	// Attempt to withdraw more than the balance
	if err := account.Withdraw(100); err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Final Balance:", account.GetBalance())
}
