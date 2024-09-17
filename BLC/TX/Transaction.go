package TX

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"hu169.ca/simpleBlockChain/BLC"
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

	unSpentTx := BLC.UnSpentTransactionsWithAddr(from)
	fmt.Println(unSpentTx)
	//money, dict :=

	var txInputs []*TxInput
	var txOutputs []*TxOutput
	// used
	//6d67b257d429418b48b908c5797e39195d73b3c22761ce4cf6b84292e752b135
	//159c52f4235d3029c5b0e45ade14c4899e6083936dee936cd76eff9903157dee
	//aa4961b83a967222277f627e660351d3f3620d872259bdd9dc92eb28e90f53cb
	preTxHash, _ := hex.DecodeString("159c52f4235d3029c5b0e45ade14c4899e6083936dee936cd76eff9903157dee")
	txInput := &TxInput{preTxHash, 0, from}
	txInputs = append(txInputs, txInput)
	// tranfer
	txOutput := &TxOutput{amount, to}
	txOutputs = append(txOutputs, txOutput)
	// retain
	txOutput = &TxOutput{4 - amount, from}
	txOutputs = append(txOutputs, txOutput)

	tx := &Transaction{[]byte{}, txInputs, txOutputs}
	// Set hash value for the tx
	tx.HashTransaction()
	return tx
}
