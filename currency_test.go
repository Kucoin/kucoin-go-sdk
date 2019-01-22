package kucoin

import (
	"encoding/json"
	"testing"
)

func TestCurrencyList(t *testing.T) {
	s := NewApiServiceFromEnv()
	cl, err := s.Currencies()
	if err != nil {
		t.Error(err)
	}
	for _, c := range cl {
		b, _ := json.Marshal(c)
		t.Log(string(b))
		switch {
		case c.Name == "":
			t.Error("Missing key 'name'")
		case c.Currency == "":
			t.Error("Missing key 'currency'")
		case c.FullName == "":
			t.Error("Missing key 'fullName'")
		case c.Precision == 0:
			t.Error("Missing key 'precision'")
		}
	}
}

func TestCurrencyDetail(t *testing.T) {
	s := NewApiServiceFromEnv()
	c, err := s.Currency("BTC")
	if err != nil {
		t.Error(err)
	}
	b, _ := json.Marshal(c)
	t.Log(string(b))
	switch {
	case c.Name == "":
		t.Error("Missing key 'name'")
	case c.Currency == "":
		t.Error("Missing key 'currency'")
	case c.FullName == "":
		t.Error("Missing key 'fullName'")
	case c.Precision == 0:
		t.Error("Missing key 'precision'")
	case c.WithdrawalMinSize == "":
		t.Error("Missing key 'withdrawalMinSize'")
	case c.WithdrawalMinFee == "":
		t.Error("Missing key 'withdrawalMinFee'")
	}
}
