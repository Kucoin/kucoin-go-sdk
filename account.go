package kucoin

import (
	"net/http"
)

type AccountModel struct {
	Id        string `json:"id"`
	Currency  string `json:"currency"`
	Type      string `json:"type"`
	Balance   string `json:"balance"`
	Available string `json:"available"`
	Holds     string `json:"holds"`
}

type AccountsModel []*AccountModel

func (as *ApiService) Accounts(currency, typo string) (*ApiResponse, error) {
	p := map[string]string{}
	if currency != "" {
		p["currency"] = currency
	}
	if typo != "" {
		p["type"] = typo
	}
	req := NewRequest(http.MethodGet, "/api/v1/accounts", p)
	return as.Call(req)
}

func (as *ApiService) Account(accountId string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/accounts/"+accountId, nil)
	return as.Call(req)
}

func (as *ApiService) CreateAccount(typo, currency string) (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v1/accounts", map[string]string{"currency": currency, "type": typo})
	return as.Call(req)
}
