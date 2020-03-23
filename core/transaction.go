package core

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"github.com/kajchang/gochain"
	"math/big"
)

var StandardCurve = elliptic.P256()

/*
65 bytes - From
65 bytes - To
8 bytes - In
8 bytes - Out
8 bytes - Timestamp
8 bytes - Nonce
64 bytes - Signature
--
226 bytes - Total
*/

type Transaction struct {
	From      []byte
	To        []byte
	In        float64
	Out       float64
	Timestamp uint64
	Nonce     uint64
	Signature []byte
}

func (t Transaction) Header() []byte {
	var buf bytes.Buffer
	buf.Write(t.From)
	buf.Write(t.To)
	buf.Write(gochain.EncodeFloat64(t.In))
	buf.Write(gochain.EncodeFloat64(t.Out))
	buf.Write(gochain.EncodeUint64(t.Timestamp))
	buf.Write(gochain.EncodeUint64(t.Nonce))
	return buf.Bytes()
}

func (t Transaction) Hash() []byte {
	h := sha256.New()
	h.Write(t.Header())
	return h.Sum(nil)
}

func (t Transaction) ToBuffer() []byte {
	var buf bytes.Buffer
	buf.Write(t.Header())
	buf.Write(t.Signature)
	return buf.Bytes()
}

func (t *Transaction) Sign(key *ecdsa.PrivateKey) (err error) {
	r, s, err := ecdsa.Sign(rand.Reader, key, t.Hash())
	if err != nil {
		return err
	}
	t.Signature = append(r.Bytes(), s.Bytes()...)
	return nil
}

func (t *Transaction) VerifySignature() bool {
	x, y := elliptic.Unmarshal(StandardCurve, t.From)
	pk := ecdsa.PublicKey{
		Curve: StandardCurve,
		X:     x,
		Y:     y,
	}
	r := new(big.Int).SetBytes(t.Signature[:32])
	s := new(big.Int).SetBytes(t.Signature[32:])
	return ecdsa.Verify(&pk, t.Hash(), r, s)
}

func (t Transaction) VerifyIntegrity() bool {
	return t.Out <= t.In && t.VerifySignature()
}
