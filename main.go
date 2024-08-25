package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"hu169.ca/simpleBlockChain/BLC"
	"log"
)

func main() {
	block := BLC.NewBlock("Test", 1,
		[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})

	fmt.Printf("%d\n", block.Nonce)
	fmt.Printf("%x\n", block.Hash)

	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open("blc.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		// Create table
		//b, err := tx.CreateBucket([]byte("blocks"))
		//if err != nil { // db process Failed
		//	return fmt.Errorf("create bucket failed: %s", err)
		//}
		b := tx.Bucket([]byte("blocks"))
		// Write on table
		if b != nil { // db process Failed
			err := b.Put([]byte("latest"), block.Serialize())
			if err != nil { // db process Failed
				log.Panic("Failed to PUT data into db!")
			}
		}
		// For db process debugging
		return nil
	})

	if err != nil { // db process Failed
		log.Panic(err)
	}

	//fmt.Println("Hi, Block Chain Project in Golang")
	//blChain := BLC.CreateBlockChainWithGenesis()

	// Add more blocks to chain
	//blChain.AddBlockToChain("Send 100 to A", blChain.Blocks[len(blChain.Blocks)-1].Height+1, blChain.Blocks[len(blChain.Blocks)-1].Hash)
	//blChain.AddBlockToChain("Send 200 to B", blChain.Blocks[len(blChain.Blocks)-1].Height+1, blChain.Blocks[len(blChain.Blocks)-1].Hash)
	//blChain.AddBlockToChain("Send 300 to C", blChain.Blocks[len(blChain.Blocks)-1].Height+1, blChain.Blocks[len(blChain.Blocks)-1].Hash)

	//fmt.Println(blChain)

	//
	//bytes := block.Serialize()
	//fmt.Println(bytes)
	//
	//newblock := BLC.Deserialize(bytes)
	//fmt.Printf("%d\n", newblock.Nonce)
	//fmt.Printf("%x\n", newblock.Hash)

}
