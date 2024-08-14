package main

import "fmt"

func main() {
	fmt.Println("ABC Bank")

	walletFacade := NewWalletFacade("abc", 1234)

	fmt.Println()

	err := walletFacade.AddMoneyToWallet("abc", 1234, 1000)

	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

	fmt.Println()

	err = walletFacade.DeductMoneyFromWallet("abc", 1234, 10000)

	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

}
