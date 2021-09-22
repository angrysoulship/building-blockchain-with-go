package main

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
