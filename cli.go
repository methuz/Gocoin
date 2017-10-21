package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type CLI struct {
	bc *Blockchain
}

func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println(" createblockchain -address ADDRESS - Create a blockchain and send genesis block reward to ADDRESS")
	fmt.Println(" createWalletCmd - Generates a new key-pair and save it into the wallet file")
	fmt.Println(" getbalance  -address ADDRESS - Get balance of ADDRESS")
	fmt.Println(" send  -from FROM -to TO -amount AMOUNT - Send money from FROM to TO for total AMOUNT")
	fmt.Println(" printchain - print all the blocks of the blockchain")
}

func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) Run() {
	cli.validateArgs()

	// Sub
	createBlockchainCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)
	createWalletCmd := flag.NewFlagSet("createwallet", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	getBalanceCmd := flag.NewFlagSet("getbalance", flag.ExitOnError)
	sendCmd := flag.NewFlagSet("send", flag.ExitOnError)

	createBlockchainAddress := createBlockchainCmd.String("address", "", "The address to send genesis block reward to")
	getBalanceAddress := getBalanceCmd.String("address", "", "The address to get balance")

	sendFrom := sendCmd.String("from", "", "The address of receiver")
	sendTo := sendCmd.String("to", "", "The address of sender")
	sendAmount := sendCmd.Int("amount", 0, "Total money to send")

	// Flag

	switch os.Args[1] {
	case "createblockchain":
		err := createBlockchainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "createwallet":
		err := createWalletCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "send":
		err := sendCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "getbalance":
		err := getBalanceCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		cli.printUsage()
		os.Exit(1)
	}

	if createBlockchainCmd.Parsed() {
		if *createBlockchainAddress == "" {
			createBlockchainCmd.Usage()
			os.Exit(1)
		}
		cli.createBlockchain(*createBlockchainAddress)
	}
	if createWalletCmd.Parsed() {
		cli.createWallet()
	}
	if getBalanceCmd.Parsed() {
		if *getBalanceAddress == "" {
			createBlockchainCmd.Usage()
			os.Exit(1)
		}
		cli.getBalance(*getBalanceAddress)
	}
	if sendCmd.Parsed() {
		if *sendFrom == "" || *sendTo == "" || *sendAmount <= 0 {
			sendCmd.Usage()
			os.Exit(1)
		}
		cli.send(*sendFrom, *sendTo, *sendAmount)
	}
	if printChainCmd.Parsed() {
		cli.printChain()
	}
}
