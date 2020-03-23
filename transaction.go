package gochain

import (
	"bytes"
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/asn1"
	"math/big"
)

type Transaction struct {
	From      []byte
	To        []byte
	Amount    float64
	Timestamp uint64
	Nonce     uint64
	Signature []byte
}

func (t Transaction) Header() []byte {
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
	h.Write(t.Header())
	return h.Sum(nil)
}

func (t Transaction) ToBuffer() []byte {
	var buf bytes.Buffer
	buf.Write(t.Header())
	buf.Write(t.Signature)
	return buf.Bytes()
}

func (t *Transaction) Sign(key *ecdsa.PrivateKey) {
	signature, _ := key.Sign(rand.Reader, t.Hash(), crypto.SHA256)
	t.Signature = signature
}

func (t *Transaction) VerifySignature() bool {
	x, y := elliptic.Unmarshal(StandardCurve, t.From)
	pk := ecdsa.PublicKey{
		Curve: StandardCurve,
		X:     x,
		Y:     y,
	}
	var esig struct {
		R, S *big.Int
	}
	_, err := asn1.Unmarshal(t.Signature, &esig)
	if err != nil {
		return false
	}
	return ecdsa.Verify(&pk, t.Hash(), esig.R, esig.S)
}
