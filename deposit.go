package kucoin

import (
	"net/http"
)

// A DepositAddressModel represents a deposit address of currency for deposit.
type DepositAddressModel struct {
	Address string `json:"address"`
	Memo    string `json:"memo"`
	Chain   string `json:"chain"`
}

// A DepositAddressesModel is the set of *DepositAddressModel.
type DepositAddressesModel DepositAddressModel

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
	Remark     string `json:"remark"`
	CreatedAt  int64  `json:"createdAt"`
	UpdatedAt  int64  `json:"updatedAt"`
}

// A DepositsModel is the set of *DepositModel.
type DepositsModel []*DepositModel

// CreateDepositAddress creates a deposit address.
func (as *ApiService) CreateDepositAddress(currency, chain string) (*ApiResponse, error) {
	params := map[string]string{"currency": currency}
	if chain != "" {
		params["chain"] = chain
	}
	req := NewRequest(http.MethodPost, "/api/v1/deposit-addresses", params)
	return as.Call(req)
}

// DepositAddresses returns the deposit address of currency for deposit.
// If return data is empty, you may need create a deposit address first.
func (as *ApiService) DepositAddresses(currency, chain string) (*ApiResponse, error) {
	params := map[string]string{"currency": currency}
	if chain != "" {
		params["chain"] = chain
	}
	req := NewRequest(http.MethodGet, "/api/v1/deposit-addresses", params)
	return as.Call(req)
}

type depositAddressV2Model struct {
	Address         string `json:"address"`
	Memo            string `json:"memo"`
	Chain           string `json:"chain"`
	ContractAddress string `json:"contract_address"`
}
type DepositAddressesV2Model []*depositAddressV2Model

// DepositAddressesV2 Get all deposit addresses for the currency you intend to deposit.
// If the returned data is empty, you may need to create a deposit address first.
func (as *ApiService) DepositAddressesV2(currency string) (*ApiResponse, error) {
	params := map[string]string{"currency": currency}
	req := NewRequest(http.MethodGet, "/api/v2/deposit-addresses", params)
	return as.Call(req)
}

// Deposits returns a list of deposit.
func (as *ApiService) Deposits(params map[string]string, pagination *PaginationParam) (*ApiResponse, error) {
	pagination.ReadParam(params)
	req := NewRequest(http.MethodGet, "/api/v1/deposits", params)
	return as.Call(req)
}

// A V1DepositModel represents a v1 deposit record.
type V1DepositModel struct {
	Amount     string `json:"amount"`
	Currency   string `json:"currency"`
	IsInner    bool   `json:"isInner"`
	WalletTxId string `json:"walletTxId"`
	Status     string `json:"status"`
	CreateAt   int64  `json:"createAt"`
}

// A V1DepositsModel is the set of *V1DepositModel.
type V1DepositsModel []*V1DepositModel

// V1Deposits returns a list of v1 historical deposits.
func (as *ApiService) V1Deposits(params map[string]string, pagination *PaginationParam) (*ApiResponse, error) {
	pagination.ReadParam(params)
	req := NewRequest(http.MethodGet, "/api/v1/hist-deposits", params)
	return as.Call(req)
}
