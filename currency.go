package kucoin

import (
	"context"
	"encoding/json"
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
func (as *ApiService) Currencies(ctx context.Context) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/currencies", nil)
	return as.Call(ctx, req)
}

// Currency returns the details of the currency.
// Deprecated: Use CurrencyV2 instead.
func (as *ApiService) Currency(ctx context.Context, currency string, chain string) (*ApiResponse, error) {
	params := map[string]string{}
	if chain != "" {
		params["chain"] = chain
	}
	req := NewRequest(http.MethodGet, "/api/v1/currencies/"+currency, params)
	return as.Call(ctx, req)
}

// ChainsModel Chains Model
type ChainsModel struct {
	ChainName         string `json:"chainName"`
	WithdrawalMinSize string `json:"withdrawalMinSize"`
	WithdrawalMinFee  string `json:"withdrawalMinFee"`
	IsWithdrawEnabled bool   `json:"isWithdrawEnabled"`
	IsDepositEnabled  bool   `json:"isDepositEnabled"`
	Confirms          int64  `json:"confirms"`
	ContractAddress   string `json:"contractAddress"`
	ChainId           string `json:"chainId"`
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
func (as *ApiService) CurrencyV2(ctx context.Context, currency string, chain string) (*ApiResponse, error) {
	params := map[string]string{}
	if chain != "" {
		params["chain"] = chain
	}
	req := NewRequest(http.MethodGet, "/api/v2/currencies/"+currency, params)
	return as.Call(ctx, req)
}

type PricesModel map[string]string

// Prices returns the fiat prices for currency.
func (as *ApiService) Prices(ctx context.Context, base, currencies string) (*ApiResponse, error) {
	params := map[string]string{}
	if base != "" {
		params["base"] = base
	}
	if currencies != "" {
		params["currencies"] = currencies
	}
	req := NewRequest(http.MethodGet, "/api/v1/prices", params)
	return as.Call(ctx, req)
}

type CurrenciesV3Model []*CurrencyV3Model

type CurrencyV3Model struct {
	Currency        string `json:"currency"`
	Name            string `json:"name"`
	FullName        string `json:"fullName"`
	Precision       int32  `json:"precision"`
	Confirms        int32  `json:"confirms"`
	ContractAddress string `json:"contractAddress"`
	IsMarginEnabled bool   `json:"isMarginEnabled"`
	IsDebitEnabled  bool   `json:"isDebitEnabled"`
	Chains          []struct {
		ChainName         string      `json:"chainName"`
		WithdrawalMinFee  json.Number `json:"withdrawalMinFee"`
		WithdrawalMinSize json.Number `json:"withdrawalMinSize"`
		WithdrawFeeRate   json.Number `json:"withdrawFeeRate"`
		DepositMinSize    json.Number `json:"depositMinSize"`
		IsWithdrawEnabled bool        `json:"isWithdrawEnabled"`
		IsDepositEnabled  bool        `json:"isDepositEnabled"`
		PreConfirms       int32       `json:"preConfirms"`
		ContractAddress   string      `json:"contractAddress"`
		ChainId           string      `json:"chainId"`
		Confirms          int32       `json:"confirms"`
	} `json:"chains"`
}

func (as *ApiService) CurrenciesV3(ctx context.Context) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v3/currencies/", nil)
	return as.Call(ctx, req)
}

// CurrencyInfoV3 Request via this endpoint to get the currency details of a specified currency
func (as *ApiService) CurrencyInfoV3(ctx context.Context, currency string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v3/currencies/"+currency, nil)
	return as.Call(ctx, req)
}
