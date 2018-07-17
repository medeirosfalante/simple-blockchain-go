package main

import "math/big"

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
