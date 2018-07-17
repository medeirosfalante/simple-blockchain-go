package main

//Blockchain is a struct for blockchain database
type Blockchain struct {
	blocks []*Block
}

//AddBlock is a func add a new block in blockchain
func (b *Blockchain) AddBlock(data string) {
	prevBlock := b.blocks[len(b.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	b.blocks = append(b.blocks, newBlock)
}

// NewGenesisBlock is a func Create a Genesis block
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

// NewBlockchain is a func create a new Blockchain
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
