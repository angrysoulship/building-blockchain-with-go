package main

import (
	"log"

	"github.com/boltdb/bolt"
)

//3. Create blockchain
type Blockchain struct {
	// blocks []*Block
	db *bolt.DB

	tail []byte //储存最后的一个区块的hash
}

const blockChainDb = "blockChain.db"

func NewBlockChain() *Blockchain {

	// return &Blockchain{
	// 	blocks: []*Block{genesisBlock},
	// }
	var lastHash []byte

	//1. 打开数据库
	db, err := bolt.Open(blockChainDb, 0600, nil)
	defer db.Close()

	if err != nil {
		log.Panic("open failed")
	}

	//改写
	db.Update(func(tx *bolt.Tx) error {
		//找到bucket，没有就创建
		bucket := tx.Bucket([]byte("b1"))
		if bucket == nil {
			bucket, err = tx.CreateBucket([]byte("b1"))
			if err != nil {
				log.Panic("创建bucket失败")
			}

			genesisBlock := GenesisBlock()

			//3. 写数据 hash作为key，block字节流作为value
			bucket.Put(genesisBlock.Hash, genesisBlock.toByte())
			bucket.Put([]byte("LastHashKey"), genesisBlock.Hash)
			lastHash = genesisBlock.Hash
		} else {
			lastHash = bucket.Get([]byte("LastHashKey"))
		}

		return nil
	})

	return &Blockchain{db, lastHash}

}

//3a. the first block
func GenesisBlock() *Block {
	return NewBlock("GenesisBlock", []byte{})
}

func (bc *Blockchain) AddBlock(data string) {
	// block := NewBlock(data, bc.blocks[len(bc.blocks)-1].Hash)
	// bc.blocks = append(bc.blocks, block)
}
