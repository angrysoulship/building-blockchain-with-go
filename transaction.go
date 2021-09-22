package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

const reward = 12.5

type Transaction struct {
	TXID      []byte //交易id
	TXInputs  []TXInput
	TXOutputs []TXOutput
}

type TXInput struct {
	//引用交易id
	TXid []byte
	//引用output index
	Index uint64
	//unlocking script, 用地址模拟
	Sig string
}

type TXOutput struct {
	//金额
	Value float64
	//locking script
	PubKeyHash string
}

//set transaction ID
func (tx *Transaction) SetHash() {
	var buffer bytes.Buffer

	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(&tx)
	if err != nil {
		log.Panic(err)
	}

	data := buffer.Bytes()
	hash := sha256.Sum256(data)
	tx.TXID = hash[:]

}

// create transaction
func NewCoinBaseTX(address string, data string) *Transaction {
	// CoinbaseTX has 1. only one input, 2 no need for import transactionID, 3. no need for import index
	// 挖矿不需要指定签名，data可以由矿工自由填写
	input := TXInput{TXid: []byte{}, Index: 0, Sig: data}
	ouput := TXOutput{Value: reward, PubKeyHash: address}

	tx := Transaction{TXID: []byte{}, TXInputs: []TXInput{input}, TXOutputs: []TXOutput{ouput}}
	tx.SetHash()

	return &tx

}
