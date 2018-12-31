package errors

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, wallet Wallet, want Dollar) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Dollar(10))
		assertBalance(t, wallet, Dollar(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Dollar(20)}
		wallet.Withdraw(Dollar(10))
		assertBalance(t, wallet, Dollar(10))
	})
}
