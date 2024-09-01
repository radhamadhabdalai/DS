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
	"crypto/elliptic"
	"fmt"
	"os"

	//"crypto/ecdsa"
	"encoding/gob"
	"io/ioutil"
	"log"
)

const walletFile = "./tmp/wallets.data"

type Wallets struct {
	Wallets map[string]*Wallet
}

// save file routine method
func (ws *Wallets) saveFile() {

	var content bytes.Buffer

	gob.Register(elliptic.P256()) //p256 not working
	//now working

	encoder := gob.NewEncoder(&content)
	err := encoder.Encode(ws)

	if err != nil {
		log.Panic(err)
	}
	err = ioutil.WriteFile(walletFile, content.Bytes(), 0644)
	Handle(err)
}

//
//func Handle(err error) {
//	if err != nil {
//		log.Panic(err)
//	}
//
//}

// load file
func (ws *Wallets) LoadFIle() error {
	if _, err := os.Stat(walletFile); os.IsNotExist(err) {
		return err
	}
	var wallets Wallets

	fileContent, err := ioutil.ReadFile(walletFile)

	gob.Register(elliptic.P256())

	decoder := gob.NewDecoder(bytes.NewReader(fileContent))
	err = decoder.Decode(&wallets)
	Handle(err)
	ws.Wallets = wallets.Wallets

	return nil
}

// add wallet
func (ws *Wallets) AddWallet() string {
	wallet := MakeWallet()
	address := fmt.Sprintf("%s", wallet.Address())

	ws.Wallets[address] = wallet
	return address
}

// get wallet
func (ws Wallets) GetWallet(address string) Wallet {
	return *ws.Wallets[address]
}

// get Wallet address
func (ws *Wallets) GetAllAddress() []string {
	var addresses []string

	for address := range ws.Wallets {
		addresses = append(addresses, address)
	}

	return addresses
}
