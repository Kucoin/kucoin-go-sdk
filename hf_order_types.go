package kucoin

import (
	"encoding/json"
	"math/big"
)

type HfPlaceOrderRes struct {
	OrderId   string `json:"orderId"`
	ClientOid string `json:"clientOid"`
	Success   bool   `json:"success"`
}

type HfSyncPlaceOrderRes struct {
	OrderId      string      `json:"orderId"`
	ClientOid    string      `json:"clientOid"`
	OrderTime    json.Number `json:"orderTime"`
	OriginSize   string      `json:"originSize"`
	DealSize     string      `json:"dealSize"`
	RemainSize   string      `json:"remainSize"`
	CanceledSize string      `json:"canceledSize"`
	Status       string      `json:"status"`
	MatchTime    json.Number `json:"matchTime"`
}

type HFCreateMultiOrderModel struct {
	ClientOid   string  `json:"clientOid"`
	Symbol      string  `json:"symbol"`
	OrderType   string  `json:"type"`
	TimeInForce string  `json:"timeInForce"`
	Stp         string  `json:"stp"`
	Side        string  `json:"side"`
	Price       string  `json:"price"`
	Size        string  `json:"size"`
	CancelAfter big.Int `json:"cancelAfter"`
	PostOnly    bool    `json:"postOnly"`
	Hidden      bool    `json:"hidden"`
	Iceberg     bool    `json:"iceberg"`
	VisibleSize string  `json:"visibleSize"`
	Tags        string  `json:"tags"`
	Remark      string  `json:"remark"`
}

type HfPlaceMultiOrdersRes []*HfPlaceOrderRes

type HfModifyOrderRes struct {
	NewOrderId string `json:"newOrderId"`
	ClientOid  string `json:"clientOid"`
}

type HfSyncCancelOrderRes struct {
	OrderId      string `json:"orderId"`
	OriginSize   string `json:"originSize"`
	OriginFunds  string `json:"originFunds"`
	DealSize     string `json:"dealSize"`
	RemainSize   string `json:"remainSize"`
	CanceledSize string `json:"canceledSize"`
	Status       string `json:"status"`
}

type HfSyncCancelOrderWithSizeRes struct {
	OrderId    string `json:"orderId"`
	CancelSize string `json:"cancelSize"`
}

type HfOrdersModel []*HfOrderModel

type HfOrderModel struct {
	Id             string      `json:"id"`
	Symbol         string      `json:"symbol"`
	OpType         string      `json:"opType"`
	Type           string      `json:"type"`
	Side           string      `json:"side"`
	Price          string      `json:"price"`
	Size           string      `json:"size"`
	Funds          string      `json:"funds"`
	DealSize       string      `json:"dealSize"`
	DealFunds      string      `json:"dealFunds"`
	Fee            string      `json:"fee"`
	FeeCurrency    string      `json:"feeCurrency"`
	Stp            string      `json:"stp"`
	TimeInForce    string      `json:"timeInForce"`
	PostOnly       bool        `json:"postOnly"`
	Hidden         bool        `json:"hidden"`
	Iceberg        bool        `json:"iceberg"`
	VisibleSize    string      `json:"visibleSize"`
	CancelAfter    int64       `json:"cancelAfter"`
	Channel        string      `json:"channel"`
	ClientOid      string      `json:"clientOid"`
	Remark         string      `json:"remark"`
	Tags           string      `json:"tags"`
	CancelExist    bool        `json:"cancelExist"`
	CreatedAt      json.Number `json:"createdAt"`
	LastUpdatedAt  json.Number `json:"lastUpdatedAt"`
	TradeType      string      `json:"tradeType"`
	InOrderBook    bool        `json:"inOrderBook"`
	Active         bool        `json:"active"`
	CancelledSize  string      `json:"cancelledSize"`
	CancelledFunds string      `json:"cancelledFunds"`
	RemainSize     string      `json:"remainSize"`
	RemainFunds    string      `json:"remainFunds"`
}

type HfAutoCancelSettingRes struct {
	CurrentTime json.Number `json:"currentTime"`
	TriggerTime json.Number `json:"triggerTime"`
}

type AUtoCancelSettingModel struct {
	Timeout     int64       `json:"timeout"`
	Symbols     string      `json:"symbols"`
	CurrentTime json.Number `json:"currentTime"`
	TriggerTime json.Number `json:"triggerTime"`
}

type HfOrderIdModel struct {
	OrderId string `json:"orderId"`
}

type HfClientOidModel struct {
	ClientOid string `json:"clientOid"`
}

type HfTransactionDetailsModel struct {
	LastId json.Number                 `json:"lastId"`
	Items  []*HfTransactionDetailModel `json:"items"`
}

type HfTransactionDetailModel struct {
	Id             json.Number `json:"id"`
	Symbol         string      `json:"symbol"`
	TradeId        json.Number `json:"tradeId"`
	OrderId        string      `json:"orderId"`
	CounterOrderId string      `json:"counterOrderId"`
	Side           string      `json:"side"`
	Liquidity      string      `json:"liquidity"`
	ForceTaker     bool        `json:"forceTaker"`
	Price          string      `json:"price"`
	Size           string      `json:"size"`
	Funds          string      `json:"funds"`
	Fee            string      `json:"fee"`
	FeeRate        string      `json:"feeRate"`
	FeeCurrency    string      `json:"feeCurrency"`
	OrderType      string      `json:"type"`
	Stop           string      `json:"stop"`
	CreatedAt      json.Number `json:"createdAt"`
	TradeType      string      `json:"tradeType"`
}

type HfCancelOrdersResultModel struct {
	SucceedSymbols []string                           `json:"succeedSymbols"`
	FailedSymbols  []*HfCancelOrdersFailedResultModel `json:"failedSymbols"`
}
type HfCancelOrdersFailedResultModel struct {
	Symbol string `json:"symbol"`
	Error  string `json:"error"`
}

type HfPlaceOrderReq struct {
	ClientOid string `json:"clientOid"`
	Symbol    string `json:"symbol"`
	Side      string `json:"side"`
	Stp       string `json:"stp"`
	Tags      string `json:"tags"`
	Remark    string `json:"remark"`

	Price       string `json:"price,omitempty"`
	Size        string `json:"size,omitempty"`
	TimeInForce string `json:"timeInForce,omitempty"`
	CancelAfter int64  `json:"cancelAfter,omitempty"`
	PostOnly    bool   `json:"postOnly,omitempty"`
	Hidden      bool   `json:"hidden,omitempty"`
	Iceberg     bool   `json:"iceberg,omitempty"`
	VisibleSize bool   `json:"visibleSize,omitempty"`

	Funds string `json:"funds,omitempty"`
}

type HfMarginOrderV3Req struct {
	ClientOid  string `json:"clientOid"`
	Symbol     string `json:"symbol"`
	Side       string `json:"side"`
	Type       string `json:"type"`
	Stp        string `json:"stp"`
	IsIsolated bool   `json:"isIsolated"`
	AutoBorrow bool   `json:"autoBorrow"`
	AutoRepay  bool   `json:"autoRepay"`

	Price       string `json:"price,omitempty"`
	Size        string `json:"size,omitempty"`
	TimeInForce string `json:"timeInForce,omitempty"`
	CancelAfter int64  `json:"cancelAfter,omitempty"`
	PostOnly    bool   `json:"postOnly,omitempty"`
	Hidden      bool   `json:"hidden,omitempty"`
	Iceberg     bool   `json:"iceberg,omitempty"`
	VisibleSize bool   `json:"visibleSize,omitempty"`

	Funds string `json:"funds,omitempty"`
}

type HfMarginOrderV3Resp struct {
	OrderNo     string `json:"orderNo"`
	BorrowSize  string `json:"borrowSize"`
	LoanApplyId string `json:"loanApplyId"`
}

type HfCancelMarinOrderV3Req struct {
	OrderId string `json:"orderId"`
	Symbol  string `json:"symbol"`
}

type HfCancelMarinOrderV3Resp struct {
	OrderId string `json:"orderId"`
}

type HfCancelClientMarinOrderV3Req struct {
	ClientOid string `json:"clientOid"`
	Symbol    string `json:"symbol"`
}

type HfCancelClientMarinOrderV3Resp struct {
	ClientOid string `json:"clientOid"`
}

type HfCancelAllMarginOrdersV3Req struct {
	TradeType string `json:"tradeType" url:"tradeType"`
	Symbol    string `json:"symbol" url:"symbol"`
}

type HfCancelAllMarginOrdersV3Resp string

type HFMarginActiveSymbolsModel struct {
	SymbolSize int32    `json:"symbolSize"`
	Symbols    []string `json:"symbols"`
}

type HfMarinActiveOrdersV3Req struct {
	TradeType string `url:"tradeType"`
	Symbol    string `url:"symbol"`
}

type MarginOrderV3Model struct {
	Id            string `json:"id"`
	Symbol        string `json:"symbol"`
	OpType        string `json:"opType"`
	Type          string `json:"type"`
	Side          string `json:"side"`
	Price         string `json:"price"`
	Size          string `json:"size"`
	Funds         string `json:"funds"`
	DealFunds     string `json:"dealFunds"`
	DealSize      string `json:"dealSize"`
	Fee           string `json:"fee"`
	FeeCurrency   string `json:"feeCurrency"`
	Stp           string `json:"stp"`
	TimeInForce   string `json:"timeInForce"`
	PostOnly      bool   `json:"postOnly"`
	Hidden        bool   `json:"hidden"`
	Iceberg       bool   `json:"iceberg"`
	VisibleSize   string `json:"visibleSize"`
	CancelAfter   int    `json:"cancelAfter"`
	Channel       string `json:"channel"`
	ClientOid     string `json:"clientOid"`
	Remark        string `json:"remark"`
	Tags          string `json:"tags"`
	Active        bool   `json:"active"`
	InOrderBook   bool   `json:"inOrderBook"`
	CancelExist   bool   `json:"cancelExist"`
	CreatedAt     int64  `json:"createdAt"`
	LastUpdatedAt int64  `json:"lastUpdatedAt"`
	TradeType     string `json:"tradeType"`
}

type HfMarinActiveOrdersV3Resp []*MarginOrderV3Model

type HfMarinDoneOrdersV3Req struct {
	TradeType string `url:"tradeType"`
	Symbol    string `url:"symbol"`
	Side      string `url:"side,omitempty"`
	Type      string `url:"type,omitempty"`
	StartAt   int64  `url:"startAt,omitempty"`
	EndAt     int64  `url:"endAt,omitempty"`
	LastId    int64  `url:"lastId,omitempty"`
	Limit     int    `url:"limit,omitempty"`
}

type HfMarinDoneOrdersV3Resp struct {
	Items  []*MarginFillModel `json:"items"`
	LastId int64              `json:"lastId"`
}

type MarginFillModel struct {
	ID            int64  `json:"id"`
	Symbol        string `json:"symbol"`
	OpType        string `json:"opType"`
	Type          string `json:"type"`
	Side          string `json:"side"`
	Price         string `json:"price"`
	Size          string `json:"size"`
	Funds         string `json:"funds"`
	DealFunds     string `json:"dealFunds"`
	DealSize      string `json:"dealSize"`
	Fee           string `json:"fee"`
	FeeCurrency   string `json:"feeCurrency"`
	STP           string `json:"stp"`
	TimeInForce   string `json:"timeInForce"`
	PostOnly      bool   `json:"postOnly"`
	Hidden        bool   `json:"hidden"`
	Iceberg       bool   `json:"iceberg"`
	VisibleSize   string `json:"visibleSize"`
	CancelAfter   int64  `json:"cancelAfter"`
	Channel       string `json:"channel"`
	ClientOid     string `json:"clientOid"`
	Remark        string `json:"remark"`
	Tags          string `json:"tags"`
	Active        bool   `json:"active"`
	InOrderBook   bool   `json:"inOrderBook"`
	CancelExist   bool   `json:"cancelExist"`
	CreatedAt     int64  `json:"createdAt"`
	LastUpdatedAt int64  `json:"lastUpdatedAt"`
	TradeType     string `json:"tradeType"`
}

type HfMarinOrderV3Req struct {
	OrderId string
	Symbol  string
}

type HfMarinOrderV3Resp struct {
	MarginFillModel
}

type HfMarinClientOrderV3Req struct {
	ClientOid string
	Symbol    string
}

type HfMarinClientOrderV3Resp struct {
	MarginFillModel
}

type HfMarinFillsV3Req struct {
	Symbol    string `url:"symbol"`
	TradeType string `url:"tradeType"`
	OrderId   string `url:"orderId,omitempty"`
	Side      string `url:"side,omitempty"`
	Type      string `url:"type,omitempty"`
	StartAt   int64  `url:"startAt,omitempty"`
	EndAt     int64  `url:"endAt,omitempty"`
	LastId    int64  `url:"lastId,omitempty"`
	Limit     int    `url:"limit,omitempty"`
}

type HfMarinFillsV3Resp struct {
	Items  []*MarginFillModel `json:"items"`
	LastId int64              `json:"lastId"`
}
