package main

import (
	"fmt"
)

func (cli *CLI) reindexUTXO() {
	bc := NewBlockchain()
	UTXOSet := UTXOSet{bc}
	UTXOSet.Reindex()
	count := UTXOSet.CountTransactions()
	fmt.Printf(" done! %d transactions", count)
}
