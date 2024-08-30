package kucoin

import (
	"context"
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
func (as *ApiService) CurrentMarkPrice(ctx context.Context, symbol string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/mark-price/%s/current", symbol), nil)
	return as.Call(ctx, req)
}

// MarginConfigModel represents a margin configuration
type MarginConfigModel struct {
	CurrencyList     []string    `json:"currencyList"`
	WarningDebtRatio json.Number `json:"warningDebtRatio"`
	LiqDebtRatio     json.Number `json:"liqDebtRatio"`
	MaxLeverage      json.Number `json:"maxLeverage"`
}

// MarginConfig returns a margin configuration
func (as *ApiService) MarginConfig(ctx context.Context) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/margin/config", nil)
	return as.Call(ctx, req)
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
func (as *ApiService) MarginAccount(ctx context.Context) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/margin/account", nil)
	return as.Call(ctx, req)
}

// CreateBorrowOrderResultModel represents the result of create a borrow order
type CreateBorrowOrderResultModel struct {
	OrderId  string `json:"orderId"`
	Currency string `json:"currency"`
}

// CreateBorrowOrder returns the result of create a borrow order
// Deprecated please use MarginBorrowV3
func (as *ApiService) CreateBorrowOrder(ctx context.Context, params map[string]string) (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v1/margin/borrow", params)
	return as.Call(ctx, req)
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
// Deprecated please use QueryMarginBorrowV3
func (as *ApiService) BorrowOrder(ctx context.Context, orderId string) (*ApiResponse, error) {
	params := map[string]string{}
	if orderId != "" {
		params["orderId"] = orderId
	}
	req := NewRequest(http.MethodGet, "/api/v1/margin/borrow", params)
	return as.Call(ctx, req)
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
func (as *ApiService) BorrowOutstandingRecords(ctx context.Context, currency string, pagination *PaginationParam) (*ApiResponse, error) {
	params := map[string]string{}
	if currency != "" {
		params["currency"] = currency
	}
	pagination.ReadParam(params)
	req := NewRequest(http.MethodGet, "/api/v1/margin/borrow/outstanding", params)
	return as.Call(ctx, req)
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
// Deprecated please use QueryMarginRepayV3
func (as *ApiService) BorrowRepaidRecords(ctx context.Context, currency string, pagination *PaginationParam) (*ApiResponse, error) {
	params := map[string]string{}
	if currency != "" {
		params["currency"] = currency
	}
	pagination.ReadParam(params)
	req := NewRequest(http.MethodGet, "/api/v1/margin/borrow/repaid", params)
	return as.Call(ctx, req)
}

// RepayAll repay borrow orders of one currency
// Deprecated please use MarginRepayV3
func (as *ApiService) RepayAll(ctx context.Context, params map[string]string) (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v1/margin/repay/all", params)
	return as.Call(ctx, req)
}

// RepaySingle repay a single borrow order
// Deprecated please use MarginRepayV3
func (as *ApiService) RepaySingle(ctx context.Context, params map[string]string) (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v1/margin/repay/single", params)
	return as.Call(ctx, req)
}

// CreateLendOrderResultModel the result of create a lend order
type CreateLendOrderResultModel struct {
	OrderId string `json:"orderId"`
}

// CreateLendOrder returns the result of create a lend order
// Deprecated please use LendingPurchaseV3
func (as *ApiService) CreateLendOrder(ctx context.Context, params map[string]string) (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v1/margin/lend", params)
	return as.Call(ctx, req)
}

// CancelLendOrder cancel a lend order
// Deprecated please use LendingRedeemV3
func (as *ApiService) CancelLendOrder(ctx context.Context, orderId string) (*ApiResponse, error) {
	req := NewRequest(http.MethodDelete, fmt.Sprintf("/api/v1/margin/lend/%s", orderId), nil)
	return as.Call(ctx, req)
}

// ToggleAutoLend set auto lend rules
// Deprecated
func (as *ApiService) ToggleAutoLend(ctx context.Context, params map[string]string) (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v1/margin/toggle-auto-lend", params)
	return as.Call(ctx, req)
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
// Deprecated
func (as *ApiService) LendActiveOrders(ctx context.Context, currency string, pagination *PaginationParam) (*ApiResponse, error) {
	params := map[string]string{}
	if currency != "" {
		params["currency"] = currency
	}

	pagination.ReadParam(params)
	req := NewRequest(http.MethodGet, "/api/v1/margin/lend/active", params)

	return as.Call(ctx, req)
}

// LendDoneOrderModel represents a history lend order
type LendDoneOrderModel struct {
	LendOrderBaseModel
	Status string `json:"status"`
}

// LendDoneOrdersModel is a list of *LendDoneOrderModel
type LendDoneOrdersModel []*LendDoneOrderModel

// LendDoneOrders returns the history lend orders
// Deprecated
func (as *ApiService) LendDoneOrders(ctx context.Context, currency string, pagination *PaginationParam) (*ApiResponse, error) {
	params := map[string]string{}
	if currency != "" {
		params["currency"] = currency
	}

	pagination.ReadParam(params)
	req := NewRequest(http.MethodGet, "/api/v1/margin/lend/done", params)

	return as.Call(ctx, req)
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
// Deprecated
func (as *ApiService) LendTradeUnsettledRecords(ctx context.Context, currency string, pagination *PaginationParam) (*ApiResponse, error) {
	params := map[string]string{}
	if currency != "" {
		params["currency"] = currency
	}

	pagination.ReadParam(params)
	req := NewRequest(http.MethodGet, "/api/v1/margin/lend/trade/unsettled", params)
	return as.Call(ctx, req)
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
// Deprecated
func (as *ApiService) LendTradeSettledRecords(ctx context.Context, currency string, pagination *PaginationParam) (*ApiResponse, error) {
	params := map[string]string{}
	if currency != "" {
		params["currency"] = currency
	}

	pagination.ReadParam(params)
	req := NewRequest(http.MethodGet, "/api/v1/margin/lend/trade/settled", params)

	return as.Call(ctx, req)
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
// Deprecated
func (as *ApiService) LendAssets(ctx context.Context, currency string) (*ApiResponse, error) {
	params := map[string]string{}
	if currency != "" {
		params["currency"] = currency
	}

	req := NewRequest(http.MethodGet, "/api/v1/margin/lend/assets", params)
	return as.Call(ctx, req)
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
// Deprecated
func (as *ApiService) MarginMarkets(ctx context.Context, params map[string]string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/margin/market", params)
	return as.Call(ctx, req)
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
// Deprecated
func (as *ApiService) MarginTradeLast(ctx context.Context, currency string) (*ApiResponse, error) {
	params := map[string]string{}
	if currency != "" {
		params["currency"] = currency
	}

	req := NewRequest(http.MethodGet, "/api/v1/margin/trade/last", params)
	return as.Call(ctx, req)
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

func (as *ApiService) MarginRiskLimit(ctx context.Context, marginModel string) (*ApiResponse, error) {
	params := map[string]string{}
	if marginModel != "" {
		params["marginModel"] = marginModel
	}

	req := NewRequest(http.MethodGet, "/api/v1/risk/limit/strategy", params)
	return as.Call(ctx, req)
}

// MarginIsolatedSymbols  query margin isolated symbols
func (as *ApiService) MarginIsolatedSymbols(ctx context.Context) (*ApiResponse, error) {
	p := map[string]string{}
	req := NewRequest(http.MethodGet, "/api/v1/isolated/symbols", p)
	return as.Call(ctx, req)
}

type MarginIsolatedSymbolsModel []*MarginIsolatedSymbolModel

type MarginIsolatedSymbolModel struct {
	Symbol                string `json:"symbol"`
	SymbolName            string `json:"symbolName"`
	BaseCurrency          string `json:"baseCurrency"`
	QuoteCurrency         string `json:"quoteCurrency"`
	MaxLeverage           int64  `json:"maxLeverage"`
	FlDebtRatio           string `json:"flDebtRatio"`
	TradeEnable           bool   `json:"tradeEnable"`
	AutoRenewMaxDebtRatio string `json:"autoRenewMaxDebtRatio"`
	BaseBorrowEnable      bool   `json:"baseBorrowEnable"`
	QuoteBorrowEnable     bool   `json:"quoteBorrowEnable"`
	BaseTransferInEnable  bool   `json:"baseTransferInEnable"`
	QuoteTransferInEnable bool   `json:"quoteTransferInEnable"`
}

// MarginIsolatedAccounts query margin isolated account
func (as *ApiService) MarginIsolatedAccounts(ctx context.Context, balanceCurrency string) (*ApiResponse, error) {
	p := map[string]string{
		"balanceCurrency": balanceCurrency,
	}
	req := NewRequest(http.MethodGet, "/api/v1/isolated/accounts", p)
	return as.Call(ctx, req)
}

type MarginIsolatedAccountsModel struct {
	TotalConversionBalance     string                              `json:"totalConversionBalance"`
	LiabilityConversionBalance string                              `json:"liabilityConversionBalance"`
	Assets                     []*MarginIsolatedAccountAssetsModel `json:"assets"`
}

type MarginIsolatedAccountAssetsModel struct {
	Symbol    string `json:"symbol"`
	Status    string `json:"status"`
	DebtRatio string `json:"debtRatio"`
	BaseAsset struct {
		Currency         string `json:"currency"`
		TotalBalance     string `json:"totalBalance"`
		HoldBalance      string `json:"holdBalance"`
		AvailableBalance string `json:"availableBalance"`
		Liability        string `json:"liability"`
		Interest         string `json:"interest"`
		BorrowableAmount string `json:"borrowableAmount"`
	} `json:"baseAsset"`
	QuoteAsset struct {
		Currency         string `json:"currency"`
		TotalBalance     string `json:"totalBalance"`
		HoldBalance      string `json:"holdBalance"`
		AvailableBalance string `json:"availableBalance"`
		Liability        string `json:"liability"`
		Interest         string `json:"interest"`
		BorrowableAmount string `json:"borrowableAmount"`
	} `json:"quoteAsset"`
}

// IsolatedAccount query margin isolated account by symbol
func (as *ApiService) IsolatedAccount(ctx context.Context, symbol string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/isolated/account/"+symbol, nil)
	return as.Call(ctx, req)
}

// IsolatedBorrow  margin isolated borrow
// Deprecated
func (as *ApiService) IsolatedBorrow(ctx context.Context, params map[string]string) (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v1/isolated/borrow", params)
	return as.Call(ctx, req)
}

type MarginIsolatedBorrowRes struct {
	OrderId    string `json:"orderId"`
	Currency   string `json:"currency"`
	ActualSize string `json:"actualSize"`
}

// IsolatedBorrowOutstandingRecord query margin isolated borrow outstanding records
// Deprecated
func (as *ApiService) IsolatedBorrowOutstandingRecord(ctx context.Context, params map[string]string, pagination *PaginationParam) (*ApiResponse, error) {
	pagination.ReadParam(params)
	req := NewRequest(http.MethodGet, "/api/v1/isolated/borrow/outstanding", params)
	return as.Call(ctx, req)
}

type IsolatedBorrowOutstandingRecordsModel []*IsolatedBorrowOutstandingRecordModel

type IsolatedBorrowOutstandingRecordModel struct {
	LoanId            string      `json:"loanId"`
	Symbol            string      `json:"symbol"`
	Currency          string      `json:"currency"`
	LiabilityBalance  string      `json:"liabilityBalance"`
	PrincipalTotal    string      `json:"principalTotal"`
	InterestBalance   string      `json:"interestBalance"`
	CreatedAt         json.Number `json:"createdAt"`
	MaturityTime      json.Number `json:"maturityTime"`
	Period            int64       `json:"period"`
	RepaidSize        string      `json:"repaidSize"`
	DailyInterestRate string      `json:"dailyInterestRate"`
}

// IsolatedBorrowRepaidRecord query margin isolated borrow repaid records
// Deprecated
func (as *ApiService) IsolatedBorrowRepaidRecord(ctx context.Context, params map[string]string, pagination *PaginationParam) (*ApiResponse, error) {
	pagination.ReadParam(params)
	req := NewRequest(http.MethodGet, "/api/v1/isolated/borrow/repaid", params)
	return as.Call(ctx, req)
}

type IsolatedBorrowRepaidRecordRecordsModel []*IsolatedBorrowRepaidRecordRecordModel

type IsolatedBorrowRepaidRecordRecordModel struct {
	LoanId            string      `json:"loanId"`
	Symbol            string      `json:"symbol"`
	Currency          string      `json:"currency"`
	PrincipalTotal    string      `json:"principalTotal"`
	InterestBalance   string      `json:"interestBalance"`
	RepaidSize        string      `json:"repaidSize"`
	CreatedAt         json.Number `json:"createdAt"`
	Period            int64       `json:"period"`
	DailyInterestRate string      `json:"dailyInterestRate"`
	RepayFinishAt     json.Number `json:"repayFinishAt"`
}

// IsolatedRepayAll repay all  isolated
// Deprecated
func (as *ApiService) IsolatedRepayAll(ctx context.Context, params map[string]string) (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v1/isolated/repay/all", params)
	return as.Call(ctx, req)
}

// IsolatedRepaySingle repay single  isolated
// Deprecated
func (as *ApiService) IsolatedRepaySingle(ctx context.Context, params map[string]string) (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v1/isolated/repay/single", params)
	return as.Call(ctx, req)
}

type MarginCurrenciesModel []*MarginCurrencyModel

type MarginCurrencyModel struct {
	Currency       string      `json:"currency"`
	NetAsset       json.Number `json:"netAsset"`
	TargetLeverage string      `json:"targetLeverage"`
	ActualLeverage string      `json:"actualLeverage"`
	IssuedSize     string      `json:"issuedSize"`
	Basket         string      `json:"basket"`
}

// MarginCurrencyInfo margin currency info
func (as *ApiService) MarginCurrencyInfo(ctx context.Context, currency string) (*ApiResponse, error) {
	p := map[string]string{
		"currency": currency,
	}
	req := NewRequest(http.MethodGet, "/api/v3/etf/info", p)
	return as.Call(ctx, req)
}

type MarginCurrenciesRiskLimitModel []*MarginCurrencyRiskLimitModel

type MarginCurrencyRiskLimitModel struct {
	Timestamp         json.Number `json:"timestamp"`
	Currency          string      `json:"currency"`
	BorrowMaxAmount   string      `json:"borrowMaxAmount"`
	BuyMaxAmount      string      `json:"buyMaxAmount"`
	HoldMaxAmount     string      `json:"holdMaxAmount"`
	BorrowCoefficient string      `json:"borrowCoefficient"`
	MarginCoefficient string      `json:"marginCoefficient"`
	Precision         int64       `json:"precision"`
	BorrowMinAmount   string      `json:"borrowMinAmount"`
	BorrowMinUnit     string      `json:"borrowMinUnit"`
	BorrowEnabled     bool        `json:"borrowEnabled"`
}

type IsolatedCurrenciesRiskLimitModel []*IsolatedCurrencyRiskLimitModel

type IsolatedCurrencyRiskLimitModel struct {
	Timestamp              json.Number `json:"timestamp"`
	Symbol                 string      `json:"symbol"`
	BaseMaxBorrowAmount    string      `json:"baseMaxBorrowAmount"`
	QuoteMaxBorrowAmount   string      `json:"quoteMaxBorrowAmount"`
	BaseMaxBuyAmount       string      `json:"baseMaxBuyAmount"`
	QuoteMaxBuyAmount      string      `json:"quoteMaxBuyAmount"`
	BaseMaxHoldAmount      string      `json:"baseMaxHoldAmount"`
	QuoteMaxHoldAmount     string      `json:"quoteMaxHoldAmount"`
	BasePrecision          int64       `json:"basePrecision"`
	QuotePrecision         int64       `json:"quotePrecision"`
	BaseBorrowCoefficient  string      `json:"baseBorrowCoefficient"`
	QuoteBorrowCoefficient string      `json:"quoteBorrowCoefficient"`
	BaseMarginCoefficient  string      `json:"baseMarginCoefficient"`
	QuoteMarginCoefficient string      `json:"quoteMarginCoefficient"`
	BaseBorrowMinAmount    string      `json:"baseBorrowMinAmount"`
	BaseBorrowMinUnit      string      `json:"baseBorrowMinUnit"`
	QuoteBorrowMinAmount   string      `json:"quoteBorrowMinAmount"`
	QuoteBorrowMinUnit     string      `json:"quoteBorrowMinUnit"`
	BaseBorrowEnabled      bool        `json:"baseBorrowEnabled"`
	QuoteBorrowEnabled     bool        `json:"quoteBorrowEnabled"`
}

// MarginCurrencies This interface can obtain the risk limit and currency configuration of cross margin/isolated margin.
func (as *ApiService) MarginCurrencies(ctx context.Context, currency, symbol, isIsolated string) (*ApiResponse, error) {
	p := map[string]string{
		"currency":   currency,
		"symbol":     symbol,
		"isIsolated": isIsolated,
	}
	req := NewRequest(http.MethodGet, "/api/v3/margin/currencies", p)
	return as.Call(ctx, req)
}

type MarginBorrowV3Req struct {
	IsIsolated  bool   `json:"isIsolated"`
	Symbol      string `json:"symbol"`
	Currency    string `json:"currency"`
	Size        string `json:"size"`
	TimeInForce string `json:"timeInForce"`
	IsHf        bool   `json:"isHf"`
}

type MarginBorrowV3Res struct {
	OrderNo    string      `json:"orderNo"`
	ActualSize json.Number `json:"actualSize"`
}

// MarginBorrowV3 initiate an application for cross or isolated margin borrowing
func (as *ApiService) MarginBorrowV3(ctx context.Context, p *MarginBorrowV3Req) (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v3/margin/borrow", p)
	return as.Call(ctx, req)
}

type MarginBorrowsV3Model []*MarginBorrowV3Model

type MarginBorrowV3Model struct {
	OrderNo     string      `json:"orderNo"`
	Symbol      string      `json:"symbol"`
	Currency    string      `json:"currency"`
	Size        json.Number `json:"size"`
	ActualSize  json.Number `json:"actualSize"`
	Status      string      `json:"status"`
	CreatedTime int64       `json:"createdTime"`
}

// QueryMarginBorrowV3 get the borrowing orders for cross and isolated margin accounts
func (as *ApiService) QueryMarginBorrowV3(ctx context.Context, p map[string]string, pagination *PaginationParam) (*ApiResponse, error) {
	pagination.ReadParam(p)
	req := NewRequest(http.MethodGet, "/api/v3/margin/borrow", p)
	return as.Call(ctx, req)
}

type MarginRepay3VReq struct {
	IsIsolated bool   `json:"isIsolated"`
	Symbol     string `json:"symbol"`
	Currency   string `json:"currency"`
	Size       string `json:"size"`
	IsHf       bool   `json:"isHf"`
}

type MarginRepayV3Res struct {
	OrderNo    string      `json:"orderNo"`
	ActualSize json.Number `json:"actualSize"`
}

// MarginRepayV3 initiate an application for the repayment of cross or isolated margin borrowing
func (as *ApiService) MarginRepayV3(ctx context.Context, p *MarginRepay3VReq) (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v3/margin/repay", p)
	return as.Call(ctx, req)
}

type MarginRepayV3Model struct {
	OrderNo     string      `json:"orderNo"`
	Symbol      string      `json:"symbol"`
	Currency    string      `json:"currency"`
	Size        json.Number `json:"size"`
	ActualSize  json.Number `json:"actualSize"`
	Status      string      `json:"status"`
	CreatedTime int64       `json:"createdTime"`
}

type MarginRepaysV3Model []*MarginRepayV3Model

// QueryMarginRepayV3 get  repay orders for cross and isolated margin accounts
func (as *ApiService) QueryMarginRepayV3(ctx context.Context, p map[string]string, pagination *PaginationParam) (*ApiResponse, error) {
	pagination.ReadParam(p)
	req := NewRequest(http.MethodGet, "/api/v3/margin/repay", p)
	return as.Call(ctx, req)
}

type MarginInterestV3Model struct {
	CreatedAt      int64       `json:"createdAt"`
	Currency       string      `json:"currency"`
	InterestAmount json.Number `json:"interestAmount"`
	DayRatio       json.Number `json:"dayRatio"`
}

type MarginInterestsV3Model []*MarginInterestV3Model

// QueryInterestV3 get the interest records of the cross/isolated margin lending
func (as *ApiService) QueryInterestV3(ctx context.Context, p map[string]string, pagination *PaginationParam) (*ApiResponse, error) {
	pagination.ReadParam(p)
	req := NewRequest(http.MethodGet, "/api/v3/margin/interest", p)
	return as.Call(ctx, req)
}

type MarginCurrencyV3Model struct {
	Currency           string `json:"currency"`
	PurchaseEnable     bool   `json:"purchaseEnable"`
	RedeemEnable       bool   `json:"redeemEnable"`
	Increment          string `json:"increment"`
	MinPurchaseSize    string `json:"minPurchaseSize"`
	MinInterestRate    string `json:"minInterestRate"`
	MaxInterestRate    string `json:"maxInterestRate"`
	InterestIncrement  string `json:"interestIncrement"`
	MaxPurchaseSize    string `json:"maxPurchaseSize"`
	MarketInterestRate string `json:"marketInterestRate"`
	AutoPurchaseEnable bool   `json:"autoPurchaseEnable"`
}

type MarginCurrenciesV3Model []*MarginCurrencyV3Model

// QueryMarginCurrenciesV3 get the interest records of the cross/isolated margin lending
func (as *ApiService) QueryMarginCurrenciesV3(ctx context.Context, currency string) (*ApiResponse, error) {
	p := map[string]string{
		"currency": currency,
	}
	req := NewRequest(http.MethodGet, "/api/v3/project/list", p)
	return as.Call(ctx, req)
}

type MarginInterestRateV3Model struct {
	Time               string `json:"time"`
	MarketInterestRate string `json:"marketInterestRate"`
}

type MarginInterestRatesV3Model []*MarginInterestRateV3Model

// QueryMarginInterestRateV3 g	et the interest rates of the margin lending market over the past 7 days
func (as *ApiService) QueryMarginInterestRateV3(ctx context.Context, currency string) (*ApiResponse, error) {
	p := map[string]string{
		"currency": currency,
	}
	req := NewRequest(http.MethodGet, "/api/v3/project/marketInterestRate", p)
	return as.Call(ctx, req)
}

type LendingPurchaseV3Req struct {
	Currency     string `json:"currency"`
	Size         string `json:"size"`
	InterestRate string `json:"interestRate"`
}

type LendingV3Res struct {
	OrderNo string `json:"orderNo"`
}

// LendingPurchaseV3 Initiate subscriptions of margin lending.
func (as *ApiService) LendingPurchaseV3(ctx context.Context, p *LendingPurchaseV3Req) (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v3/purchase", p)
	return as.Call(ctx, req)
}

type LendingRedeemV3Req struct {
	Currency        string `json:"currency"`
	Size            string `json:"size"`
	PurchaseOrderNo string `json:"purchaseOrderNo"`
}

// LendingRedeemV3 Initiate redemptions of margin lending.
func (as *ApiService) LendingRedeemV3(ctx context.Context, p *LendingRedeemV3Req) (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v3/redeem", p)
	return as.Call(ctx, req)
}

type LendingPurchaseUpdateV3Req struct {
	Currency        string `json:"currency"`
	PurchaseOrderNo string `json:"purchaseOrderNo"`
	InterestRate    string `json:"interestRate"`
}

// LendingPurchaseUpdateV3 update the interest rates of subscription orders, which will take effect at the beginning of the next hour.
func (as *ApiService) LendingPurchaseUpdateV3(ctx context.Context, p *LendingPurchaseUpdateV3Req) (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v3/lend/purchase/update", p)
	return as.Call(ctx, req)
}

type RedemptionOrderV3Model struct {
	Currency        string      `json:"currency"`
	PurchaseOrderNo string      `json:"purchaseOrderNo"`
	RedeemOrderNo   string      `json:"redeemOrderNo"`
	RedeemAmount    json.Number `json:"redeemAmount"`
	ReceiptAmount   json.Number `json:"receiptAmount"`
	ApplyTime       int64       `json:"applyTime"`
	Status          string      `json:"status"`
}
type RedemptionOrdersV3Model []*RedemptionOrderV3Model

// RedemptionOrdersV3 pagination query for the redemption orders.
func (as *ApiService) RedemptionOrdersV3(ctx context.Context, currency, status string, pagination *PaginationParam) (*ApiResponse, error) {
	p := map[string]string{
		"currency": currency,
		"status":   status,
	}
	pagination.ReadParam(p)
	req := NewRequest(http.MethodGet, "/api/v3/redeem/orders", p)
	return as.Call(ctx, req)
}

type SubscriptionOrderV3Model struct {
	Currency        string      `json:"currency"`
	PurchaseOrderNo string      `json:"purchaseOrderNo"`
	PurchaseAmount  json.Number `json:"purchaseAmount"`
	LendAmount      json.Number `json:"lendAmount"`
	RedeemAmount    json.Number `json:"redeemAmount"`
	InterestRate    json.Number `json:"interestRate"`
	IncomeAmount    json.Number `json:"incomeAmount"`
	ApplyTime       int64       `json:"applyTime"`
	Status          string      `json:"status"`
}

type SubscriptionOrdersV3Model []*SubscriptionOrderV3Model

// SubscriptionOrdersV3 pagination query for the subscription orders.
func (as *ApiService) SubscriptionOrdersV3(ctx context.Context, currency, status string, pagination *PaginationParam) (*ApiResponse, error) {
	p := map[string]string{
		"currency": currency,
		"status":   status,
	}
	pagination.ReadParam(p)
	req := NewRequest(http.MethodGet, "/api/v3/purchase/orders", p)
	return as.Call(ctx, req)
}

type MarginSymbolsV3Model struct {
	Items     []*MarginSymbolsV3Model `json:"items"`
	Timestamp int64                   `json:"timestamp"`
}

type MarginSymbolV3Model struct {
	Symbol         string  `json:"symbol"`
	Name           string  `json:"name"`
	EnableTrading  bool    `json:"enableTrading"`
	Market         string  `json:"market"`
	BaseCurrency   string  `json:"baseCurrency"`
	QuoteCurrency  string  `json:"quoteCurrency"`
	BaseIncrement  float64 `json:"baseIncrement"`
	BaseMinSize    float64 `json:"baseMinSize"`
	QuoteIncrement float64 `json:"quoteIncrement"`
	QuoteMinSize   float64 `json:"quoteMinSize"`
	BaseMaxSize    int64   `json:"baseMaxSize"`
	QuoteMaxSize   int     `json:"quoteMaxSize"`
	PriceIncrement float64 `json:"priceIncrement"`
	FeeCurrency    string  `json:"feeCurrency"`
	PriceLimitRate float64 `json:"priceLimitRate"`
	MinFunds       float64 `json:"minFunds"`
}

// MarginSymbolsV3  querying the configuration of cross margin trading pairs.
func (as *ApiService) MarginSymbolsV3(ctx context.Context) (*ApiResponse, error) {
	p := map[string]string{}
	req := NewRequest(http.MethodGet, "/api/v3/margin/symbols", p)
	return as.Call(ctx, req)
}

type UpdateUserLeverageV3Model struct {
	Symbol     string `json:"symbol"`
	Leverage   string `json:"leverage"`
	IsIsolated bool   `json:"isIsolated"`
}

// UpdateUserLeverageV3  modifying the leverage multiplier for cross margin or isolated margin.
func (as *ApiService) UpdateUserLeverageV3(ctx context.Context, p *UpdateUserLeverageV3Model) (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v3/position/update-user-leverage", p)
	return as.Call(ctx, req)
}
