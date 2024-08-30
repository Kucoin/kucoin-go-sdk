package kucoin

import (
	"context"
	"net/http"
)

// HfAccountInnerTransfer Users can transfer funds between their main account,
// trading account, and high-frequency trading account free of charge.
func (as *ApiService) HfAccountInnerTransfer(ctx context.Context, params map[string]string) (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v2/accounts/inner-transfer", params)
	return as.Call(ctx, req)
}

type HfAccountInnerTransferRes struct {
	OrderId string `json:"orderId"`
}

// HfAccounts Get a list of high-frequency trading accounts.
func (as *ApiService) HfAccounts(ctx context.Context, currency, accountType string) (*ApiResponse, error) {
	p := map[string]string{
		"currency": currency,
		"type":     accountType,
	}
	req := NewRequest(http.MethodGet, "/api/v1/accounts", p)
	return as.Call(ctx, req)
}

type HfAccountsModel []HfAccountModel

// HfAccount Get the details of the high-frequency trading account
func (as *ApiService) HfAccount(ctx context.Context, accountId string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/accounts/"+accountId, nil)
	return as.Call(ctx, req)
}

type HfAccountModel struct {
	Balance   string `json:"balance"`
	Available string `json:"available"`
	Currency  string `json:"currency"`
	Holds     string `json:"holds"`
	Type      string `json:"type"`
	Id        string `json:"id"`
}

// HfAccountTransferable This API can be used to obtain the amount of transferrable funds
// in high-frequency trading accounts.
func (as *ApiService) HfAccountTransferable(ctx context.Context, currency string) (*ApiResponse, error) {
	p := map[string]string{
		"currency": currency,
		"type":     "TRADE_HF",
	}
	req := NewRequest(http.MethodGet, "/api/v1/accounts/transferable", p)
	return as.Call(ctx, req)
}

type HfAccountTransferableModel struct {
	Balance      string `json:"balance"`
	Available    string `json:"available"`
	Currency     string `json:"currency"`
	Holds        string `json:"holds"`
	Transferable string `json:"transferable"`
}

// HfAccountLedgers  returns all transfer (in and out) records in high-frequency trading account
// and supports multi-coin queries. The query results are sorted in descending order by createdAt and id.
func (as *ApiService) HfAccountLedgers(ctx context.Context, params map[string]string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/hf/accounts/ledgers", params)
	return as.Call(ctx, req)
}

type HfAccountLedgersModel []*HfAccountLedgerModel

type HfAccountLedgerModel struct {
	Id          string `json:"id"`
	Currency    string `json:"currency"`
	Amount      string `json:"amount"`
	Fee         string `json:"fee"`
	Balance     string `json:"balance"`
	AccountType string `json:"accountType"`
	BizType     string `json:"bizType"`
	Direction   string `json:"direction"`
	CreatedAt   string `json:"createdAt"`
	Context     string `json:"context"`
}
