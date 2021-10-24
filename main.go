package main

// import (
// 	"fmt"
// )

func main() {
	bc := NewBlockChain("Alice")
	cli := CLI{bc: bc}
	cli.Run()

	// withou CLI
	// block := NewBlock("Alice has transfered a bitcoin to Bob!!", []byte{})
	// bc.AddBlock("1111111111!!")
	// bc.AddBlock("2222222222!!")

	// //create iterator
	// it := bc.NewIterator()

	// //using iterator and return each block
	// for {
	// 	// having the last block and shift to the left block
	// 	block := it.Next()

	// 	fmt.Printf("=============================================\n\n")
	// 	fmt.Printf("Previous Block hash: %x\n", block.PrevHash)
	// 	fmt.Printf("Current Block hash: %x\n", block.Hash)
	// 	fmt.Printf("Nouce: %x\n", block.Nouce)
	// 	fmt.Printf("Difficulty: %x\n", block.Difficulty)
	// 	fmt.Printf("Timestamp: %x\n", block.TimeStamp)
	// 	fmt.Printf("Data: %s\n", block.Data)

	// 	if len(block.PrevHash) == 0 {
	// 		fmt.Println("iteration ends here")
	// 		break
	// 	}

	// }

}
