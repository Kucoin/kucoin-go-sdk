package kucoin

import "testing"

func TestKcSigner_Sign(t *testing.T) {
	s := NewKcSigner("abc", "efg", "kcs")
	b := []byte("GET/api/v1/orders")
	if string(s.Sign(b)) != "iOdkcc7K6cyY8Cdr3yMcTgXCof4vhHCaDyrDSG7Qf3w=" {
		t.Error("Invalid sign")
	}
}
