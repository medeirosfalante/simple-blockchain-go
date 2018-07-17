package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

// difficulty miner
const targetBits = 24

// ProofOfWork is a struct for Proof of Work in miner
type ProofOfWork struct {
	block  *Block
	target *big.Int
}

// NewPOW is a func create a new POW
func NewPOW(b *Block) (pow *ProofOfWork) {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))
	pow = &ProofOfWork{b, target}
	return
}

func (p *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			p.block.PrevBlockHash,
			p.block.Data,
			IntToHex(p.block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)
	return data

}

// Execute is a func exec POW
func (p *ProofOfWork) Execute() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0
	maxNonce := math.MaxInt64

	fmt.Printf("MINER the block containing data \"%s\"\n", p.block.Data)
	for nonce < maxNonce {
		data := p.prepareData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(p.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Print("\n\n")
	return nonce, hash[:]
}

// Validate validate block
func (p *ProofOfWork) Validate() (isValid bool) {
	var hashInt big.Int

	data := p.prepareData(p.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid = hashInt.Cmp(p.target) == -1
}
