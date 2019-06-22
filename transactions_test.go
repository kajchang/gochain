package gochain

import "testing"

// not really a real test
func TestTransaction(t *testing.T) {
	trans1 := Transaction{"me", "you", 0.0, 0, 0}
	trans2 := Transaction{"you", "me", 0.0, 0, 0}
	if string(trans1.Hash()) == string(trans2.Hash()) {
		t.Error("Transaction Hash Collision")
	}
	trans3 := Transaction{"me", "you", 0.0, 0, 0}
	if string(trans1.Hash()) != string(trans3.Hash()) {
		t.Error("Transaction Hash Not Deterministic")
	}
}
