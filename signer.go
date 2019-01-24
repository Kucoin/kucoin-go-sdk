package kucoin

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"time"
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
	apiKey        string
	apiSecret     string
	apiPassPhrase string
}

func (ks *KcSigner) Sign(plain []byte) []byte {
	s := ks.Sha256Signer.Sign(plain)
	return []byte(base64.StdEncoding.EncodeToString(s))
}

func (ks *KcSigner) Headers(plain string) map[string]string {
	t := IntToString(time.Now().UnixNano() / 1000000)
	p := []byte(t + plain)
	s := string(ks.Sign(p))
	return map[string]string{
		"KC-API-KEY":        ks.apiKey,
		"KC-API-PASSPHRASE": ks.apiPassPhrase,
		"KC-API-TIMESTAMP":  t,
		"KC-API-SIGN":       s,
	}
}

func NewKcSigner(key, secret, passPhrase string) *KcSigner {
	ks := &KcSigner{
		apiKey:        key,
		apiSecret:     secret,
		apiPassPhrase: passPhrase,
	}
	ks.key = []byte(secret)
	return ks
}
