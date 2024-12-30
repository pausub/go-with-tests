package main

import (
	"errors"
	"fmt"
)

// by declaring type from int we can add specific methods
// to Bitcoin, such as String()
type Bitcoin int

type Wallet struct {

	// lower case letter: package private
	balance Bitcoin
}

var ErrInsufficientFunds = errors.New("insuficient funds in balance")

// w is a copied argument to a method
func (w Wallet) BuggyDeposit(amount Bitcoin) {
	fmt.Printf("address of wallet in Deposit: %p\n", &w)
	w.balance += amount
}

func (w *Wallet) Deposit(amount Bitcoin) {
	// no need for manual dereference, go handles that
	// automaticaly, so we can call pointer type directly
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if w.balance < amount {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

// no need for a pointer here, just convention
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// implement Stringer interface on Bitcoin type
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
