package kucoin

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// MarkPriceModel represents mark price of a symbol
type MarkPriceModel struct {
	Symbol      string      `json:"symbol"`
	Granularity json.Number `json:"granularity"`
	TimePoint   json.Number `json:"timePoint"`
	Value       json.Number `json:"value"`
}

// CurrentMarkPrice returns current mark price of the input symbol
func (as *ApiService) CurrentMarkPrice(symbol string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/mark-price/%s/current", symbol), nil)
	return as.Call(req)
}

// MarginConfigModel represents a margin configuration
type MarginConfigModel struct {
	CurrencyList     []string    `json:"currencyList"`
	WarningDebtRatio json.Number `json:"warningDebtRatio"`
	LiqDebtRatio     json.Number `json:"liqDebtRatio"`
	MaxLeverage      json.Number `json:"maxLeverage"`
}

// MarginConfig returns a margin configuration
func (as *ApiService) MarginConfig() (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/margin/config", nil)
	return as.Call(req)
}

// MarginAccountModel represents a margin account information
type MarginAccountModel struct {
	Accounts []struct {
		AvailableBalance json.Number `json:"availableBalance"`
		Currency         string      `json:"currency"`
		HoldBalance      json.Number `json:"holdBalance"`
		Liability        json.Number `json:"liability"`
		MaxBorrowSize    json.Number `json:"maxBorrowSize"`
		TotalBalance     json.Number `json:"totalBalance"`
	} `json:"accounts"`
	DebtRatio json.Number `json:"debtRatio"`
}

// MarginAccount returns a margin account information
func (as *ApiService) MarginAccount() (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/margin/account", nil)
	return as.Call(req)
}

// CreateBorrowOrderResultModel represents the result of create a borrow order
type CreateBorrowOrderResultModel struct {
	OrderId  string `json:"orderId"`
	Currency string `json:"currency"`
}

// CreateBorrowOrder returns the result of create a borrow order
func (as *ApiService) CreateBorrowOrder(params map[string]string) (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v1/margin/borrow", params)
	return as.Call(req)
}

// BorrowOrderModel represents a borrow order
type BorrowOrderModel struct {
	OrderId   string      `json:"orderId"`
	Currency  string      `json:"currency"`
	Size      json.Number `json:"size"`
	Filled    json.Number `json:"filled"`
	Status    string      `json:"status"`
	MatchList []struct {
		Currency     string      `json:"currency"`
		DailyIntRate json.Number `json:"dailyIntRate"`
		Size         json.Number `json:"size"`
		Term         json.Number `json:"term"`
		Timestamp    json.Number `json:"timestamp"`
		TradeId      string      `json:"tradeId"`
	} `json:"matchList"`
}

// BorrowOrder returns a specific borrow order
func (as *ApiService) BorrowOrder(orderId string) (*ApiResponse, error) {
	params := map[string]string{}
	if orderId != "" {
		params["orderId"] = orderId
	}
	req := NewRequest(http.MethodGet, "/api/v1/margin/borrow", params)
	return as.Call(req)
}

// BorrowOutstandingRecordModel represents borrow outstanding record
type BorrowOutstandingRecordModel struct {
	Currency        string      `json:"currency"`
	TradeId         string      `json:"tradeId"`
	Liability       json.Number `json:"liability"`
	Principal       json.Number `json:"principal"`
	AccruedInterest json.Number `json:"accruedInterest"`
	CreatedAt       json.Number `json:"createdAt"`
	MaturityTime    json.Number `json:"maturityTime"`
	Term            json.Number `json:"term"`
	RepaidSize      json.Number `json:"repaidSize"`
	DailyIntRate    json.Number `json:"dailyIntRate"`
}

// BorrowOutstandingRecordsModel is a list of *BorrowOutstandingRecordModel
type BorrowOutstandingRecordsModel []*BorrowOutstandingRecordModel

// BorrowOutstandingRecords returns borrow outstanding records
func (as *ApiService) BorrowOutstandingRecords(currency string, pagination *PaginationParam) (*ApiResponse, error) {
	params := map[string]string{}
	if currency != "" {
		params["currency"] = currency
	}
	pagination.ReadParam(params)
	req := NewRequest(http.MethodGet, "/api/v1/margin/borrow/outstanding", params)
	return as.Call(req)
}

// BorrowRepaidRecordModel represents a repaid borrow record
type BorrowRepaidRecordModel struct {
	Currency     string      `json:"currency"`
	DailyIntRate json.Number `json:"dailyIntRate"`
	Interest     json.Number `json:"interest"`
	Principal    json.Number `json:"principal"`
	RepaidSize   json.Number `json:"repaidSize"`
	RepayTime    json.Number `json:"repayTime"`
	Term         json.Number `json:"term"`
	TradeId      string      `json:"tradeId"`
}

// BorrowRepaidRecordsModel is a list of *BorrowRepaidRecordModel
type BorrowRepaidRecordsModel []*BorrowRepaidRecordModel

// BorrowRepaidRecords returns repaid borrow records
func (as *ApiService) BorrowRepaidRecords(currency string, pagination *PaginationParam) (*ApiResponse, error) {
	params := map[string]string{}
	if currency != "" {
		params["currency"] = currency
	}
	pagination.ReadParam(params)
	req := NewRequest(http.MethodGet, "/api/v1/margin/borrow/repaid", params)
	return as.Call(req)
}

// RepayAll repay borrow orders of one currency
func (as *ApiService) RepayAll(params map[string]string) (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v1/margin/repay/all", params)
	return as.Call(req)
}

// RepaySingle repay a single borrow order
func (as *ApiService) RepaySingle(params map[string]string) (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v1/margin/repay/single", params)
	return as.Call(req)
}

// CreateLendOrderResultModel the result of create a lend order
type CreateLendOrderResultModel struct {
	OrderId string `json:"orderId"`
}

// CreateLendOrder returns the result of create a lend order
func (as *ApiService) CreateLendOrder(params map[string]string) (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v1/margin/lend", params)
	return as.Call(req)
}

// CancelLendOrder cancel a lend order
func (as *ApiService) CancelLendOrder(orderId string) (*ApiResponse, error) {
	req := NewRequest(http.MethodDelete, fmt.Sprintf("/api/v1/margin/lend/%s", orderId), nil)
	return as.Call(req)
}

// ToggleAutoLend set auto lend rules
func (as *ApiService) ToggleAutoLend(params map[string]string) (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v1/margin/toggle-auto-lend", params)
	return as.Call(req)
}

// LendOrderBaseModel represents Base model of lend order
type LendOrderBaseModel struct {
	OrderId      string      `json:"orderId"`
	Currency     string      `json:"currency"`
	Size         json.Number `json:"size"`
	FilledSize   json.Number `json:"filledSize"`
	DailyIntRate json.Number `json:"dailyIntRate"`
	Term         json.Number `json:"term"`
	CreatedAt    json.Number `json:"createdAt"`
}

// LendActiveOrderModel represents a active lend order
type LendActiveOrderModel struct {
	LendOrderBaseModel
}

// LendActiveOrdersModel is a list of *LendActiveOrderModel
type LendActiveOrdersModel []*LendActiveOrderModel

// LendActiveOrders returns the active lend orders
func (as *ApiService) LendActiveOrders(currency string, pagination *PaginationParam) (*ApiResponse, error) {
	params := map[string]string{}
	if currency != "" {
		params["currency"] = currency
	}

	pagination.ReadParam(params)
	req := NewRequest(http.MethodGet, "/api/v1/margin/lend/active", params)

	return as.Call(req)
}

// LendDoneOrderModel represents a history lend order
type LendDoneOrderModel struct {
	LendOrderBaseModel
	Status string `json:"status"`
}

// LendDoneOrdersModel is a list of *LendDoneOrderModel
type LendDoneOrdersModel []*LendDoneOrderModel

// LendDoneOrders returns the history lend orders
func (as *ApiService) LendDoneOrders(currency string, pagination *PaginationParam) (*ApiResponse, error) {
	params := map[string]string{}
	if currency != "" {
		params["currency"] = currency
	}

	pagination.ReadParam(params)
	req := NewRequest(http.MethodGet, "/api/v1/margin/lend/done", params)

	return as.Call(req)
}

// LendTradeUnsettledRecordModel represents a unsettled lend record
type LendTradeUnsettledRecordModel struct {
	TradeId         string      `json:"tradeId"`
	Currency        string      `json:"currency"`
	Size            json.Number `json:"size"`
	AccruedInterest json.Number `json:"accruedInterest"`
	Repaid          json.Number `json:"repaid"`
	DailyIntRate    json.Number `json:"dailyIntRate"`
	Term            json.Number `json:"term"`
	MaturityTime    json.Number `json:"maturityTime"`
}

// LendTradeUnsettledRecordsModel is a list of *LendTradeUnsettledRecordModel
type LendTradeUnsettledRecordsModel []*LendTradeUnsettledRecordModel

// LendTradeUnsettledRecords returns unsettled lend records
func (as *ApiService) LendTradeUnsettledRecords(currency string, pagination *PaginationParam) (*ApiResponse, error) {
	params := map[string]string{}
	if currency != "" {
		params["currency"] = currency
	}

	pagination.ReadParam(params)
	req := NewRequest(http.MethodGet, "/api/v1/margin/lend/trade/unsettled", params)
	return as.Call(req)
}

// LendTradeSettledRecordModel represents a settled lend record
type LendTradeSettledRecordModel struct {
	TradeId      string      `json:"tradeId"`
	Currency     string      `json:"currency"`
	Size         json.Number `json:"size"`
	Interest     json.Number `json:"interest"`
	Repaid       json.Number `json:"repaid"`
	DailyIntRate json.Number `json:"dailyIntRate"`
	Term         json.Number `json:"term"`
	SettledAt    json.Number `json:"settledAt"`
	Note         string      `json:"note"`
}

// LendTradeSettledRecordsModel is a list of *LendTradeSettledRecordModel
type LendTradeSettledRecordsModel []*LendTradeSettledRecordModel

// LendTradeSettledRecords returns settled lend records
func (as *ApiService) LendTradeSettledRecords(currency string, pagination *PaginationParam) (*ApiResponse, error) {
	params := map[string]string{}
	if currency != "" {
		params["currency"] = currency
	}

	pagination.ReadParam(params)
	req := NewRequest(http.MethodGet, "/api/v1/margin/lend/trade/settled", params)

	return as.Call(req)
}

// LendAssetModel represents account lend asset
type LendAssetModel struct {
	Currency        string      `json:"currency"`
	Outstanding     json.Number `json:"outstanding"`
	FilledSize      json.Number `json:"filledSize"`
	AccruedInterest json.Number `json:"accruedInterest"`
	RealizedProfit  json.Number `json:"realizedProfit"`
	IsAutoLend      bool        `json:"isAutoLend"`
}

// LendAssetsModel is a list of *LendAssetModel
type LendAssetsModel []*LendAssetModel

// LendAssets returns account lend assets
func (as *ApiService) LendAssets(currency string) (*ApiResponse, error) {
	params := map[string]string{}
	if currency != "" {
		params["currency"] = currency
	}

	req := NewRequest(http.MethodGet, "/api/v1/margin/lend/assets", params)
	return as.Call(req)
}

// MarginMarketModel represents lending market data
type MarginMarketModel struct {
	DailyIntRate json.Number `json:"dailyIntRate"`
	Term         json.Number `json:"term"`
	Size         json.Number `json:"size"`
}

// MarginMarketsModel is a list of *MarginMarketModel
type MarginMarketsModel []*MarginMarketModel

// MarginMarkets returns lending market data
func (as *ApiService) MarginMarkets(params map[string]string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/margin/market", params)
	return as.Call(req)
}

// MarginTradeModel represents lending market trade data
type MarginTradeModel struct {
	TradeId      string      `json:"tradeId"`
	Currency     string      `json:"currency"`
	Size         json.Number `json:"size"`
	DailyIntRate json.Number `json:"dailyIntRate"`
	Term         json.Number `json:"term"`
	Timestamp    json.Number `json:"timestamp"`
}

// MarginTradesModel is a list of *MarginTradeModel
type MarginTradesModel []*MarginTradeModel

// MarginTradeLast returns latest lending market trade datas
func (as *ApiService) MarginTradeLast(currency string) (*ApiResponse, error) {
	params := map[string]string{}
	if currency != "" {
		params["currency"] = currency
	}

	req := NewRequest(http.MethodGet, "/api/v1/margin/trade/last", params)
	return as.Call(req)
}

// MarginRiskLimitItemModel is item of *MarginRiskLimitModel
type MarginRiskLimitItemModel struct {
	Currency        string      `json:"currency"`
	BorrowMaxAmount string      `json:"borrowMaxAmount"`
	BuyMaxAmount    string      `json:"buyMaxAmount"`
	Precision       json.Number `json:"precision"`
}

// MarginRiskLimitModel is a list of *MarginRiskLimitModel
type MarginRiskLimitModel []*MarginRiskLimitItemModel

func (as *ApiService) MarginRiskLimit(marginModel string) (*ApiResponse, error) {
	params := map[string]string{}
	if marginModel != "" {
		params["marginModel"] = marginModel
	}

	req := NewRequest(http.MethodGet, "/api/v1/risk/limit/strategy", params)
	return as.Call(req)
}
