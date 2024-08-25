package main

import (
	"fmt"
	"hu169.ca/simpleBlockChain/BLC"
)

func main() {
	fmt.Println("Hi, Block Chain Project in Golang")
	block := BLC.NewBlock("Genesis Block", 1,
		[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	fmt.Println(block)
}
