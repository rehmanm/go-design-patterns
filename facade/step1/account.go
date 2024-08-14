package main

import "fmt"

type Account struct {
	name string
}

func NewAccount(accountName string) *Account {
	return &Account{
		name: accountName,
	}
}

func (a *Account) CheckAccount(accountName string) error {
	if a.name != accountName {
		return fmt.Errorf("Account Name is Incorrect")
	}
	fmt.Println("Account Verified")
	return nil
}
