package gochain

import (
	"bytes"
	"crypto/sha256"
)

type Block struct {
	transactions []Transaction
	previousHash []byte
	timestamp    uint64
}

func (b Block) toBuffer() []byte {
	var buf bytes.Buffer
	buf.Write(b.previousHash)
	buf.Write(EncodeUint64(b.timestamp))
	for _, t := range b.transactions {
		buf.Write(t.Hash())
	}
	return buf.Bytes()
}

func (b Block) Hash() []byte {
	h := sha256.New()
	h.Write(b.toBuffer())
	return h.Sum(nil)
}
