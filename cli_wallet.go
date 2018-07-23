package main

import (
	"fmt"
	"log"
)

func (cli *CLI) createWallet() {
	wallets, _ := NewWallets()
	address := wallets.CreateWallet()
	wallets.SaveToFile()
	fmt.Printf("you address %s\n", address)
}

func (cli *CLI) getBalance(address string) {
	if !ValidateAddress(address) {
		log.Panic("ERROR: Address is a not valid")
	}
	bc := NewBlockchain(address)
	defer bc.db.Close()

	balance := 0
	pubKeyhash := Base58Decode([]byte(address))
	pubKeyhash = pubKeyhash[1 : len(pubKeyhash)-4]
	UTXOs := bc.FindUTXO(pubKeyhash)

	for _, out := range UTXOs {
		balance += out.Value
	}

	fmt.Printf("Balance '%s' : %d\n", address, balance)

}

func (cli *CLI) listAddresses() {
	wallets, err := NewWallets()
	if err != nil {
		log.Panic(err)
	}
	addresses := wallets.GetAddresses()

	for _, address := range addresses {
		fmt.Println(address)
	}
}
