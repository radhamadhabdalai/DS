/*
 * Copyright (c) 2020 Radhamadhab Dalai
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE.
 */


package BCHAIN

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
)

// Print USage for main block with transactions
func (cli \*CommandLine) PrintUsage(){

	fmt.Println("Usage:")
	fmt.Println("getbalance -address ADDRESS - get balance for ADDRESS")
	fmt.Println("createblockchain -address ADDRESS creates a blockchain and reward")
    fmt.Println("printchain - Prints the block in the chain")
	fmt.Println("send - from FROM -to TO -amount AMOUNT - Send amount of coins from A to B")
}


func (cli *CommandLine) createBlockChain(address string) {

	newChain := blockchain.InitBLockChain(address)
	newChain.Database.Close()
	fmt.Println("Finished creating Blockchain")
}

// get balance routine


func(cli *CommandLine) getBalance(address string){

	chain := blockchain.ContinueBlockCHain(address)
	defer chain.Database.Close()

	balance := 0

	UTXOs := chain.FindUTXO(address)
	for _, out := range UTXOs {

		balance += out.Value

	}

	fmt.Printf("Balance of %s : %d \n", address, balance)
}

// send money routine

func (cli *CommandLine) send (from , to string , amount int) {
	chain := blockchain.ContinueBlockChain(from)

	defer chain.Database.Close()

	tx := blockchain.NewTransaction(from , to , amount , chain)
	chain.AddBlock([]*Blockchain.Transaction{tx})
	fmt.Println("Success")


}

// run command

func (cli *CommandLine) run(){
	cli.validateArgs()

	getBalanceCmd := flag.NewFlagSet("getBalance", flag.ExitOnError)
	createBlockChainCmd := flag.NewFlagSet("createBlockChain", flag.ExitOnError)
    sendCmd := flag.NewFlagSet("sendMoney", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printChain", flag.ExitOnError)
	createWalletCmd := flag.NewFlagSet("createwallet", flag.ExitOnError)
	listAddressCmd := flag.NewFlagSet("listaddress", flag.ExitOnError)

	getBalanceAddress := getBalanceCmd.String("address", "", "Addres")
	createBlockChainAddress := createBlockChainCmd.String("Blockchain Address","","Create Block chain Address")
	sendFrom := sendCmd.String("address", "", "Source Wallet address")
	sendAmount := sendCmd.Int("amount", 0,"Amount to send")

    switch os.Args[1] {

	case "getBalance" :
		err := getBalanceCmd.Parse(os.Args[2:])
		if err!= nil {
			log.Panic(err)
		}
	}

		case "createBlockChain" :
		err := createBlockChainCmd.Parse(os.Args[2:])
		if err!= nil {
		log.Panic(err)
		}
		}
	case "sendMoney" :
	err := sendCmd.Parse(os.Args[2:])
	if err!= nil {
	log.Panic(err)
	}
	}

		case "printChain" :
		err := printChainCmd.Parse(os.Args[2:])
	if err!= nil {
	log.Panic(err)
	}
	}

	case "listaddresses":
		err := listAddressCmd.Parse(os.Args[2:])
		Handle(err)

		case "createwallet":
     err := createWalletCmd.Parse(os.Args[2:])
	 Handle(err)



	default:
		cli.printUsage()
		runtime.GoExit()


}

if getBalanceCmd.Parsed(){
	if *getBalanceAddress =="" {
		getBalancedCmd.Usage()
		runtime.Goexit()

}
if createBlockChainCmd.Parsed(){
if *createBlockChainAddress =="" {
createBlockChainCmd.Usage()
runtime.Goexit()

}
cli.createBlockChain(*createBlockChainAddress)
}


if printChainCmd.Parsed(){
	cli.PrintChain()
}

if sendCmd.Parsed(){
	if *sendForm == "" !! *sendTo =="" || *sendAmount < 0 {
		sendCmd.Usage()
		runtime.Goexxit()

}

cli.send(*sendFrom, *sendTo, *SendAmount)
}


if listAddressCmd.Parsed(){
	cli.listAddress()
}

if createWalletCmd.Parsed(){
	cli.createWallet()
}



}


func main(){

	defer os.Exit(0)

	cli.CommandLine{}

	cli.Run()
}
