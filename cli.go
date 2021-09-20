package main

import (
	"fmt"
	"os"
)

// for receiving command from cli

type CLI struct {
	bc *Blockchain
}

const Usage = `
	addBlock --data DATA            "add data to blockchain"
	printChain                      "print blockchain data"
`

//接受参数的动作

func (cli *CLI) Run() {

	// ./block printChain
	// ./block addBlock

	// 1. get all commands from cli
	args := os.Args
	// 校验
	if len(args) < 2 {
		fmt.Println(Usage)
		return
	}

	// 2. analysis the command
	cmd := args[1]
	switch cmd {
	case "addBlock":
		fmt.Println("添加区块")
	case "printChain":
		fmt.Println("打印区块")
	default:
		fmt.Println("pleasecheckagian")
		fmt.Println(Usage)
	}

	// 3. execute the command

}
