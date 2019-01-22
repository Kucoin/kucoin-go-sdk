package kucoin

import (
	"net/http"
)

type CurrencyModel struct {
	Name              string `json:"name"`
	Currency          string `json:"currency"`
	FullName          string `json:"fullName"`
	Precision         uint8  `json:"precision"`
	WithdrawalMinSize string `json:"withdrawalMinSize"`
	WithdrawalMinFee  string `json:"withdrawalMinFee"`
	IsWithdrawEnabled bool   `json:"isWithdrawEnabled"`
	IsDepositEnabled  bool   `json:"isDepositEnabled"`
}

type CurrenciesModel []*CurrencyModel

func (as *ApiService) CurrencyList() (CurrenciesModel, error) {
	req := NewRequest(http.MethodGet, "/api/v1/currencies", nil)
	rsp, err := as.Call(req)
	if err != nil {
		return nil, err
	}
	type Data struct {
		Data CurrenciesModel `json:"data"`
	}
	v := &Data{}
	rsp.ApiData(v)
	return v.Data, nil
}

func (as *ApiService) CurrencyDetail(currency string) (*CurrencyModel, error) {
	req := NewRequest(http.MethodGet, "/api/v1/currencies/"+currency, nil)
	rsp, err := as.Call(req)
	if err != nil {
		return nil, err
	}
	type Data struct {
		Data *CurrencyModel `json:"data"`
	}
	v := &Data{}
	rsp.ApiData(v)
	return v.Data, nil
}
