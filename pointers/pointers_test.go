package main

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		// doesnt actually add to balance
		wallet.BuggyDeposit(Bitcoin(10))

		// adds to balance
		wallet.Deposit(Bitcoin(10))

		// %p: base 16 with leading 0x for address printing
		fmt.Printf("address of wallet in tests: %p\n", &wallet)

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))

		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	if wallet.balance != want {
		// String() is called when using %s formatter
		t.Errorf("got %s, want %s", wallet.balance, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	// making sure we dont call .Error on nil
	if got == nil {
		// stop the test instead of continuing
		t.Fatal("Wanted error, but didn't get one")
	}

	// getting message form error and comparing to want:
	if got != want {
		t.Errorf("Got %q, want %q", got, want)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		t.Errorf("Expected no error, but got: %s", err.Error())
	}
}
