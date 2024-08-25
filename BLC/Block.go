package BLC

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
	"strconv"
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
	err := decoder.Decode(block)
	if err != nil {
		log.Panic(err)
	}
	return &block
}

func NewBlock(data string, height int64, prevBlockHash []byte) *Block {
	// Create a new block
	block := &Block{height, prevBlockHash, []byte(data), time.Now().Unix(), nil, 0}
	block.SetHash()

	// Call POW function to get HASH & NONCE
	//pow := NewProofOfWork(block)

	// Validate mining work
	//hash, nonce := pow.Run()

	return block
}

func CreateGenesisBlock(data string) *Block {
	return NewBlock(data, 1,
		[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
}

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

	// TODO: remove the println()
	//fmt.Println(heightBytes)
	//fmt.Println(timeString)
	//fmt.Println(timeBytes)

}
