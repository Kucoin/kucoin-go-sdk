package kucoin

import (
	"net/http"
	"strconv"
)

type DepositAddressModel struct {
	Address string `json:"address"`
	Memo    string `json:"memo"`
}

type DepositAddressesModel []DepositAddressModel

type DepositModel struct {
	Address    string  `json:"address"`
	Memo       string  `json:"memo"`
	Amount     int64   `json:"amount"`
	Fee        float32 `json:"fee"`
	Currency   string  `json:"currency"`
	IsInner    bool    `json:"isInner"`
	WalletTxId string  `json:"walletTxId"`
	Status     string  `json:"status"`
	CreatedAt  int64   `json:"createdAt"`
	UpdatedAt  int64   `json:"updatedAt"`
}

type DepositsModel []DepositModel

func (as *ApiService) CreateDepositAddress(currency string) (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v1/deposit-addresses", map[string]string{"currency": currency})
	return as.Call(req)
}

func (as *ApiService) DepositAddresses(currency string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/deposit-addresses", map[string]string{"currency": currency})
	return as.Call(req)
}

func (as *ApiService) Deposits(currency, status string, startAt, endAt int64) (*ApiResponse, error) {
	p := map[string]string{}
	if currency != "" {
		p["currency"] = currency
	}
	if status != "" {
		p["status"] = status
	}
	if startAt > 0 {
		p["startAt"] = strconv.FormatInt(startAt, 10)
	}
	if endAt > 0 {
		p["endAt"] = strconv.FormatInt(endAt, 10)
	}
	req := NewRequest(http.MethodGet, "/api/v1/deposits", p)
	return as.Call(req)
}
