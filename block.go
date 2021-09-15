package main

import (
	"crypto/sha256"
	"time"
)

//0. 定义结构 Define a block
type Block struct {
	//version
	Version uint64
	//1. previous Hash
	PrevHash []byte
	//2. Merkle tree
	MerkleRoot []byte
	//4. 时间
	TimeStamp uint64
	//5. 难度值
	Difficulty uint64
	//6. 随机数
	Nouce uint64

	//a. current Hash, 正常比特币没有当前区块hash，这么做是为了简化
	Hash []byte
	//b. data
	Data []byte
}

//实现辅助函数，uint64转[]byte

func Uint64ToByte(num uint64) []byte {
	return []byte{}
}

//1. Create a block
func NewBlock(data string, prevBlockhash []byte) *Block {
	block := Block{
		Version:    00,
		PrevHash:   prevBlockhash,
		MerkleRoot: []byte{},
		TimeStamp:  uint64(time.Now().Unix()),
		Difficulty: 0, // 无效
		Nouce:      0, //无效
		Hash:       []byte{},
		Data:       []byte(data),
	}

	block.SetHash()

	return &block
}

//2. CreateHash

func (block *Block) SetHash() {
	//1. 拼装数据
	var blockInfo []byte
	blockInfo = append(blockInfo, byte(block.Version))
	blockInfo = append(blockInfo, block.PrevHash...)
	blockInfo = append(blockInfo, block.MerkleRoot...)
	blockInfo = append(blockInfo, byte(block.TimeStamp))
	blockInfo = append(blockInfo, byte(block.Difficulty))
	blockInfo = append(blockInfo, byte(block.Nouce))
	blockInfo = append(blockInfo, block.Data...)

	//2. SHA256
	hash := sha256.Sum256(blockInfo)
	block.Hash = hash[:]
}

//3. 补充区块字段
//4. 更新计算哈希函数
//5. 优化代码
