package main

import (
	"encoding/hex"
	"fmt"
	"github.com/kajchang/gochain"
)

func main() {
	myAddress := []byte ("~~~~~~~~~~~~~~~~ME~~~~~~~~~~~~~~~~")
	blockchain := gochain.Genesis(myAddress)
	for len(blockchain) < 50 {
		blockchain = append(blockchain, blockchain.MineBlock(blockchain.GenerateCoinbaseTransaction(myAddress).ToBuffer()))
	}
	fmt.Println(blockchain)
}
