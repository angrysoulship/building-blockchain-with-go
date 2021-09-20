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
		fmt.Printf("添加区块\n")
		if len(args) == 4 && args[2] == "--data" {
			data := args[3]
			cli.AddBlock(data)
		} else {
			fmt.Println("wrong input, please check again.")
			fmt.Println(Usage)
		}
	case "printChain":
		fmt.Printf("打印区块\n")
		cli.PrintBlockChain()
	default:
		fmt.Printf("pleasecheckagian\n")
		fmt.Println(Usage)
	}

	// 3. execute the command

}
