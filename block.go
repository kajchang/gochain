package gochain

import (
	"bytes"
	"crypto/sha256"
)

type Block struct {
	Data         []byte
	PreviousHash []byte
	Timestamp    uint64
	Nonce        uint64
}

func (b Block) ToBuffer() []byte {
	var buf bytes.Buffer
	buf.Write(b.PreviousHash)
	buf.Write(EncodeUint64(b.Timestamp))
	buf.Write(EncodeUint64(b.Nonce))
	buf.Write(b.Data)
	return buf.Bytes()
}

func (b Block) Hash() []byte {
	h := sha256.New()
	h.Write(b.ToBuffer())
	return h.Sum(nil)
}
