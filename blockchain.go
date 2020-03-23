package gochain

import (
	"bytes"
	"encoding/hex"
	"strconv"
	"time"
)

const StartingCoinbase = 50
var CoinbaseAddress = make([]byte, 65)

type Blockchain []Block

func (blockchain Blockchain) GetDifficulty() []byte {
	return []byte {1}
}

func (blockchain Blockchain) GetCoinbase() float64 {
	return StartingCoinbase
}

func (blockchain Blockchain) GetBalance(address []byte) float64 {
	balance := 0.0
	for _, block := range blockchain {
		for _, transaction := range block.Transactions {
			if bytes.Equal(transaction.To, address) {
				balance += transaction.Out
			}
			if bytes.Equal(transaction.From, address) {
				balance -= transaction.In
			}
		}
	}
	return balance
}

func (blockchain Blockchain) GenerateCoinbaseTransaction(minerAddress []byte) Transaction {
	return Transaction{
		From:      CoinbaseAddress,
		To:        minerAddress,
		In:        0,
		Out:       blockchain.GetCoinbase(),
		Timestamp: uint64 (time.Now().Unix()),
		Nonce:     0,
		Signature: make([]byte, 64),
	}
}

func (blockchain Blockchain) MineBlock(transactions []Transaction) Block {
	var nonce uint64 = 0
	for {
		previousHash := make([]byte, 32)
		if len(blockchain) != 0 {
			previousHash = blockchain[len(blockchain) - 1].Hash()
		}
		timestamp := uint64 (time.Now().Unix())
		candidateBlock := Block{
			Transactions: transactions,
			PreviousHash: previousHash,
			Timestamp:    timestamp,
			Nonce:        nonce,
		}
		if blockchain.ValidNextHash(candidateBlock) {
			return candidateBlock
		}
		nonce++
	}
}

func (blockchain Blockchain) ValidNextHash(block Block) bool {
	return bytes.HasPrefix(block.Hash(), blockchain.GetDifficulty())
}

func (blockchain Blockchain) VerifyBlock(block Block) bool {
	net := 0.0
	for _, transaction := range block.Transactions {
		if !bytes.Equal(transaction.From, CoinbaseAddress) &&
			(
				transaction.Out > transaction.In ||
				!transaction.VerifySignature() ||
				blockchain.GetBalance(transaction.From) < transaction.In) {
			return false
		}
		net -= transaction.In
		net += transaction.Out
	}

	return blockchain.ValidNextHash(block) && net <= blockchain.GetCoinbase()
}

func Genesis(genesisAddress []byte) Blockchain {
	blockchain := Blockchain{}
	genesisCoinbase := blockchain.GenerateCoinbaseTransaction(genesisAddress)
	blockchain = append(blockchain, blockchain.MineBlock([]Transaction{genesisCoinbase}))
	return blockchain
}

func (blockchain Blockchain) String() string {
	result := "|" + MiddlePadString("Index", 8) + "|" +
		      MiddlePadString("Previous Hash", 68) + "|" +
		      MiddlePadString("Hash", 68) + "|" +
		      MiddlePadString("Size", 8) + "|\n"
	for i, block := range blockchain {
		result += "|" + MiddlePadString(strconv.Itoa(i), 8) + "|" +
			      MiddlePadString(hex.EncodeToString(block.PreviousHash), 68) + "|" +
			      MiddlePadString(hex.EncodeToString(block.Hash()), 68) + "|" +
			      MiddlePadString(strconv.Itoa(len(block.Transactions)), 8) + "|\n"
	}
	return result
}
