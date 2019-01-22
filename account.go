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

func (as *ApiService) Accounts(currency, typ string) (AccountsModel, error) {
	p := map[string]string{}
	if currency != "" {
		p["currency"] = currency
	}
	if typ != "" {
		p["type"] = typ
	}
	req := NewRequest(http.MethodGet, "/api/v1/accounts", p)
	rsp, err := as.Call(req)
	if err != nil {
		return nil, err
	}
	v := AccountsModel{}
	rsp.ReadData(&v)
	return v, nil
}

func (as *ApiService) Account(accountId string) (*AccountModel, error) {
	req := NewRequest(http.MethodGet, "/api/v1/accounts/"+accountId, nil)
	rsp, err := as.Call(req)
	if err != nil {
		return nil, err
	}
	v := &AccountModel{}
	rsp.ReadData(v)
	return v, nil
}
