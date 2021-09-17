package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"log"
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

	// block.SetHash()
	// 创建pow对象
	pow := NewProofOfWork(&block)
	// 计算查找目标随机数
	hash, nouce := pow.Run()

	// 根据挖矿结果进行对区块数据进行更新
	block.Hash = hash
	block.Nouce = nouce

	return &block
}

func (block *Block) toByte() []byte {
	return []byte{}
}

// uint64转[]byte辅助函数
func Uint64ToByte(num uint64) []byte {
	var buffer bytes.Buffer

	err := binary.Write(&buffer, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buffer.Bytes()
}

//2. CreateHash

func (block *Block) SetHash() {
	// 1. 拼装数据
	// var blockInfo []byte
	// blockInfo = append(blockInfo, byte(block.Version))
	// blockInfo = append(blockInfo, block.PrevHash...)
	// blockInfo = append(blockInfo, block.MerkleRoot...)
	// blockInfo = append(blockInfo, byte(block.TimeStamp))
	// blockInfo = append(blockInfo, byte(block.Difficulty))
	// blockInfo = append(blockInfo, byte(block.Nouce))
	// blockInfo = append(blockInfo, block.Data...)

	tmp := [][]byte{
		Uint64ToByte(block.Version),
		block.PrevHash,
		block.MerkleRoot,
		Uint64ToByte(block.TimeStamp),
		Uint64ToByte(block.Difficulty),
		Uint64ToByte(block.Nouce),
		block.Data,
	}

	blockInfo := bytes.Join(tmp, []byte{})

	//2. SHA256
	hash := sha256.Sum256(blockInfo)
	block.Hash = hash[:]
}
