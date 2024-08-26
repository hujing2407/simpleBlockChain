package main

import (
	"hu169.ca/simpleBlockChain/BLC"
)

func main() {
	blChain := BLC.CreateBlockChainWithGenesis()
	defer blChain.DB.Close()

	cli := BLC.CLI{blChain}
	cli.Run()

	//
	//// Add more blocks to chain
	//blChain.AddBlockToChain("Send 100 to A")
	//blChain.AddBlockToChain("Send 200 to B")
	//blChain.AddBlockToChain("Send 300 to C")
	//blChain.AddBlockToChain("Send 400 to D")
	//blChain.AddBlockToChain("Send 500 to E")
	//
	//blChain.PrintChain()

}
