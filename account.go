package kucoin

import (
	"fmt"
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

type AccountHistoryModel struct {
	Currency  string            `json:"currency"`
	Amount    string            `json:"amount"`
	Fee       string            `json:"fee"`
	Balance   string            `json:"balance"`
	BizType   string            `json:"bizType"`
	Direction string            `json:"direction"`
	CreatedAt int64             `json:"createdAt"`
	Context   map[string]string `json:"context"`
}

type AccountHistoriesModel []AccountHistoryModel

func (as *ApiService) AccountHistories(accountId string, startAt, endAt int64) (*ApiResponse, error) {
	p := map[string]string{}
	if startAt > 0 {
		p["startAt"] = IntToString(startAt)
	}
	if endAt > 0 {
		p["endAt"] = IntToString(endAt)
	}
	req := NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/accounts/%s/ledgers", accountId), p)
	return as.Call(req)
}

type AccountHoldModel struct {
	Currency   string `json:"currency"`
	HoldAmount string `json:"holdAmount"`
	BizType    string `json:"bizType"`
	OrderId    string `json:"orderId"`
	CreatedAt  int64  `json:"createdAt"`
	UpdatedAt  int64  `json:"updatedAt"`
}

type AccountHoldsModel []AccountHoldModel

func (as *ApiService) AccountHolds(accountId string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/accounts/%s/holds", accountId), nil)
	return as.Call(req)
}

type InterTransferResultModel struct {
	OrderId string `json:"orderId"`
}

func (as *ApiService) InnerTransfer(clientOid, payAccountId, recAccountId, amount string) (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v1/accounts/inner-transfer", map[string]string{
		"clientOid":    clientOid,
		"payAccountId": payAccountId,
		"recAccountId": recAccountId,
		"amount":       amount,
	})
	return as.Call(req)
}
