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

func (as *ApiService) Currencies() (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/currencies", nil)
	return as.call(req)
}

func (as *ApiService) Currency(currency string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/currencies/"+currency, nil)
	return as.call(req)
}
