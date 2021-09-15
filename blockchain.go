package main

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
