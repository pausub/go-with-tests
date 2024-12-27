package pointers

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		// doesnt actually add to balance
		wallet.BuggyDeposit(Bitcoin(10))

		// adds to balance
		wallet.Deposit(Bitcoin(10))

		// %p: base 16 with leading 0x for address printing
		fmt.Printf("address of wallet in tests: %p\n", &wallet)

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	if wallet.balance != want {
		// String() is called when using %s formatter
		t.Errorf("got %s, want %s", wallet.balance, want)
	}
}
