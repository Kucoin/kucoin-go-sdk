package kucoin

import (
	"encoding/json"
	"testing"
)

func TestApiService_Currencies(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.Currencies()
	if err != nil {
		t.Fatal(err)
	}
	cl := CurrenciesModel{}
	rsp.ReadData(&cl)
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
	rsp, err := s.Currency("BTC")
	if err != nil {
		t.Fatal(err)
	}
	c := &CurrencyModel{}
	rsp.ReadData(c)
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
