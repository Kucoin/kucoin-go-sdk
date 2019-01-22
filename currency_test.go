package kucoin

import (
	"encoding/json"
	"testing"
)

func TestApiService_Currencies(t *testing.T) {
	s := NewApiServiceFromEnv()
	cl := CurrenciesModel{}
	if _, err := s.Currencies(&cl); err != nil {
		t.Fatal(err)
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

func TestApiService_Currency(t *testing.T) {
	s := NewApiServiceFromEnv()
	c := &CurrencyModel{}
	if _, err := s.Currency(c, "BTC"); err != nil {
		t.Fatal(err)
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
