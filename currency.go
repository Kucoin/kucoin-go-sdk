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

func (as *ApiService) Currencies() (CurrenciesModel, error) {
	req := NewRequest(http.MethodGet, "/api/v1/currencies", nil)
	rsp, err := as.Call(req)
	if err != nil {
		return nil, err
	}
	v := CurrenciesModel{}
	rsp.ReadData(&v)
	return v, nil
}

func (as *ApiService) Currency(currency string) (*CurrencyModel, error) {
	req := NewRequest(http.MethodGet, "/api/v1/currencies/"+currency, nil)
	rsp, err := as.Call(req)
	if err != nil {
		return nil, err
	}
	type Data struct {
		Data *CurrencyModel `json:"data"`
	}
	v := &CurrencyModel{}
	rsp.ReadData(v)
	return v, nil
}
