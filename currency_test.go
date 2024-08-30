package kucoin

import (
	"context"
	"testing"
)

func TestApiService_Currencies(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.Currencies(context.Background())
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
	rsp, err := s.Currency(context.Background(), "BTC", "")
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

func TestApiService_Currency_V2(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.CurrencyV2(context.Background(), "BTC", "")
	if err != nil {
		t.Fatal(err)
	}
	c := &CurrencyV2Model{}
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
	}
}

func TestApiService_Prices(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.Prices(context.Background(), "USD", "BTC,KCS")
	if err != nil {
		t.Fatal(err)
	}
	p := PricesModel{}
	if err := rsp.ReadData(&p); err != nil {
		t.Fatal(err)
	}
	if len(p) == 0 {
		t.Error("Empty prices")
	}
	t.Log(ToJsonString(p))
}

func TestApiServiceCurrenciesV3(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.CurrenciesV3(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	p := CurrenciesV3Model{}
	if err := rsp.ReadData(&p); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(p))
}

func TestApiService_CurrencyInfoV3(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.CurrencyInfoV3(context.Background(), "BTC")
	if err != nil {
		t.Fatal(err)
	}
	p := CurrencyV3Model{}
	if err := rsp.ReadData(&p); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(p))
}
