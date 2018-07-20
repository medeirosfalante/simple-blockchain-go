package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
	"time"
)

// Block structure
type Block struct {
	Timestamp     int64
	Transactions  []*Transaction
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

// NewBlock is a func create a new block
func NewBlock(transactions []*Transaction, PrevBlockHash []byte) (block *Block) {
	block = &Block{time.Now().Unix(), transactions, PrevBlockHash, []byte{}, 0}
	pow := NewPOW(block)
	nonce, hash := pow.Execute()
	block.Hash = hash[:]
	block.Nonce = nonce
	return
}

// Serialize is a func serializa
func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}
	return result.Bytes()
}

// DeserializeBlock is a func deserialize a data
func DeserializeBlock(d []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}
	return &block
}

// NewGenesisBlock is a func Create a Genesis block
func NewGenesisBlock(coinbase *Transaction) *Block {
	return NewBlock([]*Transaction{coinbase}, []byte{})
}

// HashTransactions return a hash transaction
func (b *Block) HashTransactions() []byte {
	var txHashes [][]byte
	var txHash [32]byte

	for _, tx := range b.Transactions {
		txHashes = append(txHashes, tx.ID)
	}
	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))
	return txHash[:]
}
