package main

import (
	"crypto/sha256"
	"fmt"
)

//0. 定义结构 Define a block
type Block struct {
	//1. previous Hash
	PrevHash []byte
	//2. current Hash
	Hash []byte
	//3. data
	Data []byte
}

//1. Create a block
func NewBlock(data string, prevBlockhash []byte) *Block {
	block := Block{
		PrevHash: prevBlockhash,
		Hash:     []byte{},
		Data:     []byte(data),
	}

	block.SetHash()

	return &block
}

//2. CreateHash

func (block *Block) SetHash() {
	//1. 拼装数据
	blockInfo := append(block.PrevHash, block.Data...)

	//2. SHA256
	hash := sha256.Sum256(blockInfo)
	block.Hash = hash[:]
}

//3. Create blockchain
type Blockchain struct {
	blocks []*Block
}

//3a. the first block
func GenesisBlock() *Block {
	return NewBlock("GenesisBlock", []byte{})
}

func NewBlockChain() *Blockchain {
	genesisBlock := GenesisBlock()
	return &Blockchain{
		blocks: []*Block{genesisBlock},
	}
}

//4. Add block into blockchain
func (bc *Blockchain) AddBlock(data string) {
	block := NewBlock(data, bc.blocks[len(bc.blocks)-1].Hash)
	bc.blocks = append(bc.blocks, block)
}

func main() {
	bc := NewBlockChain()
	// block := NewBlock("Alice has transfered a bitcoin to Bob!!", []byte{})
	bc.AddBlock("Alice has transfered a bitcoin to Bob!!")
	bc.AddBlock("Alice has transfered a bitcoin to EO!!")

	for i, block := range bc.blocks {
		fmt.Printf("======  Current Blockchain Height: %d  ======= \n", i)
		fmt.Printf("Previous Block hash: %x\n", block.PrevHash)
		fmt.Printf("Current Block hash: %x\n", block.Hash)
		fmt.Printf("Data: %s\n\n", block.Data)
	}

}
