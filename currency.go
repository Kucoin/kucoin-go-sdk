package kucoin

import (
	"net/http"
)

type Currency struct {
	Name              string `json:"name"`
	Currency          string `json:"currency"`
	FullName          string `json:"fullName"`
	Precision         uint8  `json:"precision"`
	WithdrawalMinSize string `json:"withdrawalMinSize"`
	WithdrawalMinFee  string `json:"withdrawalMinFee"`
	IsWithdrawEnabled bool   `json:"isWithdrawEnabled"`
	IsDepositEnabled  bool   `json:"isDepositEnabled"`
}

type Currencies []*Currency

func CurrencyList() (Currencies, error) {
	req := NewRequest(http.MethodGet, "/api/v1/currencies", nil)
	rsp, err := Api.Call(req)
	if err != nil {
		return nil, err
	}
	type Data struct {
		Data Currencies `json:"data"`
	}
	v := &Data{}
	rsp.ApiData(v)
	return v.Data, nil
}

func CurrencyDetail(currency string) (*Currency, error) {
	req := NewRequest(http.MethodGet, "/api/v1/currencies/"+currency, nil)
	rsp, err := Api.Call(req)
	if err != nil {
		return nil, err
	}
	type Data struct {
		Data *Currency `json:"data"`
	}
	v := &Data{}
	rsp.ApiData(v)
	return v.Data, nil
}
