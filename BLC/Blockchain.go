package BLC

import (
	"fmt"
	"github.com/boltdb/bolt"
	"hu169.ca/simpleBlockChain/BLC/TX"
	"log"
	"math/big"
	"os"
	"time"
)

const dbName = "blc.db"
const blockTableName = "blocks"

type Blockchain struct {
	Tip []byte // The hash of the latest block
	DB  *bolt.DB
}

func DBExisted() bool {
	if _, err := os.Stat(dbName); os.IsNotExist(err) {
		return false
	}
	return true
}
func (blc *Blockchain) PrintChain() {

	blcIterator := blc.Iterator()
	for {
		block := blcIterator.Next()
		fmt.Printf("\nBlock Height: %d\n", block.Height)
		fmt.Printf("PrevBlockHash: %x\n", block.PrevBlockHash)

		fmt.Printf("Timestamp: %s\n", time.Unix(block.Timestamp, 0).Format("2006-01-02 03:04:05 PM"))
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Nonce: %d\n", block.Nonce)
		for i, tx := range block.TXs {
			fmt.Printf("TXs[%d]: Hash [%x]\n", i, tx.TxHash)
			for _, in := range tx.Vins {
				fmt.Printf("\tVins Hash:\t%x\n", in.TxHash)
				fmt.Printf("\tVins Vout:\t%d\n", in.Vout)
				fmt.Printf("\tVins ScriptSig:\t%s\n", in.ScriptSig)
			}
			for _, out := range tx.Vouts {
				fmt.Printf("\tVouts Value:\t%d\n", out.Value)
				fmt.Printf("\tVouts Pubkey:\t%s\n", out.ScriptPubkey)
			}
		}
		var hashInt big.Int
		hashInt.SetBytes(block.PrevBlockHash)
		if big.NewInt(0).Cmp(&hashInt) == 0 {
			break
		}
	}
}

/* Create a block chain with Genesis block */
func CreateBlockChainWithGenesis(addr string) {
	if DBExisted() {
		fmt.Println("Failed: Genesis block is existed!")
		os.Exit(1)
	}
	// Open the blc.db data file
	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		//Create table
		b, err := tx.CreateBucket([]byte(blockTableName))
		if err != nil { // db process Failed
			return fmt.Errorf("create bucket failed: %s", err)
		}

		// Write on table
		if b != nil { // db process Failed
			// Create coinbase Tx for Genesis Block
			txCoinbase := TX.NewCoinbaseTransaction(addr)
			genesisBlock := CreateGenesisBlock([]*TX.Transaction{txCoinbase})
			err := b.Put(genesisBlock.Hash, genesisBlock.Serialize())
			if err != nil { // db process Failed
				log.Panicf("Failed to PUT block into db! %s", err)
			}

			err = b.Put([]byte("latest"), genesisBlock.Hash)
			if err != nil { // db process Failed
				log.Panicf("Failed to PUT latest hash into db! %s", err)
			}
		}
		// For db process debugging
		return nil
	})

	if err != nil {
		log.Panic(err)
	}
}

/* Add a block to the chain */
func (blc *Blockchain) AddBlockToChain(txs []*TX.Transaction) {

	err := blc.DB.Update(func(tx *bolt.Tx) error {
		// 1. Get table
		b := tx.Bucket([]byte(blockTableName))

		// 2. Write on table
		if b != nil { // db process

			// Get the current latest block
			lastBlockBytes := b.Get(blc.Tip)
			lastBlock := Deserialize(lastBlockBytes)

			newBlock := NewBlock(txs, lastBlock.Height+1, blc.Tip)
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

func GetBlockChain() *Blockchain {

	// Open the blc.db data file
	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	var tip []byte
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if b != nil {
			tip = b.Get([]byte("latest"))
		}
		return nil
	})
	return &Blockchain{tip, db}
}
