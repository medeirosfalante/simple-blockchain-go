package main

import (
	"fmt"
)

func main() {
	bc := NewBlockchain()

	bc.AddBlock("Send 1 BTC to RAFAEL")
	bc.AddBlock("Send 2 more MARC")

	for _, block := range bc.blocks {
		fmt.Printf("tutorial Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("tutorial Data: %s\n", block.Data)
		fmt.Printf(" tutorial Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
