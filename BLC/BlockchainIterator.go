package BLC

import (
	"github.com/boltdb/bolt"
	"log"
)

type BlockchainIterator struct {
	CurrentHash []byte
	DB          *bolt.DB
}

func (blc *Blockchain) Iterator() *BlockchainIterator {
	return &BlockchainIterator{blc.Tip, blc.DB}
}

func (blcIterator *BlockchainIterator) Next() *Block {
	var currentBlock *Block
	err := blcIterator.DB.View(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte(blockTableName))
		if b != nil {
			// Get the current latest block
			currentBlockBytes := b.Get(blcIterator.CurrentHash)
			currentBlock = Deserialize(currentBlockBytes)

			// Update the current hash of Iterator
			blcIterator.CurrentHash = currentBlock.PrevBlockHash
		}

		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	return currentBlock
}
