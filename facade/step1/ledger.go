package main

import "fmt"

type Ledger struct{}

func (l *Ledger) MakeEntry(accountID, txnType string, amount int) {
	fmt.Printf("Make ledger entry for accountID %s with txnType %s for amount %d\n", accountID, txnType, amount)
}
