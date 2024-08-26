package BLC

import (
	"flag"
	"fmt"
	"hu169.ca/simpleBlockChain/BLC/TX"
	"log"
	"os"
)

type CLI struct {
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("\tcreateGenesis -address \t-> Create genesis block")
	fmt.Println("\taddBlock -data DATA \t-> TX data")
	fmt.Println("\tprintChain \t\t-> Print out blocks details")
}
func isValidArgs() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) Run() {
	isValidArgs()
	addBlockCmd := flag.NewFlagSet("addBlock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printChain", flag.ExitOnError)
	createGenesisCmd := flag.NewFlagSet("createGenesis", flag.ExitOnError)

	flagAddBlockData := addBlockCmd.String("data", "github.com", "Tx data")
	flagCreateGenesisWithAddress := createGenesisCmd.String("address", "", "Create Genesis Block")
	switch os.Args[1] {
	case "createGenesis":
		err := createGenesisCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "addBlock":
		err := addBlockCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "printChain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}

	default:
		printUsage()
		os.Exit(1)
	}

	if createGenesisCmd.Parsed() {
		if *flagCreateGenesisWithAddress == "" {
			fmt.Println("Failed: Address is empty!")
			printUsage()
			os.Exit(1)
		}
		cli.createGenesis(*flagCreateGenesisWithAddress)
	}

	if addBlockCmd.Parsed() {
		if *flagAddBlockData == "" {
			printUsage()
			os.Exit(1)
		}
		cli.addBlock([]*TX.Transaction{})
	}

	if printChainCmd.Parsed() {
		fmt.Println("=== Blocks details are following ===")
		cli.printChain()
	}
}
func (cli *CLI) createGenesis(address string) {
	CreateBlockChainWithGenesis(address)
}
func (cli *CLI) addBlock(txs []*TX.Transaction) {
	if DBExisted() == false {
		fmt.Println("Error: Database is not existed!")
		os.Exit(1)
	}
	blc := GetBlockChain()
	defer blc.DB.Close()
	blc.AddBlockToChain(txs)
}

func (cli *CLI) printChain() {
	if DBExisted() == false {
		fmt.Println("Error: Database is not existed!")
		os.Exit(1)
	}
	blc := GetBlockChain()
	defer blc.DB.Close()
	blc.PrintChain()
}
