package kucoin

import (
	"crypto/sha256"
	"encoding/hex"
)

type Signer interface {
	Sign(plain []byte) string
}

type Sha256Signer struct {
	Key []byte
}

func (ss *Sha256Signer) Sign(plain []byte) string {
	h := sha256.New()
	h.Write(plain)
	return hex.EncodeToString(h.Sum(nil))
}
