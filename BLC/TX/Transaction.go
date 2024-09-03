package TX

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"log"
)

// UTXO
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

// Build normal transaction
func NewSimpleTransaction(from string, to string, amount int64) *Transaction {

	var txInputs []*TxInput
	var txOutputs []*TxOutput
	// used
	preTxHash, _ := hex.DecodeString("6d67b257d429418b48b908c5797e39195d73b3c22761ce4cf6b84292e752b135")
	txInput := &TxInput{preTxHash, amount, from}
	txInputs = append(txInputs, txInput)
	// tranfer
	txOutput := &TxOutput{amount, to}
	txOutputs = append(txOutputs, txOutput)
	// retain
	txOutput = &TxOutput{10 - amount, from}
	txOutputs = append(txOutputs, txOutput)

	tx := &Transaction{[]byte{}, txInputs, txOutputs}
	// Set hash value for the tx
	tx.HashTransaction()
	return tx
}
