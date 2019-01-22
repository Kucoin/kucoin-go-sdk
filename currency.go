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

func (as *ApiService) Currencies(v *CurrenciesModel) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/currencies", nil)
	rsp, err := as.Call(req)
	if err != nil {
		return nil, err
	}
	if err := rsp.ReadData(v); err != nil {
		return rsp, err
	}
	return rsp, nil
}

func (as *ApiService) Currency(v *CurrencyModel, currency string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/currencies/"+currency, nil)
	rsp, err := as.Call(req)
	if err != nil {
		return nil, err
	}
	if err := rsp.ReadData(v); err != nil {
		return rsp, err
	}
	return rsp, nil
}
