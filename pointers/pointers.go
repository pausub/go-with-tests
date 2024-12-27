package pointers

import "fmt"

type Wallet struct {
	balance int
}

// w is a copied argument to a method
func (w Wallet) BuggyDeposit(amount int) {
	fmt.Printf("address of wallet in Deposit: %p\n", &w)
	w.balance += amount
}

func (w *Wallet) Deposit(amount int) {
	// no need for manual dereference, go handles that
	// automaticaly, so we can call pointer type directly
	w.balance += amount
}

// no need for a pointer here, just convention
func (w *Wallet) Balance() int {
	return w.balance
}
