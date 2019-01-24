package kucoin

import (
	"net/http"
)

type WithdrawalModel struct {
	Id         string `json:"id"`
	Address    string `json:"address"`
	Memo       string `json:"memo"`
	Currency   string `json:"currency"`
	Amount     string `json:"amount"`
	Fee        string `json:"fee"`
	WalletTxId string `json:"walletTxId"`
	IsInner    bool   `json:"isInner"`
	Status     string `json:"status"`
	CreatedAt  int64  `json:"createdAt"`
	UpdatedAt  int64  `json:"updatedAt"`
}

type WithdrawalsModel []WithdrawalModel

func (as *ApiService) Withdrawals(currency, status string, startAt, endAt int64, pagination *PaginationParam) (*ApiResponse, error) {
	p := map[string]string{}
	if currency != "" {
		p["currency"] = currency
	}
	if status != "" {
		p["status"] = status
	}
	if startAt > 0 {
		p["startAt"] = IntToString(startAt)
	}
	if endAt > 0 {
		p["endAt"] = IntToString(endAt)
	}
	pagination.ReadParam(p)
	req := NewRequest(http.MethodGet, "/api/v1/withdrawals", p)
	return as.call(req)
}

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
}

func (as *ApiService) WithdrawalQuotas(currency string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/withdrawals/quotas", map[string]string{"currency": currency})
	return as.call(req)
}

type ApplyWithdrawalResultModel struct {
	WithdrawalId string `json:"withdrawalId"`
}

func (as *ApiService) ApplyWithdrawal(currency, address, amount string, options map[string]string) (*ApiResponse, error) {
	p := map[string]string{
		"currency": currency,
		"address":  address,
		"amount":   amount,
	}
	for k, v := range options {
		p[k] = v
	}
	req := NewRequest(http.MethodPost, "/api/v1/withdrawals", p)
	return as.call(req)
}

type CancelWithdrawalResultModel struct {
	CancelledWithdrawIds []string `json:"cancelledWithdrawIds"`
}

func (as *ApiService) CancelWithdrawal(withdrawalId string) (*ApiResponse, error) {
	req := NewRequest(http.MethodDelete, "/api/v1/withdrawals/"+withdrawalId, nil)
	return as.call(req)
}
