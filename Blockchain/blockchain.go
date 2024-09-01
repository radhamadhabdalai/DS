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
	"encoding/hex"
	"fmt"
	"github.com/dgraph-io/badger" // This is our database import
	"log"
	"os"
	"runtime"
)

const (
	dbPath      = "./tmps/blocks"
	dbFile      = "./tmps/blocks/MANIFEST"
	GenesisData = "First Genesis"
)

type BlockChain2 struct {
	LastHash []byte
	Database *badger.DB
}

type BlockChainIterator struct {
	CurrentHash []byte
	Database    *badger.DB
}

// check db exists or not

func DBExists(db) bool {
	if _, err := os.Stat(db); os.IsNotExist(err) {
		return false
	}
	return true
}

// updated routine for initializing block chain with  previous hash and database

// change the code for checking database exists

func InitBlockChain() *BlockChain2 {

	var lastHash []byte
	opts := badger.DefaultOptions(dbPath)
	db, err := badger.Open(opts)
	Handle(err)

	err = db.Update(func(txn *badger.Txn) error {

		if _, err := txn.Get([]byte("lh")); err == badger.ErrKeyNotFound {
			fmt.Println("No existing blockchain found")
			genesis := Genesis()
			fmt.Println("Genesis proved")
			err = txn.Set(genesis.Hash, genesis.Serialize())
			Handle(err)
			err = txn.Set([]byte("lh"), genesis.Hash)

			lastHash := genesis.Hash

			return err

		} else {
			item, err := txn.Get([]byte("lh"))
			Handle(err)
			err = item.Value(func(val []byte) error {
				lastHash = val
				return nil
			})
			Handle(err)
			return err

		}

	})

	Handle(err)
	//blockHash := BLOCK{[]byte{}, []byte{}, lastHash, 0}

	//blockChain := BlockChain{lastHash, db}

	blockChain := BlockChain2{lastHash, db}
	return &blockChain
}

// updated initBlock chain for tranaction handling
func InitBlockChain(address string) *BlockChain {
	var lastHash []byte

	if DBExists(dbFile) {

		fmt.Println("Db file already exists")
		runtime.Goexit()
	}

	opts := badger.DefaultOptions(dbPath)
	db, err := badger.Open(opts)
	Handle(err)

	err = db.Update(func(txn *badger.Txn) error {

		cbTx := CoinbaseTx(address, genesisData)
		genesis := Genesis(cbTx)
		fmt.Println("Genesis Created")
		err := txn.Set(genesis.Hash, genesis.Serialize())
		Handle(err)
		err = txn.Set([]byte("lh"), genesis.Hash)
		lastHash = genesis.Hash
		return err
	})

}

// go routine for continue block chain

func ContinueBlockChain(address string) *BlockChain {

	if DBExists(dbFile) == false {
		fmt.Println("No Blockchain found")
		runtime.Goexit()

	}

	var lastHash []byte

	opts := badger.DefaultOptions(dbPath)
	db, err := badger.Open(opts)
	Handle(err)

	err = db.Update(func(txn *badger.Txn) error {

		item, err := txn.Get([]byte("lh"))
		Handle(err)
		err = item.Value(func(val []byte) error {
			lastHash = val
			return nil
		})
		Handle(err)
		return err
	})

	chain := BlockChain{lastHash, db} // check the error
	return &chain
}

// routine for add block

func (chain *BlockChain2) AddBlock(data string) {

	var lastHash []byte

	err := chain.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("lh"))
		Handle(err)
		err = item.Value(func(val []byte) error {
			lastHash = val
			return nil
		})

		Handle(err)
		return err
	})

	Handle(err)
	newBlock := Createblock(data, lastHash)

	err = chain.Database.Update(func(transaction *badger.Txn) error {
		err := transaction.Set(newBlock.Hash, newBlock.Serialize())

		chain.LastHash = newBlock.Hash
		return err
	})

	Handle(err)
}

// routine for unspent transactions
func (chain *BlockChain) FindUnspentTransactions(address string) []Transaction {
	var unspentTxs []Transaction

	spentTxns := make(map[string][]int)

	iter := chain.Iterator()
	for {
		block := iter.Next()

		for _, tx := range block.Transactions {

			txID := hex.EncodeToString(tx.ID)

		Outputs:
			for outIDx, out := range tx.Outputs {
				if spentTxns[txID] != nil {
					for _, spentOut := range spentTxns[txID] {
						if spentOut == outIDx {
							continue Outputs
						}
					}
				}
				if out.CanBeUnlocked(address) {
					unspentTxs = append(unspentTxs, *tx)

				}
			}

			if tx.IsCoinBase() == false {

				for _, in := range tx.Inputs {
					if in.CanUnlock(address) {
						inTxID := hex.EncodeToString(in.ID)
						spentTxns[inTxID] = append(spentTxns[inTxID], in.Out)

					}
				}

			}

		}

		if len(block.PrevHash) == 0 {
			break
		}
	}

	return unspentTxs
}

// find unspendable outputs

func (chain *Blockchain) FindSpendableOutputs(address string, amount int) (int, map[string][]int) {

	unspentAmounts := make(map[string][]int)
	unspentTxs := chain.FindUnspentTransactions(address)
	accumulated := 0
Work:
	for _, tx := range unpspentTxs {
		txID := hex.EncodeToString(tx.ID)

		for outIds, out := range tx.Outputs {
			if out.CanBeUnlocked(address) && accumulated < amount {
				accumulated += out.Value
				unspentAmounts[txID] = append(unspentTxs[txId], outIds)

				if accumulated >= amount {
					break Work
				}

			}
		}
	}
	return accumulated, unspentAmounts
}

// data string -> Transaction block
func (chain *Blockchain) Addblock(transactions []*Transaction) {
	var lastHash []byte

	err := chain.Database.View(func(txn *badger.Txn) error {

		item, err := txn.Get([]byte("lh"))
		Handle(err)
		err = item.Value(func(val []byte) error {
			lastHash = val
			return nil

		})

		Handle(err)
		return err
	})

	Handle(err)

	newBlock := Createblock(transactions, lastHash)

	err = chain.Database.Update(func(transaction *badger.Txn) error {
		err := transaction.Set(newBlock.Hash, newBlock.Serialize())
		Handle(err)
		err = transaction.Set([]byte("lh"), newBlock.Hash)
		chain.LastHash = newBlock.Hash
		return err
	})

	Handle(err)

}

// New Transaction

func NewTransaction(from, to string, amount int, chain *Blockchain) *Transaction {
	var inputs []TxInput
	var ouputs []TxOutput

	acc, ValidOutputs := chain.FindSpendableOutputs(from, amount)

	if acc < amount {
		log.Panic("Error, Not funds")

	}
	for txid, outs := range ValidOutputs {
		txID, err := hex.DecodeString(txid)
		Handle(err)
	}

	for _, out := range outs {

		input := TxInput{txID, out, from}
		inputs = append(inputs, input)
	}
}
     outputs = append(outputs, TxOutput { amount ,to})


	  if acc > amount {

            outputs = append(outputs, TxOutput { acc - amount ,from})
	  }

	 tx := Transaction{nil, inputs, outputs}
	 tx.setID()

return &tx
}

// find unspent transactions output

func (chain *BlockChain) findUTXO(address string) []TxOutput {

	var UTXOs []TxOutput

	unspentTransactions := chain.FindUnspentTransactions(address)

	for _, tx := range unspentTransactions {

		for _, out := range tx.Outputs {
			if CanBeUnocked(address) {
				UTXOs = append(UTXOs, out)
			}
		}
	}

	return UTXOs
}

// routine to iterator process
func (chain *BlockChain2) Iterator() *BlockChainIterator {
	iterator := BlockChainIterator{chain.LastHash, chain.Database}
	return &iterator
}

// iterate through database
func (iterator *BlockChainIterator) Next() *BLOCK {

	var block *BLOCK

	err := iterator.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get(iterator.CurrentHash)

		Handle(err)

		err = item.Value(func(val []byte) error {
			block = Deserialize(val)
			return nil
		})

		Handle(err)
		return err
	})

	Handle(err)

	iterator.CurrentHash = block.prevHash
	return block
}
