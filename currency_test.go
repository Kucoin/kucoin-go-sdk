package kucoin

import (
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
		t.Log(ToJsonString(c))
		switch {
		case c.Name == "":
			t.Error("Empty key 'name'")
		case c.Currency == "":
			t.Error("Empty key 'currency'")
		case c.FullName == "":
			t.Error("Empty key 'fullName'")
		case c.WithdrawalMinSize == "":
			t.Error("Empty key 'withdrawalMinSize'")
		case c.WithdrawalMinFee == "":
			t.Error("Empty key 'withdrawalMinFee'")
		}
	}
}

func TestApiService_Currency(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.Currency("BTC", "")
	if err != nil {
		t.Fatal(err)
	}
	c := &CurrencyModel{}
	if err := rsp.ReadData(c); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(c))
	switch {
	case c.Name == "":
		t.Error("Empty key 'name'")
	case c.Currency == "":
		t.Error("Empty key 'currency'")
	case c.FullName == "":
		t.Error("Empty key 'fullName'")
	case c.WithdrawalMinSize == "":
		t.Error("Empty key 'withdrawalMinSize'")
	case c.WithdrawalMinFee == "":
		t.Error("Empty key 'withdrawalMinFee'")
	}
}

func TestApiService_Prices(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.Prices("USD", "BTC,KCS")
	if err != nil {
		t.Fatal(err)
	}
	p := map[string]string{}
	if err := rsp.ReadData(&p); err != nil {
		t.Fatal(err)
	}
	if len(p) == 0 {
		t.Error("Empty prices")
	}
	t.Log(ToJsonString(p))
}
