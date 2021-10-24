package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

// TO DO

type ProofOfWork struct {
	// block
	block *Block
	// target nouce bigint是一个非常大数，很多方法可以用
	target *big.Int
}

func NewProofOfWork(block *Block) *ProofOfWork {
	pow := ProofOfWork{
		block: block,
	}

	//difficulty string, and change it to bigint
	targetStr := "000001000000000000000000000000000000000000000000000000000000000000"
	tmpInt := big.Int{}
	tmpInt.SetString(targetStr, 16)

	pow.target = &tmpInt
	return &pow
}

//find the correct Hash
func (pow *ProofOfWork) Run() ([]byte, uint64) {
	var nouce uint64
	block := pow.block
	var hash [32]byte

	// 拼装数据
	for {
		tmp := [][]byte{
			Uint64ToByte(block.Version),
			block.PrevHash,
			block.MerkleRoot,
			Uint64ToByte(block.TimeStamp),
			Uint64ToByte(block.Difficulty),
			Uint64ToByte(nouce),
		}

		blockInfo := bytes.Join(tmp, []byte{})

		// 计算hash
		hash = sha256.Sum256((blockInfo))

		//compare pow.target and hash
		tmpInt := big.Int{}

		tmpInt.SetBytes(hash[:])
		if tmpInt.Cmp(pow.target) == -1 {
			fmt.Println("Mining start...")
			fmt.Printf("Mining Success! Hash: %x, nouce: %d\n", hash, nouce)
			break
		} else {
			nouce++
		}
	}

	return hash[:], nouce
}
