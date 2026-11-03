// Solution to the episode 8 exercise. Run it with: go run ./solution
package main

import (
	"errors"
	"fmt"
)

// Account is a bank account. balance is lowercase, so only this package can change it.
type Account struct {
	Owner   string
	balance float64
}

// Deposit adds money. It changes the account, so it needs a pointer receiver.
func (a *Account) Deposit(amount float64) {
	a.balance += amount
}

// Withdraw takes money out, but never lets the balance go below zero.
func (a *Account) Withdraw(amount float64) error {
	if amount > a.balance {
		return errors.New("insufficient funds")
	}
	a.balance -= amount
	return nil
}

// Balance only reads the account.
func (a *Account) Balance() float64 {
	return a.balance
}

func main() {
	acc := Account{Owner: "Alex"}
	acc.Deposit(100)
	if err := acc.Withdraw(30); err != nil {
		fmt.Println("error:", err)
	}
	if err := acc.Withdraw(500); err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%s has %.2f\n", acc.Owner, acc.Balance())
}
