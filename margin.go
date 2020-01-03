package kucoin

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// A MarkPriceModel represents mark price of a symbol
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

// A MarginConfigModel represents a margin configuration
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

// A MarginAccountModel represents a margin account information
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

//发布借入委托
type CreateBorrowOrderResultModel struct {
	OrderId  string `json:"orderId"`
	Currency string `json:"currency"`
}

func (as *ApiService) CreateBorrowOrder(params map[string]string) (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v1/margin/borrow", params)
	return as.Call(req)
}

//查询借入委托
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

func (as *ApiService) BorrowOrder(orderId string) (*ApiResponse, error) {
	params := map[string]string{}
	if orderId != "" {
		params["orderId"] = orderId
	}
	req := NewRequest(http.MethodGet, "/api/v1/margin/borrow", params)
	return as.Call(req)
}

//查询待还款记录
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

type BorrowOutstandingRecordsModel []*BorrowOutstandingRecordModel

func (as *ApiService) BorrowOutstandingRecords(currency string, pagination *PaginationParam) (*ApiResponse, error) {
	params := map[string]string{}
	if currency != "" {
		params["currency"] = currency
	}
	pagination.ReadParam(params)
	req := NewRequest(http.MethodGet, "/api/v1/margin/borrow/outstanding", params)
	return as.Call(req)
}

//查询已还款记录
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
type BorrowRepaidRecordsModel []*BorrowRepaidRecordModel

func (as *ApiService) BorrowRepaidRecords(currency string, pagination *PaginationParam) (*ApiResponse, error) {
	params := map[string]string{}
	if currency != "" {
		params["currency"] = currency
	}
	pagination.ReadParam(params)
	req := NewRequest(http.MethodGet, "/api/v1/margin/borrow/repaid", params)
	return as.Call(req)
}

//一键还款
func (as *ApiService) RepayAll(params map[string]string) (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v1/margin/repay/all", params)
	return as.Call(req)
}

//一键还款单笔
func (as *ApiService) RepaySingle(params map[string]string) (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v1/margin/repay/single", params)
	return as.Call(req)
}

//发布借出委托
type CreateLendOrderResultModel struct {
	OrderId string `json:"orderId"`
}

func (as *ApiService) CreateLendOrder(params map[string]string) (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v1/margin/lend", params)
	return as.Call(req)
}

//撤销借出委托
func (as *ApiService) CancelLendOrder(orderId string) (*ApiResponse, error) {
	req := NewRequest(http.MethodDelete, fmt.Sprintf("/api/v1/margin/lend/%s", orderId), nil)
	return as.Call(req)
}

//设置自动借出
func (as *ApiService) ToggleAutoLend(params map[string]string) (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v1/margin/toggle-auto-lend", params)
	return as.Call(req)
}

//查询活跃借出委托
type LendOrderBaseModel struct {
	OrderId      string      `json:"orderId"`
	Currency     string      `json:"currency"`
	Size         json.Number `json:"size"`
	FilledSize   json.Number `json:"filledSize"`
	DailyIntRate json.Number `json:"dailyIntRate"`
	Term         json.Number `json:"term"`
	CreatedAt    json.Number `json:"createdAt"`
}

type LendActiveOrderModel struct {
	LendOrderBaseModel
}

type LendActiveOrdersModel []*LendActiveOrderModel

func (as *ApiService) LendActiveOrders(currency string, pagination *PaginationParam) (*ApiResponse, error) {
	params := map[string]string{}
	if currency != "" {
		params["currency"] = currency
	}

	pagination.ReadParam(params)
	req := NewRequest(http.MethodGet, "/api/v1/margin/lend/active", params)

	return as.Call(req)
}

//查询历史借出委托
type LendDoneOrderModel struct {
	LendOrderBaseModel
	Status string `json:"status"`
}

type LendDoneOrdersModel []*LendDoneOrderModel

func (as *ApiService) LendDoneOrders(currency string, pagination *PaginationParam) (*ApiResponse, error) {
	params := map[string]string{}
	if currency != "" {
		params["currency"] = currency
	}

	pagination.ReadParam(params)
	req := NewRequest(http.MethodGet, "/api/v1/margin/lend/done", params)

	return as.Call(req)
}

//查询未结算出借记录
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

type LendTradeUnsettledRecordsModel []*LendTradeUnsettledRecordModel

func (as *ApiService) LendTradeUnsettledRecords(currency string, pagination *PaginationParam) (*ApiResponse, error) {
	params := map[string]string{}
	if currency != "" {
		params["currency"] = currency
	}

	pagination.ReadParam(params)
	req := NewRequest(http.MethodGet, "/api/v1/margin/lend/trade/unsettled", params)
	return as.Call(req)
}

//查询已结算出借记录
type LendTradeSettledRecordModel struct {
	TradeId      string      `json:"tradeId"`
	Currency     string      `json:"currency"`
	Size         json.Number `json:"size"`
	Interest     json.Number `json:"interest"`
	Repaid       json.Number `json:"repaid"`
	DailyIntRate json.Number `json:"dailyIntRate"`
	Term         json.Number `json:"term"`
	SettledAt    json.Number `json:"settledAt"`
	Note         json.Number `json:"note"`
}

type LendTradeSettledRecordsModel []*LendTradeSettledRecordModel

func (as *ApiService) LendTradeSettledRecords(currency string, pagination *PaginationParam) (*ApiResponse, error) {
	params := map[string]string{}
	if currency != "" {
		params["currency"] = currency
	}

	pagination.ReadParam(params)
	req := NewRequest(http.MethodGet, "/api/v1/margin/lend/trade/settled", params)

	return as.Call(req)
}

//资产借出记录
type LendAssetModel struct {
	Currency        string      `json:"currency"`
	Outstanding     json.Number `json:"outstanding"`
	FilledSize      json.Number `json:"filledSize"`
	AccruedInterest json.Number `json:"accruedInterest"`
	RealizedProfit  json.Number `json:"realizedProfit"`
	IsAutoLend      bool        `json:"isAutoLend"`
}

type LendAssetsModel []*LendAssetModel

func (as *ApiService) LendAssets(currency string) (*ApiResponse, error) {
	params := map[string]string{}
	if currency != "" {
		params["currency"] = currency
	}

	req := NewRequest(http.MethodGet, "/api/v1/margin/lend/assets", params)
	return as.Call(req)
}

//借出市场信息
type MarginMarketModel struct {
	DailyIntRate json.Number `json:"dailyIntRate"`
	Term         json.Number `json:"term"`
	Size         json.Number `json:"size"`
}

type MarginMarketsModel []*MarginMarketModel

func (as *ApiService) MarginMarkets(params map[string]string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/margin/market", params)
	return as.Call(req)
}

//借贷市场成交信息
type MarginTradeModel struct {
	TradeId      string      `json:"tradeId"`
	Currency     string      `json:"currency"`
	Size         json.Number `json:"size"`
	DailyIntRate json.Number `json:"dailyIntRate"`
	Term         json.Number `json:"term"`
	Timestamp    json.Number `json:"timestamp"`
}

type MarginTradesModel []*MarginTradeModel

func (as *ApiService) MarginTradeLast(currency string) (*ApiResponse, error) {
	params := map[string]string{}
	if currency != "" {
		params["currency"] = currency
	}

	req := NewRequest(http.MethodGet, "/api/v1/margin/trade/last", params)
	return as.Call(req)
}
