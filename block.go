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

func (b Block) toBuffer() []byte {
	var buf bytes.Buffer
	buf.Write(b.PreviousHash)
	buf.Write(EncodeUint64(b.Timestamp))
	buf.Write(EncodeUint64(b.Nonce))
	for _, t := range b.Transactions {
		buf.Write(t.toBuffer())
	}
	return buf.Bytes()
}

func (b Block) Hash() []byte {
	h := sha256.New()
	h.Write(b.toBuffer())
	return h.Sum(nil)
}
