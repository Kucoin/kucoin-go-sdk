package kucoin

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"time"
)

// Signer interface contains Sign() method.
type Signer interface {
	Sign(plain []byte) []byte
}

// Sha256Signer is the sha256 Signer.
type Sha256Signer struct {
	key []byte
}

// Sign makes a signature by sha256.
func (ss *Sha256Signer) Sign(plain []byte) []byte {
	hm := hmac.New(sha256.New, ss.key)
	hm.Write(plain)
	return hm.Sum(nil)
}

// KcSigner is the implement of Signer for KuCoin.
type KcSigner struct {
	Sha256Signer
	apiKey        string
	apiSecret     string
	apiPassPhrase string
	apiKeyVersion string
}

// Sign makes a signature by sha256 with `apiKey` `apiSecret` `apiPassPhrase`.
func (ks *KcSigner) Sign(plain []byte) []byte {
	s := ks.Sha256Signer.Sign(plain)
	return []byte(base64.StdEncoding.EncodeToString(s))
}

// Headers returns a map of signature header.
func (ks *KcSigner) Headers(plain string) map[string]string {
	t := IntToString(time.Now().UnixNano() / 1000000)
	p := []byte(t + plain)
	s := string(ks.Sign(p))
	ksHeaders := map[string]string{
		"KC-API-KEY":        ks.apiKey,
		"KC-API-PASSPHRASE": ks.apiPassPhrase,
		"KC-API-TIMESTAMP":  t,
		"KC-API-SIGN":       s,
	}

	if ks.apiKeyVersion != "" && ks.apiKeyVersion != ApiKeyVersionV1 {
		ksHeaders["KC-API-KEY-VERSION"] = ks.apiKeyVersion
	}

	return ksHeaders
}

// NewKcSigner creates a instance of KcSigner.
func NewKcSigner(key, secret, passPhrase string) *KcSigner {
	ks := &KcSigner{
		apiKey:        key,
		apiSecret:     secret,
		apiPassPhrase: passPhrase,
		apiKeyVersion: ApiKeyVersionV1,
	}
	ks.key = []byte(secret)
	return ks
}

// NewKcSignerV2 creates a instance of KcSigner.
func NewKcSignerV2(key, secret, passPhrase string) *KcSigner {
	ks := &KcSigner{
		apiKey:        key,
		apiSecret:     secret,
		apiPassPhrase: passPhraseEncrypt([]byte(secret), []byte(passPhrase)),
		apiKeyVersion: ApiKeyVersionV2,
	}
	ks.key = []byte(secret)
	return ks
}

// passPhraseEncrypt, encrypt passPhrase
func passPhraseEncrypt(key, plain []byte) string {
	hm := hmac.New(sha256.New, key)
	hm.Write(plain)
	return base64.StdEncoding.EncodeToString(hm.Sum(nil))
}
