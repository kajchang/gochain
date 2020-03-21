package gochain

import (
	"bytes"
	"crypto/sha256"
)

type Transaction struct {
	From      []byte
	To        []byte
	Amount    float64
	Timestamp uint64
	Nonce     uint64
}

func (t Transaction) ToBuffer() []byte {
	var buf bytes.Buffer
	buf.Write(t.From)
	buf.Write(t.To)
	buf.Write(EncodeFloat64(t.Amount))
	buf.Write(EncodeUint64(t.Timestamp))
	buf.Write(EncodeUint64(t.Nonce))
	return buf.Bytes()
}

func (t Transaction) Hash() []byte {
	h := sha256.New()
	h.Write(t.ToBuffer())
	return h.Sum(nil)
}
