package kucoin

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

type Signer interface {
	Sign(plain []byte) []byte
}

type Sha256Signer struct {
	key []byte
}

func (ss *Sha256Signer) Sign(plain []byte) []byte {
	hm := hmac.New(sha256.New, ss.key)
	hm.Write(plain)
	return hm.Sum(nil)
}

type KcSigner struct {
	Sha256Signer
	ApiKey        string
	ApiSecret     string
	ApiPassPhrase string
}

func (ks *KcSigner) Sign(plain []byte) []byte {
	s := ks.Sha256Signer.Sign(plain)
	return []byte(base64.StdEncoding.EncodeToString(s))
}

func NewKcSigner(key, secret, passPhrase string) *KcSigner {
	ks := &KcSigner{
		ApiKey:        key,
		ApiSecret:     secret,
		ApiPassPhrase: passPhrase,
	}
	ks.key = []byte(secret)
	return ks
}
