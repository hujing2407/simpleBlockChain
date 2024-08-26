package BLC

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"hu169.ca/simpleBlockChain/BLC/TX"
	"log"
	"time"
)

type Block struct {
	Height        int64
	PrevBlockHash []byte
	TXs           []*TX.Transaction
	Timestamp     int64
	Hash          []byte //256 bits
	Nonce         int64
}

// Serialize a block to byte array
func (block *Block) Serialize() []byte {

	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)
	err := encoder.Encode(block)
	if err != nil {
		log.Panic(err)
	}
	return res.Bytes()
}

// Deserialize byte array to a block
func Deserialize(blockBytes []byte) *Block {

	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(blockBytes))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}
	return &block
}

func NewBlock(txs []*TX.Transaction, height int64, prevBlockHash []byte) *Block {
	// Create a new block
	block := &Block{height, prevBlockHash, txs, time.Now().Unix(), nil, 0}
	// Replaced the following call with POW progress
	//block.SetHash()

	// Call POW function to get HASH & NONCE
	pow := NewProofOfWork(block)

	// Validate mining work
	hash, nonce := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

func CreateGenesisBlock(txs []*TX.Transaction) *Block {
	log.Println("Creating Genesis Block...")
	return NewBlock(txs, 1,
		[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
}

// Convert TXs to []byte
func (block *Block) HashTransactions() []byte {
	var txHash [32]byte
	var txHashes [][]byte
	for _, tx := range block.TXs {
		txHashes = append(txHashes, tx.TxHash)
	}
	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))
	return txHash[:]
}
