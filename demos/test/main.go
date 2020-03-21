package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"github.com/kajchang/gochain"
	"time"
)

func main() {
	myKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	myAddress := elliptic.Marshal(myKey, myKey.X, myKey.Y)

	anotherKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	anotherAddress := elliptic.Marshal(anotherKey, anotherKey.X, anotherKey.Y)

	blockchain := gochain.Genesis(myAddress)
	for len(blockchain) < 50 {
		blockchain = append(blockchain, blockchain.MineBlock(blockchain.GenerateCoinbaseTransaction(myAddress).ToBuffer()))
	}

	myTransaction := gochain.Transaction{
		From:      myAddress,
		To:        anotherAddress,
		Amount:    50,
		Timestamp: uint64 (time.Now().Unix()),
		Nonce:     0,
		Signature: nil,
	}

	myTransaction.Sign(myKey)

	fmt.Println(blockchain)
}
