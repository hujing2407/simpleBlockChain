package BLC

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

const dbName = "blc.db"
const blockTableName = "blocks"

type Blockchain struct {
	Tip []byte // The hash of the latest block
	DB  *bolt.DB
}

// Create a block chain with Genesis block
func CreateBlockChainWithGenesis() *Blockchain {
	// Open the blc.db data file
	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	var blockHash []byte
	err = db.Update(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte(blockTableName))
		if b == nil {
			//If table is not existed, Create table
			b, err = tx.CreateBucket([]byte(blockTableName))
			if err != nil { // db process Failed
				return fmt.Errorf("create bucket failed: %s", err)
			}
		}

		// Write on table
		if b != nil { // db process Failed
			genesisBlock := CreateGenesisBlock("Genesis Data ......")
			err := b.Put(genesisBlock.Hash, genesisBlock.Serialize())
			if err != nil { // db process Failed
				log.Panicf("Failed to PUT block into db! %s", err)
			}

			err = b.Put([]byte("latest"), genesisBlock.Hash)
			if err != nil { // db process Failed
				log.Panicf("Failed to PUT latest hash into db! %s", err)
			}
			blockHash = genesisBlock.Hash
		}
		// For db process debugging
		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	return &Blockchain{blockHash, db}
}

// Add a block to the chain
func (blc *Blockchain) AddBlockToChain(data string) {

	err := blc.DB.Update(func(tx *bolt.Tx) error {
		// 1. Get table
		b := tx.Bucket([]byte(blockTableName))

		// 2. Write on table
		if b != nil { // db process

			// Get the current latest block
			lastBlockBytes := b.Get(blc.Tip)
			lastBlock := Deserialize(lastBlockBytes)

			newBlock := NewBlock(data, lastBlock.Height+1, blc.Tip)
			err := b.Put(newBlock.Hash, newBlock.Serialize())
			if err != nil { // db process Failed
				log.Panicf("Failed to PUT block into db! %s", err)
			}

			err = b.Put([]byte("latest"), newBlock.Hash)
			if err != nil { // db process Failed
				log.Panicf("Failed to PUT latest hash into db! %s", err)
			}
			// 3. Update latest hash (Tip)
			blc.Tip = newBlock.Hash
		} else {
			log.Panic("Get bucket failed.")
		}
		// For db process debugging
		return nil
	})

	if err != nil {
		log.Panic(err)
	}
}
