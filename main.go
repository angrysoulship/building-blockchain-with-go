package main

import (
	"fmt"
)

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
