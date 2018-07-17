package main

import "math/big"

// difficulty miner
const targetBits = 24

// ProofOfWork is a struct for Proof of Work in miner
type ProofOfWork struct {
	block  *Block
	target *big.Int
}
