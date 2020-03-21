package main

import (
	"encoding/hex"
	"fmt"
	"github.com/kajchang/gochain"
	"time"
)

func main() {
	transaction := gochain.Transaction{
		From: []byte ("Coinbase"),
		To:   []byte ("Human   "),
		Amount: 50,
		Timestamp: uint64 (time.Now().Unix()),
		Nonce: 0,
	}
	fmt.Printf(hex.EncodeToString(transaction.Hash()))
}
