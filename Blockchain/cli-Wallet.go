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

	//"wallet"
)

func (cli *CommandLine) printUsageWallet() {

	fmt.Println("createwallet - Creates a new wallet list")
	fmt.Println("listaddress - Lists the addresses in wallet file")
	//fmt.Println("createwallet - Creates a new wallet list")

}

// list address method
func (cli *CommandLine) listAddress() {
	wallets, _ := wallet.CreateWallet()

	addresses := address.GetAllAddresses()

	for _, address := range addresses {
		fmt.Println(address)

	}

}

//create wallet

func (cli *CommandLine) CreateWallet() {
	wallets, _ := wallet.CreateWallets()
	address := wallets.AddWallet()
	wallets.SaveFile()

	fmt.Printf(" New Address is %s\n", address)

}

// Run method
//func (cli *CommandLine) Run() {}
func (cli *CommandLine) Run() {
	cli.validateArgs()
	getBalanceCmd := flag.NewFlagSet("getbalance", flag.ExitOnError)
	createBlockchainCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError
	sendCmd := flag.NewFlagSet("send", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	createWalletCmd := flag.NewFlagSet("createwallet", flag.ExitOnError) // thi
	listAddressesCmd := flag.NewFlagSet("listaddresses", flag.ExitOnError) //


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
