package main

import "fmt"

type Wallet struct {
	balance int
}

func NewWallet() *Wallet {
	return &Wallet{balance: 0}
}

func (w *Wallet) CreditBalance(amount int) {
	w.balance += amount
	fmt.Println("Amount Credited to Wallet")
}

func (w *Wallet) DebitBalance(amount int) error {
	if w.balance < amount {
		return fmt.Errorf("not enough money")
	}
	fmt.Println("Amount Debited from Wallet")
	w.balance -= amount
	return nil
}
