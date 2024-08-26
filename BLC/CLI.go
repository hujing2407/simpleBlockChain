package BLC

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type CLI struct {
	Blc *Blockchain
}

func (cli *CLI) Run() {
	isValidArgs()
	addBlockCmd := flag.NewFlagSet("addBlock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printChain", flag.ExitOnError)
	createGenesisCmd := flag.NewFlagSet("createGenesis", flag.ExitOnError)

	flagAddBlockData := addBlockCmd.String("data", "github.com", "Tx data")
	createGenesisData := createGenesisCmd.String("data", "Genesis block data...", "Genesis Info")
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
		if *flagAddBlockData == "" {
			fmt.Println("Failed: Tx data is empty!")
			printUsage()
			os.Exit(1)
		}
		cli.createGenesis(*createGenesisData)
	}

	if addBlockCmd.Parsed() {
		if *flagAddBlockData == "" {
			printUsage()
			os.Exit(1)
		}
		cli.addBlock(*flagAddBlockData)
	}

	if printChainCmd.Parsed() {
		fmt.Println("=== Blocks details are following ===")
		cli.printChain()
	}
}
func (cli *CLI) createGenesis(data string) {
	CreateBlockChainWithGenesis(data)
}
func (cli *CLI) addBlock(data string) {
	cli.Blc.AddBlockToChain(data)
}

func (cli *CLI) printChain() {
	cli.Blc.PrintChain()
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("\tcreateGenesis -data \t-> Create genesis block")
	fmt.Println("\taddBlock -data DATA \t-> TX data")
	fmt.Println("\tprintChain \t\t-> Print out blocks details")
}
func isValidArgs() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}
}
