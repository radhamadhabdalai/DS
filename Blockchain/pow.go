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
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
)

const DIFFICULTY = 12

// create a proof of work struct
type ProofOfWork struct {
	Block  *BLOCK
	Target *big.Int
}

// initialize pow struct
func NewProofOfWork(b *BLOCK) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-DIFFICULTY))

	pow := &ProofOfWork{b, target}

	return pow
}

// convert int to --> []byte

func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)

	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}

// create a nonce

func (pow *ProofOfWork) InitNonce(nonce int) []byte {
	data := bytes.Join([][]byte{pow.Block.prevHash, pow.Block.Data, ToHex(int64(nonce)), ToHex(int64(DIFFICULTY))}, []byte{})

	return data
}

// running the algorithm
func (pow *ProofOfWork) Run() (int, []byte) {
	var intHash big.Int
	var hash [32]byte

	nonce := 0

	for nonce < math.MaxInt64 { // going for an infinite loop

		data := pow.InitNonce(nonce) // look at here
		hash = sha256.Sum256(data)   // look at here

		fmt.Printf("\r%x", hash)
		intHash.SetBytes(hash[:])

		if intHash.Cmp(pow.Target) == -1 {
			break
		} else {
			nonce++
		}
	}

	return nonce, hash[:]
}

// validating algorithm
func (pow *ProofOfWork) Validate() bool {
	var intHash big.Int
	data := pow.InitNonce(pow.Block.nonce) // check it for "Nonce" as it is not in the fields of BLOCK struct

	hash := sha256.Sum256(data)
	intHash.SetBytes(hash[:])

	return intHash.Cmp(pow.Target) == -1
}

// initData for transaction
func (pow *ProofOfWork) initData(nonce int64) []byte {
	data := bytes.Join(
		[][]byte{pow.Block.prevHash, pow.Block.HashTransations(), ToHex(nonce), ToHex(int64(Difference))}, []byte{})

	return data
}
