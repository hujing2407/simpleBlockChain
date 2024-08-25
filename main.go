package main

import (
	"fmt"
	"hu169.ca/simpleBlockChain/BLC"
)

func main() {
	//fmt.Println("Hi, Block Chain Project in Golang")
	genesisBlc := BLC.CreateBlockChainWithGenesis()
	fmt.Println(genesisBlc)
	fmt.Println(genesisBlc.Blocks[0])
}
