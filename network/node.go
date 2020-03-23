package network

import (
	"bytes"
	"github.com/kajchang/gochain/core"
)

type Node struct {
	Blockchain      *core.Blockchain
	transactionPool []core.Transaction
}

func (node *Node) AddTransactionToPool(transaction core.Transaction) {
	if transaction.VerifyIntegrity() {
		node.transactionPool = append(node.transactionPool, transaction)
	}
}

func (node Node) AssembleValidTransactions() []core.Transaction {
	validTransactions := make([]core.Transaction, 0)

	addressCached := make(map[string]bool)
	nonceCache := make(map[string]uint64)
	balanceCache := make(map[string]float64)

	flag := false

	for {
		for _, transaction := range node.transactionPool {
			strAddress := string(transaction.From)
			if !addressCached[strAddress] {
				nonceCache[strAddress] = node.Blockchain.GetAddressNonce(transaction.From)
				balanceCache[strAddress] = node.Blockchain.GetBalance(transaction.From)
				addressCached[strAddress] = true
			}

			if transaction.Nonce == nonceCache[strAddress] &&
				transaction.In <= balanceCache[strAddress] {
				nonceCache[strAddress]++
				balanceCache[strAddress] -= transaction.In
				validTransactions = append(validTransactions, transaction)
				flag = true
			}
		}

		if flag {
			flag = false
		} else {
			break
		}
	}

	return validTransactions
}

func (node Node) VerifyBlock(block core.Block) bool {
	transactionCache := make(map[int]bool)

	net := 0.0

	addressCached := make(map[string]bool)
	nonceCache := make(map[string]uint64)
	balanceCache := make(map[string]float64)

	flag := false

	for {
		for i, transaction := range block.Transactions {
			if transactionCache[i] {
				continue
			}

			strAddress := string(transaction.From)
			if !addressCached[strAddress] {
				nonceCache[strAddress] = node.Blockchain.GetAddressNonce(transaction.From)
				balanceCache[strAddress] = node.Blockchain.GetBalance(transaction.From)
				addressCached[strAddress] = true
			}

			if !bytes.Equal(transaction.From, core.CoinbaseAddress) &&
				(!transaction.VerifyIntegrity() || transaction.In > balanceCache[strAddress]) {
				return false
			}

			if bytes.Equal(transaction.From, core.CoinbaseAddress) || transaction.Nonce == nonceCache[strAddress] {
				nonceCache[strAddress]++
				balanceCache[strAddress] -= transaction.In
				transactionCache[i] = true
				flag = true
			}

			net -= transaction.In
			net += transaction.Out
		}

		if len(transactionCache) == len(block.Transactions) {
			return node.Blockchain.ValidNextHash(block) && net <= node.Blockchain.GetCoinbase()
		}

		if flag {
			flag = false
		} else {
			return false
		}
	}
}
