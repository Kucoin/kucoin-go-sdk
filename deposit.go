package kucoin

import (
	"net/http"
)

// A DepositAddressModel represents a deposit address of currency for deposit.
type DepositAddressModel struct {
	Address string `json:"address"`
	Memo    string `json:"memo"`
}

// A DepositAddressesModel is the set of *DepositAddressModel.
type DepositAddressesModel []*DepositAddressModel

// A DepositModel represents a deposit record.
type DepositModel struct {
	Address    string `json:"address"`
	Memo       string `json:"memo"`
	Amount     string `json:"amount"`
	Fee        string `json:"fee"`
	Currency   string `json:"currency"`
	IsInner    bool   `json:"isInner"`
	WalletTxId string `json:"walletTxId"`
	Status     string `json:"status"`
	CreatedAt  int64  `json:"createdAt"`
	UpdatedAt  int64  `json:"updatedAt"`
}

// A DepositsModel is the set of *DepositModel.
type DepositsModel []*DepositModel

// CreateDepositAddress creates a deposit address.
func (as *ApiService) CreateDepositAddress(currency string) (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v1/deposit-addresses", map[string]string{"currency": currency})
	return as.Call(req)
}

// DepositAddresses returns the deposit address of currency for deposit.
// If return data is empty, you may need create a deposit address first.
func (as *ApiService) DepositAddresses(currency string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/deposit-addresses", map[string]string{"currency": currency})
	return as.Call(req)
}

// Deposits returns a list of deposit.
func (as *ApiService) Deposits(currency, status string, startAt, endAt int64, pagination *PaginationParam) (*ApiResponse, error) {
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
	req := NewRequest(http.MethodGet, "/api/v1/deposits", p)
	return as.Call(req)
}
