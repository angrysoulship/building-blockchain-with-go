package main

import "math/big"

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
	targetStr := "0000100000000000000000000000000000000000000000000000000000000000000"
	tmpInt := big.Int{}
	tmpInt.SetString(targetStr, 16)

	pow.target = &tmpInt
	return &pow
}

//find the correct Hash
func (pow *ProofOfWork) Run() ([]byte, uint64) {

	return []byte("hello world"), 10
}
