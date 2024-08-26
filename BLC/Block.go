package BLC

import (
	"bytes"
	"encoding/gob"
	"log"
	"time"
)

type Block struct {
	Height        int64
	PrevBlockHash []byte
	Data          []byte
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

func NewBlock(data string, height int64, prevBlockHash []byte) *Block {
	// Create a new block
	block := &Block{height, prevBlockHash, []byte(data), time.Now().Unix(), nil, 0}
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

func CreateGenesisBlock(data string) *Block {
	log.Println("Creating Genesis Block...")
	return NewBlock(data, 1,
		[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
}

/* Replaced the SetHash() to POW process
func (block *Block) SetHash() {
	// Convert Height,Timestamp to []byte
	heightBytes := IntToHex(block.Height)
	timeString := strconv.FormatInt(block.Timestamp, 2)
	timeBytes := []byte(timeString)

	// Combine all attributes
	blockBytes := bytes.Join([][]byte{
		heightBytes, block.PrevBlockHash, block.Data, timeBytes, block.Hash}, []byte{})
	// Generate Hash value
	hash := sha256.Sum256(blockBytes)
	block.Hash = hash[:]
}*/
