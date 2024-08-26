package main

import (
	"hu169.ca/simpleBlockChain/BLC"
)

func main() {
	//blChain := BLC.CreateBlockChainWithGenesis()
	//defer blChain.DB.Close()

	cli := BLC.CLI{}
	cli.Run()
}
