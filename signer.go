package kucoin

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

type Signer interface {
	Sign(plain []byte) string
}

type Sha256Signer struct {
	key []byte
}

func (ss *Sha256Signer) Sign(plain []byte) string {
	hm := hmac.New(sha256.New, ss.key)
	hm.Write(plain)
	return hex.EncodeToString(hm.Sum(nil))
}

type KcSigner struct {
	Sha256Signer
	ApiKey        string
	ApiSecret     string
	ApiPassphrase string
}

func NewKcSigner(key, secret, passphrase string) *KcSigner {
	ks := &KcSigner{
		ApiKey:        key,
		ApiSecret:     secret,
		ApiPassphrase: passphrase,
	}
	ks.key = []byte(secret)
	return ks
}
