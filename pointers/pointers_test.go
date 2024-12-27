package pointers

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	// doesnt actually add to balance
	wallet.BuggyDeposit(10)

	// adds to balance
	wallet.Deposit(10)

	// %p: base 16 with leading 0x for address printing
	fmt.Printf("address of wallet in tests: %p\n", &wallet)

	got := wallet.Balance()
	want := 10

	assertCorrectResult(t, got, want)
}

func assertCorrectResult(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
