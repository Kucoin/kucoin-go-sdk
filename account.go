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

func (as *ApiService) Accounts(v *AccountsModel, currency, typo string) (*ApiResponse, error) {
	p := map[string]string{}
	if currency != "" {
		p["currency"] = currency
	}
	if typo != "" {
		p["type"] = typo
	}
	req := NewRequest(http.MethodGet, "/api/v1/accounts", p)
	rsp, err := as.Call(req)
	if err != nil {
		return nil, err
	}
	if err := rsp.ReadData(v); err != nil {
		return rsp, err
	}
	return rsp, nil
}

func (as *ApiService) Account(v *AccountModel, accountId string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/accounts/"+accountId, nil)
	rsp, err := as.Call(req)
	if err != nil {
		return nil, err
	}
	if err := rsp.ReadData(v); err != nil {
		return rsp, err
	}
	return rsp, nil
}

func (as *ApiService) CreateAccount(v *AccountModel, typo, currency string) (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v1/accounts", map[string]string{"currency": currency, "type": typo})
	rsp, err := as.Call(req)
	if err != nil {
		return nil, err
	}
	if err := rsp.ReadData(v); err != nil {
		return rsp, err
	}
	return rsp, nil
}
