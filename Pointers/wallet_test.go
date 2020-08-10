package main

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(t *testing.T, got error, want string) {
		t.Helper()
		if got == nil {

			t.Fatal("wanted an error but didn't get one")
		}
		if got.Error() != want {
			t.Errorf("got %q want %q", got, want)

		}

	}

	t.Run("Testing the wallet deposit", func(t *testing.T) {

		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		got := wallet.Balance()
		want := Bitcoin(10)
		fmt.Printf("The balance address in test is %v \n", &wallet.balance)
		if got != want {
			t.Errorf("got %s and wanted %s", got, want)
		}
	})
	t.Run("Testing the wallet withdrawl", func(t *testing.T) {

		wallet := Wallet{balance: 30}
		wallet.Withdraw(Bitcoin(10))
		got := wallet.Balance()
		want := Bitcoin(20)
		//fmt.Printf("The balance address in test is %v \n", &wallet.balance)
		if got != want {
			assertBalance(t, wallet, want)
		}
	})
	t.Run("Testing the wallet withdrawl more than balance", func(t *testing.T) {

		wallet := Wallet{balance: 30}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, "Insufficient fund to withdraw")

	})

}
