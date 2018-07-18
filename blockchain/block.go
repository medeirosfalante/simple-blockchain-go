package blockchain

import (
	"bytes"
	"encoding/gob"
	"log"
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
