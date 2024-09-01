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
	"blockchain"
	"flag"
	"fmt"
	"os"
	"runtime"
	"strconv"
)

// create a basic block (transaction block) first
// here is the structure of block
/*
type BLOCK struct {
	Hash     []byte // byte array
	Data     []byte
	prevHash []byte
	nonce    int
}

//build a chain for this block

type BlockChain struct {
	blocks []*BLOCK //array of BLOCK pointers

}

// write a function to create a hash for  current data and previous hash

func (b *BLOCK) DeriveHsh() {
	info := bytes.Join([][]byte{b.Data, b.prevHash}, []byte{})

	hash := sha256.Sum256(info)

	b.Hash = hash[:]

}

// function for creating a new block

func Createblock(data string, prevHash []byte) *BLOCK {

	block := &BLOCK{[]byte{}, []byte(data), prevHash, 0} //empty byte array - first parameter

	block.DeriveHsh()

	return block
}

// function for adding a block to existing chain
func (chain *BlockChain) AddBlock(data string) {

	prevBlock := chain.blocks[len(chain.blocks)-1]
	newBlock := Createblock(data, prevBlock.Hash)

	chain.blocks = append(chain.blocks, newBlock)

}

// OG Block for previous hash
func Genesis() *BLOCK {
	return Createblock("Genesis", []byte{})
}

// initialize a new blockchain with new Genesis block made earlier
// after new block struct type
func InitializeBlockchain() *BlockChain {
	return &BlockChain{[]*BLOCK{Genesis()}}
}
*/
func main() {

	chain := InitializeBlockchain()

	chain.AddBlock("1st Genesis")
	chain.AddBlock("2nd Genesis")
	chain.AddBlock("3rd Genesis")

	for _, block := range chain.blocks {

		//fmt.Printf("%x", block.prevHash)  //old calls
		//fmt.Printf("%s", block.Data)
		//fmt.Printf("%s", block.Hash)

		fmt.Printf("%x", block.prevHash) //new calls
		fmt.Printf("%s", block.Data)
		fmt.Printf("%s", block.Hash)

		pow := NewProofOfWork(block)

		fmt.Printf("POW:%s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

	}

}

// after iterator database
type CommandLine struct {
	blockchain *blockchain.BlockChain2
}

// printUsage will display the options
func (cli *CommandLine) printUsage() {
	fmt.Println("Usage")
	fmt.Println(" add - block <BLOCK Data> add a block to chain")
	fmt.Println("print the block chain")

}

// validate Args

func (cli *CommandLine) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		runtime.Goexit()
	}
}

// add block allow users to add blocks to the chain via cli
func (cli *CommandLine) addBlock(data string) {
	cli.blockchain.AddBlock(data)
	fmt.Println("Added Block!")

}

// Print chain will display the entire contents of the block chain
func (cli *CommandLine) printChain() {
	iterator := cli.blockchain.Iterator()

	for {
		block := iterator.Next()
		fmt.Printf("Previous Hash: %s \n", block.prevHash)
		fmt.Printf("Data Hash: %s \n", block.Data)
		fmt.Printf("Hash: %s \n", block.Hash)

		pow := cli.blockchain.NewProofWork(block) // new blockchain
		fmt.Printf("Pow %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		if len(block.prevHash) == 0 {
			break

		}

	}
}

// run command Line Program
func (cli *CommandLine) run() {
	cli.validateArgs()

	addBlockCmd := flag.NewFlagSet("add", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("print", flag.ExitOnError)
	addBlockData := addBlockCmd.String("", "", "Block data")

	switch os.Args[1] {
	case "add":
		err := addBlockCmd.Parse(os.Args[2:])
		blockchain.Handle(err)

	case "print":
		err := printChainCmd.Parse(os.Args[2:])
		blockchain.Handle(err)

	default:
		cli.printUsage()
		runtime.Goexit()
	}

	//

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			runtime.Goexit()
		}
		cli.addBlock(*addBlockData)
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}

}

// new function main
func main() {

	defer os.Exit(0)

	chain := blockchain.InitBlockChain()

	defer chain.Database.Close()
	cli := Commandline{chain}
	cli.run()
}
