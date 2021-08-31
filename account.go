package kucoin

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// An AccountModel represents an account.
type AccountModel struct {
	Id        string `json:"id"`
	Currency  string `json:"currency"`
	Type      string `json:"type"`
	Balance   string `json:"balance"`
	Available string `json:"available"`
	Holds     string `json:"holds"`
}

// An AccountsModel is the set of *AccountModel.
type AccountsModel []*AccountModel

// Accounts returns a list of accounts.
// See the Deposits section for documentation on how to deposit funds to begin trading.
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

// Account returns an account when you know the accountId.
func (as *ApiService) Account(accountId string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/accounts/"+accountId, nil)
	return as.Call(req)
}

// A SubAccountUserModel represents a sub-account user.
type SubAccountUserModel struct {
	UserId  string `json:"userId"`
	SubName string `json:"subName"`
	Remarks string `json:"remarks"`
	Type    int    `json:"type"`
}

// A SubAccountUsersModel is the set of *SubAccountUserModel.
type SubAccountUsersModel []*SubAccountUserModel

// SubAccountUsers returns a list of sub-account user.
func (as *ApiService) SubAccountUsers() (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/sub/user", nil)
	return as.Call(req)
}

// A SubAccountsModel is the set of *SubAccountModel.
type SubAccountsModel []*SubAccountModel

// SubAccounts returns the aggregated balance of all sub-accounts of the current user.
func (as *ApiService) SubAccounts() (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/sub-accounts", nil)
	return as.Call(req)
}

// AccountsTransferableModel  RESPONSES of AccountsTransferable
type AccountsTransferableModel struct {
	Currency     string `json:"currency"`
	Balance      string `json:"balance"`
	Available    string `json:"available"`
	Holds        string `json:"holds"`
	Transferable string `json:"transferable"`
}

// AccountsTransferable  returns the transferable balance of a specified account.
func (as *ApiService) AccountsTransferable(currency, typo string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/accounts/transferable", map[string]string{"currency": currency, "type": typo})
	return as.Call(req)
}

// A SubAccountModel represents the balance of a sub-account user.
type SubAccountModel struct {
	SubUserId    string `json:"subUserId"`
	SubName      string `json:"subName"`
	MainAccounts []struct {
		Currency          string `json:"currency"`
		Balance           string `json:"balance"`
		Available         string `json:"available"`
		Holds             string `json:"holds"`
		BaseCurrency      string `json:"baseCurrency"`
		BaseCurrencyPrice string `json:"baseCurrencyPrice"`
		BaseAmount        string `json:"baseAmount"`
	} `json:"mainAccounts"`
	TradeAccounts []struct {
		Currency          string `json:"currency"`
		Balance           string `json:"balance"`
		Available         string `json:"available"`
		Holds             string `json:"holds"`
		BaseCurrency      string `json:"baseCurrency"`
		BaseCurrencyPrice string `json:"baseCurrencyPrice"`
		BaseAmount        string `json:"baseAmount"`
	} `json:"tradeAccounts"`
	MarginAccounts []struct {
		Currency          string `json:"currency"`
		Balance           string `json:"balance"`
		Available         string `json:"available"`
		Holds             string `json:"holds"`
		BaseCurrency      string `json:"baseCurrency"`
		BaseCurrencyPrice string `json:"baseCurrencyPrice"`
		BaseAmount        string `json:"baseAmount"`
	} `json:"marginAccounts"`
}

// SubAccount returns the detail of a sub-account.
func (as *ApiService) SubAccount(subUserId string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/sub-accounts/"+subUserId, nil)
	return as.Call(req)
}

// CreateAccountModel represents The account id returned from creating an account
type CreateAccountModel struct {
	Id string `json:"id"`
}

// CreateAccount creates an account according to type(main|trade) and currency
// Parameter #1 typo is type of account.
func (as *ApiService) CreateAccount(typo, currency string) (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v1/accounts", map[string]string{"currency": currency, "type": typo})
	return as.Call(req)
}

// An AccountLedgerModel represents account activity either increases or decreases your account balance.
type AccountLedgerModel struct {
	ID          string          `json:"id"`
	Currency    string          `json:"currency"`
	Amount      string          `json:"amount"`
	Fee         string          `json:"fee"`
	Balance     string          `json:"balance"`
	AccountType string          `json:"accountType"`
	BizType     string          `json:"bizType"`
	Direction   string          `json:"direction"`
	CreatedAt   int64           `json:"createdAt"`
	Context     json.RawMessage `json:"context"`
}

// An AccountLedgersModel the set of *AccountLedgerModel.
type AccountLedgersModel []*AccountLedgerModel

// AccountLedgers returns a list of ledgers.
// Deprecated: This interface was discontinued on Nov 05, 2020. Please use AccountLedgersV2.
// Account activity either increases or decreases your account balance.
// Items are paginated and sorted latest first.
// Deprecated
func (as *ApiService) AccountLedgers(accountId string, startAt, endAt int64, options map[string]string, pagination *PaginationParam) (*ApiResponse, error) {
	p := map[string]string{}
	if startAt > 0 {
		p["startAt"] = IntToString(startAt)
	}
	if endAt > 0 {
		p["endAt"] = IntToString(endAt)
	}
	for k, v := range options {
		p[k] = v
	}
	pagination.ReadParam(p)
	req := NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/accounts/%s/ledgers", accountId), p)
	return as.Call(req)
}

// AccountLedgersV2 returns a list of ledgers.
// Recommended for use on Nov 05, 2020.
// Account activity either increases or decreases your account balance.
// Items are paginated and sorted latest first.
func (as *ApiService) AccountLedgersV2(params map[string]string, pagination *PaginationParam) (*ApiResponse, error) {
	pagination.ReadParam(params)

	req := NewRequest(http.MethodGet, "/api/v1/accounts/ledgers", params)
	return as.Call(req)
}

// An AccountHoldModel represents the holds on an account for any active orders or pending withdraw requests.
// As an order is filled, the hold amount is updated.
// If an order is canceled, any remaining hold is removed.
// For a withdraw, once it is completed, the hold is removed.
type AccountHoldModel struct {
	Currency   string `json:"currency"`
	HoldAmount string `json:"holdAmount"`
	BizType    string `json:"bizType"`
	OrderId    string `json:"orderId"`
	CreatedAt  int64  `json:"createdAt"`
	UpdatedAt  int64  `json:"updatedAt"`
}

// An AccountHoldsModel is the set of *AccountHoldModel.
type AccountHoldsModel []*AccountHoldModel

// AccountHolds returns a list of currency hold.
// Holds are placed on an account for any active orders or pending withdraw requests.
// As an order is filled, the hold amount is updated.
// If an order is canceled, any remaining hold is removed.
// For a withdraw, once it is completed, the hold is removed.
func (as *ApiService) AccountHolds(accountId string, pagination *PaginationParam) (*ApiResponse, error) {
	p := map[string]string{}
	pagination.ReadParam(p)
	req := NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/accounts/%s/holds", accountId), p)
	return as.Call(req)
}

// An InnerTransferResultModel represents the result of a inner-transfer operation.
type InnerTransferResultModel struct {
	OrderId string `json:"orderId"`
}

// InnerTransferV2 makes a currency transfer internally.
// Recommended for use on June 5, 2019.
// The inner transfer interface is used for transferring assets between the accounts of a user and is free of charges.
// For example, a user could transfer assets from their main account to their trading account on the platform.
func (as *ApiService) InnerTransferV2(clientOid, currency, from, to, amount string) (*ApiResponse, error) {
	p := map[string]string{
		"clientOid": clientOid,
		"currency":  currency,
		"from":      from,
		"to":        to,
		"amount":    amount,
	}
	req := NewRequest(http.MethodPost, "/api/v2/accounts/inner-transfer", p)
	return as.Call(req)
}

// A SubTransferResultModel represents the result of a sub-transfer operation.
type SubTransferResultModel InnerTransferResultModel

// SubTransfer transfers between master account and sub-account.
// Deprecated: This interface was discontinued on Oct 28, 2020. Please use SubTransferV2.
func (as *ApiService) SubTransfer(params map[string]string) (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v1/accounts/sub-transfer", params)
	return as.Call(req)
}

// SubTransferV2 transfers between master account and sub-account.
// Recommended for use on Oct 28, 2020.
func (as *ApiService) SubTransferV2(params map[string]string) (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v2/accounts/sub-transfer", params)
	return as.Call(req)
}

// BaseFeeModel RESPONSES of BaseFee endpoint
type BaseFeeModel struct {
	TakerFeeRate string `json:"takerFeeRate"`
	MakerFeeRate string `json:"makerFeeRate"`
}

// BaseFee returns the basic fee rate of users.
func (as *ApiService) BaseFee() (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/base-fee", nil)
	return as.Call(req)
}

type TradeFeesResultModel []struct {
	Symbol       string `json:"symbol"`
	TakerFeeRate string `json:"takerFeeRate"`
	MakerFeeRate string `json:"makerFeeRate"`
}

// ActualFee returns the actual fee rate of the trading pair.
// You can inquire about fee rates of 10 trading pairs each time at most.
func (as *ApiService) ActualFee(symbols string) (*ApiResponse, error) {
	p := map[string]string{
		"symbols": symbols,
	}
	req := NewRequest(http.MethodGet, "/api/v1/trade-fees", p)
	return as.Call(req)
}
