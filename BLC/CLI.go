package BLC

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type CLI struct {
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("\tcreateGenesis -address \t-> Create genesis block")
	fmt.Println("\tsend -from SOURCE -to DEST -amount AMOUNT \t-> TX info")
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
	sendBlockCmd := flag.NewFlagSet("send", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printChain", flag.ExitOnError)
	createGenesisCmd := flag.NewFlagSet("createGenesis", flag.ExitOnError)

	flagFrom := sendBlockCmd.String("from", "", "Source Address")
	flagTo := sendBlockCmd.String("to", "", "Destination Address")
	flagAmount := sendBlockCmd.String("amount", "", "Tx Amount")
	flagCreateGenesisWithAddress := createGenesisCmd.String("address", "", "Create Genesis Block")
	switch os.Args[1] {
	case "createGenesis":
		err := createGenesisCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "send":
		err := sendBlockCmd.Parse(os.Args[2:])
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

	if sendBlockCmd.Parsed() {
		if *flagFrom == "" || *flagTo == "" || *flagAmount == "" {
			printUsage()
			os.Exit(1)
		}
		//cli.addBlock([]*TX.Transaction{})
		//fmt.Println(*flagFrom)
		//fmt.Println(*flagTo)
		//fmt.Println(*flagAmount)

		//fmt.Println(JSONToArray(*flagFrom))
		//fmt.Println(JSONToArray(*flagTo))
		//fmt.Println(JSONToArray(*flagAmount))
		from := JSONToArray(*flagFrom)
		to := JSONToArray(*flagTo)
		amount := JSONToArray(*flagAmount)
		cli.send(from, to, amount)

	}

	if printChainCmd.Parsed() {
		fmt.Println("=== Blocks details are following ===")
		cli.printChain()
	}
}

// Create genesis block
func (cli *CLI) createGenesis(address string) {
	blockchain := CreateBlockChainWithGenesis(address)
	defer blockchain.DB.Close()
}

// Transaction
func (cli *CLI) send(from []string, to []string, amount []string) {
	if DBExisted() == false {
		fmt.Println("Error: Database is not existed!")
		os.Exit(1)
	}
	blockchain := BlockChainObject()
	defer blockchain.DB.Close()
	blockchain.MineNewBlock(from, to, amount)
}

//func (cli *CLI) addBlock(txs []*TX.Transaction) {
//	if DBExisted() == false {
//		fmt.Println("Error: Database is not existed!")
//		os.Exit(1)
//	}
//	blc := BlockChainObject()
//	defer blc.DB.Close()
//	blc.AddBlockToChain(txs)
//}

func (cli *CLI) printChain() {
	if DBExisted() == false {
		fmt.Println("Error: Database is not existed!")
		os.Exit(1)
	}
	blc := BlockChainObject()
	defer blc.DB.Close()
	blc.PrintChain()
}
