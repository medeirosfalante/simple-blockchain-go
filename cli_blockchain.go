package main

import (
	"fmt"
	"log"
)

func (cli *CLI) createBlockchain(address string, nodeID string) {
	if !ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}
	bc := CreateBlockchain(address, nodeID)
	bc.db.Close()

	//UTXOSet := UTXOSet{bc}
	// UTXOSet.Reindex()
	fmt.Println("Done!")
}
