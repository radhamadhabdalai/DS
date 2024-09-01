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
	"bytes"
	"crypto/sha256"
	"fmt"
)

// create a basic block (transaction block) first
// here is the structure of block

type BLOCK struct {
	Hash     []byte // byte array
	Data     []byte
	prevHash []byte
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

	block := &BLOCK{[]byte{}, []byte(data), prevHash}

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

//initialize a new blockchain with new Genesis block made earlier

func InitializeBlockchain() *BlockChain {
	return &BlockChain{[]*BLOCK{Genesis()}}
}

func main() {
	chain := InitializeBlockchain()

	chain.AddBlock("1st Genesis")
	chain.AddBlock("2nd Genesis")
	chain.AddBlock("3rd Genesis")

	for _, block := range chain.blocks {

		fmt.Printf("%x", block.prevHash)
		fmt.Printf("%s", block.Data)
		fmt.Printf("%s", block.Hash)

	}

}
