package main

import (
	"log"

	"github.com/boltdb/bolt"
)

type BlockChainIterator struct {
	db                 *bolt.DB
	currentHashPointer []byte
}

func (bc *Blockchain) NewIterator() *BlockChainIterator {
	return &BlockChainIterator{
		db: bc.db,
		// 最初指向区块链的最后一个区块，随着next调用，不断变化
		currentHashPointer: bc.tail,
	}

}

func (it *BlockChainIterator) Next() *Block {
	var block Block
	it.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			log.Panic("没找到bucket")
		}

		blockTmp := bucket.Get(it.currentHashPointer)

		//解码
		block = Deserialize(blockTmp)

		//hash左移
		it.currentHashPointer = block.PrevHash

		return nil
	})

	return &block
}
