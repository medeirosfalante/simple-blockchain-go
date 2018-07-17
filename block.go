package main

import (
	"time"
)

// Block structure
type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

// NewBlock is a func create a new block
func NewBlock(data string, PrevBlockHash []byte) (block *Block) {
	block = &Block{time.Now().Unix(), []byte(data), PrevBlockHash, []byte{}, 0}
	pow := NewPOW(block)
	nonce, hash := pow.Execute()
	block.Hash = hash[:]
	block.Nonce = nonce
	return
}
