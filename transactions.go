package gochain

import (
	"bytes"
	"crypto/sha256"
)

type Transaction struct {
	from      string
	to        string
	amount    float64
	timestamp uint64
	nonce     uint64
}

func (t Transaction) toBuffer() []byte {
	var buf bytes.Buffer
	buf.Write([]byte(t.from))
	buf.Write([]byte(t.to))
	buf.Write(EncodeFloat64(t.amount))
	buf.Write(EncodeUint64(t.timestamp))
	buf.Write(EncodeUint64(t.nonce))
	return buf.Bytes()
}

func (t Transaction) Hash() []byte {
	h := sha256.New()
	h.Write(t.toBuffer())
	return h.Sum(nil)
}
