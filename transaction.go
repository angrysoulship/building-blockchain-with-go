package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

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
