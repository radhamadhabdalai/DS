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
	"encoding/gob"
	"github.com/dgraph-io/badger"
	"log"
)

// add new Block structure

type BLOCK struct {
	Hash         []byte // byte array
	Data         []byte
	prevHash     []byte
	Transactions []*Transaction
	nonce        int
}

type BlockChain struct {
	blocks   []*BLOCK
	Database *badger.DB
}

func (b *BLOCK) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.prevHash}, []byte{})

	hash := sha256.Sum256(info)

	b.Hash = hash[:]

}
func createBlock(data string, prevHash []byte) *BLOCK {
	block := &BLOCK{[]byte{}, []byte(data), prevHash, 0}

	block.DeriveHash()
	return block
}

// creating a new block with Transaction // added latest
func CreateBlock(txs []*Transaction, prevHash []byte) *Block {
	block := &Block{[]byte{}, txs, prevHash, nil}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce
	return block

}

func Genesis(coinbase *Transaction) *BLOCK {
	return CreateBlock([]*Transaction{coinbase}, []byte{})
}

// data hashing

func (b *Block) HashTransactions() []byte {

	var txHashes [][]byte
	var txHash [32]byte

	for _, tx := range b.Transactions {
		txHashes = append(txHashes, tx.ID)

	}

	txHash = sha256.Sum256(bytes.Join(txHashes, []byte))
	return txHash[:]
}

// function for creating a new block

func Createblock(data string, prevHash []byte) *BLOCK {

	block := &BLOCK{[]byte{}, []byte(data), prevHash, 0} //empty byte array - first parameter

	block.DeriveHash()

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
	return &BlockChain{blocks: []*BLOCK{Genesis()}}
}

// new changes

// Handle error
// handle every single bit of error

func Handle(err error) {
	if err != nil {

		log.Panic(err)
	}

}

// Serialize - Convert Block Struct to Bytes

func (b *BLOCK) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)

	err := encoder.Encode(b)

	Handle(err)

	return res.Bytes()
}

// deserialize - convert byte [] to Block struct
func Deserialize(data []byte) *BLOCK {
	var block BLOCK
	decoder := gob.NewDecoder(bytes.NewReader(data))

	err := decoder.Decode(&block)

	Handle(err)

	return &block
}
