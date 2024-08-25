package BLC

type Blockchain struct {
	Blocks []*Block
}

// Create a block chain with Genesis block
func CreateBlockChainWithGenesis() *Blockchain {
	genesisBlock := CreateGenesisBlock("Genesis Data ......")
	return &Blockchain{[]*Block{genesisBlock}}
}
