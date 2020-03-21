package gochain

import (
	"bytes"
	"encoding/hex"
	"strconv"
	"time"
)

type Blockchain []Block

func (blockchain Blockchain) GetDifficulty() []byte {
	return []byte {1}
}

func (blockchain Blockchain) GetCoinbase() float64 {
	return StartingCoinbase
}

func (blockchain Blockchain) GenerateCoinbaseTransaction(minerAddress []byte) Transaction {
	return Transaction{
		From:      make([]byte, 65),
		To:        minerAddress,
		Amount:    blockchain.GetCoinbase(),
		Timestamp: uint64 (time.Now().Unix()),
		Nonce:     0,
		Signature: make([]byte, 71),
	}
}

func (blockchain Blockchain) MineBlock(data []byte) Block {
	var nonce uint64 = 0
	for {
		previousHash := make([]byte, 32)
		if len(blockchain) != 0 {
			previousHash = blockchain[len(blockchain) - 1].Hash()
		}
		timestamp := uint64 (time.Now().Unix())
		candidateBlock := Block{
			Data:         data,
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

func Genesis(genesisAddress []byte) Blockchain {
	blockchain := Blockchain{}
	genesisCoinbase := blockchain.GenerateCoinbaseTransaction(genesisAddress)
	blockchain = append(blockchain, blockchain.MineBlock(genesisCoinbase.ToBuffer()))
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
			      MiddlePadString(strconv.Itoa(len(block.Data)), 8) + "|\n"
	}
	return result
}
