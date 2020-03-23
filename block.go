package gochain

import (
	"bytes"
	"crypto/sha256"
)

type Block struct {
	Transactions []Transaction
	PreviousHash []byte
	Timestamp    uint64
	Nonce        uint64
}

func (b Block) Header() []byte {
	var buf bytes.Buffer
	buf.Write(b.PreviousHash)
	buf.Write(EncodeUint64(b.Timestamp))
	buf.Write(EncodeUint64(b.Nonce))
	for _, transaction := range b.Transactions {
		buf.Write(transaction.ToBuffer())
	}
	return buf.Bytes()
}

func (b Block) Hash() []byte {
	h := sha256.New()
	h.Write(b.Header())
	return h.Sum(nil)
}
