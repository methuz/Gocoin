package main

import (
	"fmt"
)

func (cli *CLI) createWallet() {
	wallets, _ := NewWallets()
	address := wallets.CreateWallet()
	wallets.SaveToFile()
	fmt.Printf("Address : %s\n", address)
}
