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

func (block *Block) SetHash() {
	//1. 拼装数据
	blockInfo := append(block.PrevHash, block.Data...)

	//2. SHA256
	hash := sha256.Sum256(blockInfo)
	block.Hash = hash[:]
}

func main() {
	block := NewBlock("Alice has transfered a bitcoin to Bob!!", []byte{})

	fmt.Printf("Previous Block hash: %x\n", block.PrevHash)
	fmt.Printf("Current Block hash: %x\n", block.Hash)
	fmt.Printf("Data: %s\n", block.Data)
	fmt.Println("hello")
}
