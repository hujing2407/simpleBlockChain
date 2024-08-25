package BLC

type Blockchain struct {
	Blocks []*Block
}

// Create a block chain with Genesis block
func CreateBlockChainWithGenesis() *Blockchain {
	genesisBlock := CreateGenesisBlock("Genesis Data ......")
	return &Blockchain{[]*Block{genesisBlock}}
}

// Add a block to the chain
func (blc *Blockchain) AddBlockToChain(data string, height int64, preHash []byte) {
	newBlock := NewBlock(data, height, preHash)
	blc.Blocks = append(blc.Blocks, newBlock)
}
