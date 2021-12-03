package kucoin

import (
	"net/http"
)

// A CurrencyModel represents a model of known currency.
type CurrencyModel struct {
	Name              string `json:"name"`
	Currency          string `json:"currency"`
	FullName          string `json:"fullName"`
	Precision         uint8  `json:"precision"`
	Confirms          int64  `json:"confirms"`
	ContractAddress   string `json:"contractAddress"`
	WithdrawalMinSize string `json:"withdrawalMinSize"`
	WithdrawalMinFee  string `json:"withdrawalMinFee"`
	IsWithdrawEnabled bool   `json:"isWithdrawEnabled"`
	IsDepositEnabled  bool   `json:"isDepositEnabled"`
	IsMarginEnabled   bool   `json:"isMarginEnabled"`
	IsDebitEnabled    bool   `json:"isDebitEnabled"`
}

// A CurrenciesModel is the set of *CurrencyModel.
type CurrenciesModel []*CurrencyModel

// Currencies returns a list of known currencies.
func (as *ApiService) Currencies() (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/currencies", nil)
	return as.Call(req)
}

// Currency returns the details of the currency.
// Deprecated: Use CurrencyV2 instead.
func (as *ApiService) Currency(currency string, chain string) (*ApiResponse, error) {
	params := map[string]string{}
	if chain != "" {
		params["chain"] = chain
	}
	req := NewRequest(http.MethodGet, "/api/v1/currencies/"+currency, params)
	return as.Call(req)
}

// ChainsModel Chains Model
type ChainsModel struct {
	ChainName         string `json:"chainName"`
	WithdrawalMinSize string `json:"withdrawalMinSize"`
	WithdrawalMinFee  string `json:"withdrawalMinFee"`
	IsWithdrawEnabled bool   `json:"isWithdrawEnabled"`
	IsDepositEnabled  bool   `json:"isDepositEnabled"`
	Confirms          int64  `json:"confirms"`
	ContractAddress   string   `json:"contractAddress"`
}

// CurrencyV2Model CurrencyV2 Model
type CurrencyV2Model struct {
	Name            string         `json:"name"`
	Currency        string         `json:"currency"`
	FullName        string         `json:"fullName"`
	Precision       uint8          `json:"precision"`
	Confirms        int64          `json:"confirms"`
	ContractAddress string         `json:"contractAddress"`
	IsMarginEnabled bool           `json:"isMarginEnabled"`
	IsDebitEnabled  bool           `json:"isDebitEnabled"`
	Chains          []*ChainsModel `json:"chains"`
}

// CurrencyV2 returns the details of the currency.
func (as *ApiService) CurrencyV2(currency string, chain string) (*ApiResponse, error) {
	params := map[string]string{}
	if chain != "" {
		params["chain"] = chain
	}
	req := NewRequest(http.MethodGet, "/api/v2/currencies/"+currency, params)
	return as.Call(req)
}

type PricesModel map[string]string

// Prices returns the fiat prices for currency.
func (as *ApiService) Prices(base, currencies string) (*ApiResponse, error) {
	params := map[string]string{}
	if base != "" {
		params["base"] = base
	}
	if currencies != "" {
		params["currencies"] = currencies
	}
	req := NewRequest(http.MethodGet, "/api/v1/prices", params)
	return as.Call(req)
}
