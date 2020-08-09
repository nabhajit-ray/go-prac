package main

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)

}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amt Bitcoin) {

	w.balance = w.balance + amt
	//fmt.Printf("The address of balance in deposit is %v \n", &w.balance)

}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amt Bitcoin) error {
	if w.balance < amt {
		return errors.New("Insufficient fund to withdraw")
	} else {
		w.balance = w.balance - amt
		return nil
	}

}
