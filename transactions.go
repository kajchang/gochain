package gochain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"math"
)

type Transaction struct {
	from      string
	to        string
	amount    float64
	timestamp uint64
	nonce     uint64
}

func encodeFloat64(f float64) []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, math.Float64bits(f))
	return buf
}

func encodeInt64(i uint64) []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, i)
	return buf
}

func (t Transaction) toBuffer() []byte {
	var buf bytes.Buffer
	buf.Write([]byte(t.from))
	buf.Write([]byte(t.to))
	buf.Write(encodeFloat64(t.amount))
	buf.Write(encodeInt64(t.timestamp))
	buf.Write(encodeInt64(t.nonce))
	return buf.Bytes()
}

func (t Transaction) Hash() []byte {
	h := sha256.New()
	h.Write(t.toBuffer())
	return h.Sum(nil)
}
