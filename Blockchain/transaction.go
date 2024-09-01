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
	"fmt"
)

const reward = 100

type Transaction struct {
	ID      []byte
	Inputs  []TxInput
	Outputs []TxOutput
}

type TxOutput struct {
	Value  int
	PubKey string
}

type TxInput struct {
	ID  []byte
	Out int
	Sig string
}

func CoinbaseTx(toAddress, data string) *Transaction {

	if data == "" {

		data = fmt.Sprintf("Coins to %s", toAddress)

	}

	txIn := TxInput{[]byte{}, -1, data}
	txOut := TxOutput{reward, toAddress}

	tx := Transaction{nil, []TxInput{txIn}, []TxOutput{txOut}}

	return &tx
}

func (tx *Transaction) SetID() {
	var encoded bytes.Buffer
	var hash [32]byte

	encoder := gob.NewEncoder(&encoded)
	err := encoder.Encode(tx)
	Handle(err)

	hash = sha256.Sum256(encoded.Bytes())
	tx.ID := hash[:]

}

func (in *TxInput) CanUnlock(data string ) bool {
	return in.Sig == data
}

func (out *TxOutput) CanBeUnlocked(data string) bool {
	return out.PubKey == data
}


func (tx *Transaction) IsCoinbase() bool {

	return len(tx.Inputs) == 1 && len (tx.Inputs[0].ID) ==0

}
