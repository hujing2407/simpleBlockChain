package main

import (
	"hu169.ca/simpleBlockChain/BLC"
)

func main() {

	blChain := BLC.CreateBlockChainWithGenesis()
	defer blChain.DB.Close()

	// Add more blocks to chain
	blChain.AddBlockToChain("Send 100 to A")
	blChain.AddBlockToChain("Send 200 to B")
	blChain.AddBlockToChain("Send 300 to C")
	blChain.AddBlockToChain("Send 400 to D")
	blChain.AddBlockToChain("Send 500 to E")

	//fmt.Println(blChain)

	//
	//bytes := block.Serialize()
	//fmt.Println(bytes)
	//
	//newblock := BLC.Deserialize(bytes)
	//fmt.Printf("%d\n", newblock.Nonce)
	//fmt.Printf("%x\n", newblock.Hash)

}
