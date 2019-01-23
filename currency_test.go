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
	if err := rsp.ReadData(&cl); err != nil {
		t.Fatal(err)
	}
	for _, c := range cl {
		b, _ := json.Marshal(c)
		t.Log(string(b))
		switch {
		case c.Name == "":
			t.Error("Empty key 'name'")
		case c.Currency == "":
			t.Error("Empty key 'currency'")
		case c.FullName == "":
			t.Error("Empty key 'fullName'")
		case c.Precision == 0:
			t.Error("Empty key 'precision'")
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
	if err := rsp.ReadData(c); err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(c)
	t.Log(string(b))
	switch {
	case c.Name == "":
		t.Error("Empty key 'name'")
	case c.Currency == "":
		t.Error("Empty key 'currency'")
	case c.FullName == "":
		t.Error("Empty key 'fullName'")
	case c.Precision == 0:
		t.Error("Empty key 'precision'")
	case c.WithdrawalMinSize == "":
		t.Error("Empty key 'withdrawalMinSize'")
	case c.WithdrawalMinFee == "":
		t.Error("Empty key 'withdrawalMinFee'")
	}
}
