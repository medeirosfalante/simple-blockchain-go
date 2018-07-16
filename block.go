package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// Block structure
type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

// SetHash is a func setter a hash in blockch
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	header := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(header)
	b.Hash = hash[:]
}

// NewBlock is a func create a new block
func NewBlock(data string, PrevBlockHash []byte) (block *Block) {
	block = &Block{time.Now().Unix(), []byte(data), PrevBlockHash, []byte{}}
	block.SetHash()
	return
}
