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
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"github.com/mr-tron/base58"
	"log"
)

// wallet implementation

const (
	checksumLength = 4
	version        = byte(0x00)
)

type Wallet struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey  []byte
}

func NewKeyPair() (ecdsa.PrivateKey, []byte) {

	curve := elliptic.P256()

	private, err := ecdsa.GenerateKey(curve, rand.Reader)

	if err != nil {
		log.Panicln(err)

	}

	pub := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)

	return *private, pub

}

func base58Encode(input []byte) []byte {
	encode := base58.Encode(input)
	return []byte(encode)
}

func base58Decode(input []byte) []byte {
	decode, err := base58.Decode(string(input[:]))
	if err != nil {
		log.Panic(err)
	}

	return decode
}

// private key to public key
func PublicKeyHash(publicKey []byte) []byte {
	hashedPublicKey := sha256.Sum256(publicKey)
	hasher := ripemd160.New()
	_, err := hasher.Write(hashedPublicKey[:])

	if err != nil {
		log.Panic(err)
	}

	publicRipeMd := hasher.Sum(nil)

	return publicRipeMd
}

//check sum

func CheckSum(ripeMdHash []byte) []byte {
	firstHash := sha256.Sum256(ripeMdHash)
	secondHash := sha256.Sum256(firstHash[:])

	return secondHash[:checksumLength]

}

func (w *Wallet) Address() []byte {

	pubHash := PublicKeyHash(w.PublicKey)
	versionedHash := append([]byte{version}, pubHash...)
	checkSum := CheckSum(versionedHash)
	finalHash := append(versionedHash, checkSum...)
	address := base58Encode(finalHash)

	return address
}

//initialize a wallet

func MakeWallet() *Wallet {

	privateKey, PublicKey := NewKeyPair()
	wallet := Wallet{privateKey, PublicKey}

	return &wallet
}
