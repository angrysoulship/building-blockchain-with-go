package main

import "fmt"

func (cli *CLI) AddBlock(data string) {
	// cli.bc.AddBlock(txs)
	fmt.Printf("Add Block Successful\n")
}

func (cli *CLI) PrintBlockChain() {

	bc := cli.bc
	it := bc.NewIterator()

	for {

		block := it.Next()

		fmt.Printf("=============================================\n\n")
		fmt.Printf("Version: %x\n", block.Version)
		fmt.Printf("Previous Block hash: %x\n", block.PrevHash)
		fmt.Printf("Merkle tree: %x\n", block.MerkleRoot)
		fmt.Printf("Current Block hash: %x\n", block.Hash)
		fmt.Printf("Nouce: %x\n", block.Nouce)
		fmt.Printf("Difficulty: %x\n", block.Difficulty)
		fmt.Printf("Timestamp: %x\n", block.TimeStamp)
		fmt.Printf("CoinbaseTX data: %s\n", block.Transaction[0].TXInputs[0].Sig)

		if len(block.PrevHash) == 0 {
			fmt.Println("iteration ends here")
			break
		}

	}
}
