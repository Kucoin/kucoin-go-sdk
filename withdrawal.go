package kucoin

import (
	"context"
	"net/http"
)

// A WithdrawalModel represents a withdrawal.
type WithdrawalModel struct {
	Chain      string `json:"chain"`
	Id         string `json:"id"`
	Address    string `json:"address"`
	Memo       string `json:"memo"`
	Currency   string `json:"currency"`
	Amount     string `json:"amount"`
	Fee        string `json:"fee"`
	WalletTxId string `json:"walletTxId"`
	IsInner    bool   `json:"isInner"`
	Status     string `json:"status"`
	Remark     string `json:"remark"`
	CreatedAt  int64  `json:"createdAt"`
	UpdatedAt  int64  `json:"updatedAt"`
}

// A WithdrawalsModel is the set of *WithdrawalModel.
type WithdrawalsModel []*WithdrawalModel

// Withdrawals returns a list of withdrawals.
func (as *ApiService) Withdrawals(ctx context.Context, params map[string]string, pagination *PaginationParam) (*ApiResponse, error) {
	pagination.ReadParam(params)
	req := NewRequest(http.MethodGet, "/api/v1/withdrawals", params)
	return as.Call(ctx, req)
}

// A V1WithdrawalModel represents a v1 historical withdrawal.
type V1WithdrawalModel struct {
	Address    string `json:"address"`
	Amount     string `json:"amount"`
	Currency   string `json:"currency"`
	IsInner    bool   `json:"isInner"`
	WalletTxId string `json:"walletTxId"`
	Status     string `json:"status"`
	CreateAt   int64  `json:"createAt"`
}

// A V1WithdrawalsModel is the set of *V1WithdrawalModel.
type V1WithdrawalsModel []*V1WithdrawalModel

// V1Withdrawals returns a list of v1 historical withdrawals.
func (as *ApiService) V1Withdrawals(ctx context.Context, params map[string]string, pagination *PaginationParam) (*ApiResponse, error) {
	pagination.ReadParam(params)
	req := NewRequest(http.MethodGet, "/api/v1/hist-withdrawals", params)
	return as.Call(ctx, req)
}

// A WithdrawalQuotasModel represents the quotas for a currency.
type WithdrawalQuotasModel struct {
	Currency            string `json:"currency"`
	AvailableAmount     string `json:"availableAmount"`
	RemainAmount        string `json:"remainAmount"`
	WithdrawMinSize     string `json:"withdrawMinSize"`
	LimitBTCAmount      string `json:"limitBTCAmount"`
	InnerWithdrawMinFee string `json:"innerWithdrawMinFee"`
	UsedBTCAmount       string `json:"usedBTCAmount"`
	IsWithdrawEnabled   bool   `json:"isWithdrawEnabled"`
	WithdrawMinFee      string `json:"withdrawMinFee"`
	Precision           uint8  `json:"precision"`
	Chain               string `json:"chain"`
}

// WithdrawalQuotas returns the quotas of withdrawal.
func (as *ApiService) WithdrawalQuotas(ctx context.Context, currency, chain string) (*ApiResponse, error) {
	params := map[string]string{"currency": currency}
	if chain != "" {
		params["chain"] = chain
	}
	req := NewRequest(http.MethodGet, "/api/v1/withdrawals/quotas", params)
	return as.Call(ctx, req)
}

// ApplyWithdrawalResultModel represents the result of ApplyWithdrawal().
type ApplyWithdrawalResultModel struct {
	WithdrawalId string `json:"withdrawalId"`
}

// ApplyWithdrawal applies a withdrawal.
func (as *ApiService) ApplyWithdrawal(ctx context.Context, currency, address, amount string, options map[string]string) (*ApiResponse, error) {
	p := map[string]string{
		"currency": currency,
		"address":  address,
		"amount":   amount,
	}
	for k, v := range options {
		p[k] = v
	}
	req := NewRequest(http.MethodPost, "/api/v1/withdrawals", p)
	return as.Call(ctx, req)
}

// CancelWithdrawalResultModel represents the result of CancelWithdrawal().
type CancelWithdrawalResultModel struct {
	CancelledWithdrawIds []string `json:"cancelledWithdrawIds"`
}

// CancelWithdrawal cancels a withdrawal by withdrawalId.
func (as *ApiService) CancelWithdrawal(ctx context.Context, withdrawalId string) (*ApiResponse, error) {
	req := NewRequest(http.MethodDelete, "/api/v1/withdrawals/"+withdrawalId, nil)
	return as.Call(ctx, req)
}
