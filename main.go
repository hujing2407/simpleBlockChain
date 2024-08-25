package main

import (
	"fmt"
	"hu169.ca/simpleBlockChain/BLC"
)

func main() {
	//fmt.Println("Hi, Block Chain Project in Golang")
	//blChain := BLC.CreateBlockChainWithGenesis()

	// Add more blocks to chain
	//blChain.AddBlockToChain("Send 100 to A", blChain.Blocks[len(blChain.Blocks)-1].Height+1, blChain.Blocks[len(blChain.Blocks)-1].Hash)
	//blChain.AddBlockToChain("Send 200 to B", blChain.Blocks[len(blChain.Blocks)-1].Height+1, blChain.Blocks[len(blChain.Blocks)-1].Hash)
	//blChain.AddBlockToChain("Send 300 to C", blChain.Blocks[len(blChain.Blocks)-1].Height+1, blChain.Blocks[len(blChain.Blocks)-1].Hash)

	//fmt.Println(blChain)

	block := BLC.NewBlock("Test", 1,
		[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})

	fmt.Printf("%d\n", block.Nonce)
	fmt.Printf("%x\n", block.Hash)

	pow := BLC.NewProofOfWork(block)
	fmt.Printf("%v", pow.IsValid())
}
