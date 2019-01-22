package kucoin

import (
	"net/http"
)

type AccountModel struct {
	Id        string `json:"name"`
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
	type Data struct {
		Data AccountsModel `json:"data"`
	}
	v := &Data{}
	rsp.ApiData(v)
	return v.Data, nil
}
