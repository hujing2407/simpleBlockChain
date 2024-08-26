package TX

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

type Transaction struct {
	TxHash []byte      // 1. Tx Hash
	Vins   []*TxInput  // 2. input
	Vouts  []*TxOutput // 3. output
}

/* Genesis block transaction */
func NewCoinbaseTransaction(address string) *Transaction {
	// used
	txInput := &TxInput{[]byte{}, -1, "Genesis Data"}
	// total
	txOutput := &TxOutput{10, address}

	txCoinbase := &Transaction{[]byte{}, []*TxInput{txInput}, []*TxOutput{txOutput}}
	// Set hash value for the tx
	txCoinbase.HashTransaction()
	return txCoinbase
}

// Serialize TX to []byte, and then Hash it
func (tx *Transaction) HashTransaction() {

	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)
	err := encoder.Encode(tx)
	if err != nil {
		log.Panic(err)
	}
	hash := sha256.Sum256(res.Bytes())
	tx.TxHash = hash[:]
}
