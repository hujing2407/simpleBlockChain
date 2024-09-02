package BLC

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

const targetBits = 16

type ProofOfWork struct {
	Block  *Block   // current block to validate
	target *big.Int // difficulty of mining
}

// Create a new ProofOfWork instance
func NewProofOfWork(block *Block) *ProofOfWork {

	// 1. Initiate the target value (1)
	target := big.NewInt(1)

	// 2. Right move target (256 - targetBits) bits, all hash values which are less than it is Valid.
	target = target.Lsh(target, 256-targetBits)

	return &ProofOfWork{block, target}
}

func (pow *ProofOfWork) preparedata(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevBlockHash,
			pow.Block.HashTransactions(),
			IntToHex(pow.Block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
			IntToHex(int64(pow.Block.Height))},
		[]byte{})
	return data
}

func (proofOfWork *ProofOfWork) Run() ([]byte, int64) {

	nonce := 0
	var hashInt big.Int
	var hash [32]byte

	for {
		// 1. Combine []byte of all attributes
		dataBytes := proofOfWork.preparedata(nonce)

		// 2. Generate hash value
		hash = sha256.Sum256(dataBytes)
		// TODO: remove
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		// 3. Validate the hash value
		if proofOfWork.target.Cmp(&hashInt) == 1 {
			break
		}
		nonce++
	}
	fmt.Println()
	return hash[:], int64(nonce)
}

// Validate the block
func (proofOfWork *ProofOfWork) IsValid() bool {
	var hashInt big.Int
	hashInt.SetBytes(proofOfWork.Block.Hash)
	if proofOfWork.target.Cmp(&hashInt) == 1 {
		return true
	}
	return false
}
