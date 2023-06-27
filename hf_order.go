package kucoin

import (
	"encoding/json"
	"math/big"
	"net/http"
)

// HfPlaceOrder There are two types of orders:
// (limit) order: set price and quantity for the transaction.
// (market) order : set amount or quantity for the transaction.
func (as *ApiService) HfPlaceOrder(params map[string]string) (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v1/hf/orders", params)
	return as.Call(req)
}

type HfPlaceOrderRes struct {
	OrderId string `json:"orderId"`
	Success bool   `json:"success"`
}

// HfSyncPlaceOrder The difference between this interface
// and "Place hf order" is that this interface will synchronously
// return the order information after the order matching is completed.
// For higher latency requirements, please select the "Place hf order" interface.
// If there is a requirement for returning data integrity, please select this interface
func (as *ApiService) HfSyncPlaceOrder(params map[string]string) (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v1/hf/orders/sync", params)
	return as.Call(req)
}

type HfSyncPlaceOrderRes struct {
	OrderId      string      `json:"orderId"`
	OrderTime    json.Number `json:"orderTime"`
	OriginSize   string      `json:"originSize"`
	DealSize     string      `json:"dealSize"`
	RemainSize   string      `json:"remainSize"`
	CanceledSize string      `json:"canceledSize"`
	Status       string      `json:"status"`
	MatchTime    json.Number `json:"matchTime"`
}

// HfPlaceMultiOrders This endpoint supports sequential batch order placement from a single endpoint.
// A maximum of 5orders can be placed simultaneously.
// The order types must be limit orders of the same trading pair
// （this endpoint currently only supports spot trading and does not support margin trading）
func (as *ApiService) HfPlaceMultiOrders(orders []*HFCreateMultiOrderModel) (*ApiResponse, error) {
	p := map[string]interface{}{
		"orderList": orders,
	}
	req := NewRequest(http.MethodPost, "/api/v1/hf/orders/multi", p)
	return as.Call(req)
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

// HfSyncPlaceMultiOrders The request parameters of this interface
// are the same as those of the "Sync place multiple hf orders" interface
// The difference between this interface and "Sync place multiple hf orders" is that
// this interface will synchronously return the order information after the order matching is completed.
func (as *ApiService) HfSyncPlaceMultiOrders(orders []*HFCreateMultiOrderModel) (*ApiResponse, error) {
	p := map[string]interface{}{
		"orderList": orders,
	}
	req := NewRequest(http.MethodPost, "/api/v1/hf/orders/multi/sync", p)
	return as.Call(req)
}

type HfSyncPlaceMultiOrdersRes []*HfSyncPlaceOrderRes

// HfModifyOrder
// This interface can modify the price and quantity of the order according to orderId or clientOid.
func (as *ApiService) HfModifyOrder(params map[string]string) (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v1/hf/orders/alter", params)
	return as.Call(req)
}

type HfModifyOrderRes struct {
	NewOrderId string `json:"newOrderId"`
}

// HfCancelOrder This endpoint can be used to cancel a high-frequency order by orderId.
func (as *ApiService) HfCancelOrder(orderId, symbol string) (*ApiResponse, error) {
	p := map[string]string{
		"symbol": symbol,
	}
	req := NewRequest(http.MethodDelete, "/api/v1/hf/orders/"+orderId, p)
	return as.Call(req)
}

// HfSyncCancelOrder The difference between this interface and "Cancel orders by orderId" is that
// this interface will synchronously return the order information after the order canceling is completed.
func (as *ApiService) HfSyncCancelOrder(orderId, symbol string) (*ApiResponse, error) {
	p := map[string]string{
		"symbol": symbol,
	}
	req := NewRequest(http.MethodDelete, "/api/v1/hf/orders/sync/"+orderId, p)
	return as.Call(req)
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

// HfCancelOrderByClientId This endpoint sends out a request to cancel a high-frequency order using clientOid.
func (as *ApiService) HfCancelOrderByClientId(clientOid, symbol string) (*ApiResponse, error) {
	p := map[string]string{
		"symbol": symbol,
	}
	req := NewRequest(http.MethodDelete, "/api/v1/hf/orders/client-order/"+clientOid, p)
	return as.Call(req)
}

// HfSyncCancelOrderByClientId The difference between this interface and "Cancellation of order by clientOid"
// is that this interface will synchronously return the order information after the order canceling is completed.
func (as *ApiService) HfSyncCancelOrderByClientId(clientOid, symbol string) (*ApiResponse, error) {
	p := map[string]string{
		"symbol": symbol,
	}
	req := NewRequest(http.MethodDelete, "/api/v1/hf/orders/sync/client-order/"+clientOid, p)
	return as.Call(req)
}

// HfSyncCancelOrderWithSize This interface can cancel the specified quantity of the order according to the orderId.
func (as *ApiService) HfSyncCancelOrderWithSize(orderId, symbol, cancelSize string) (*ApiResponse, error) {
	p := map[string]string{
		"symbol":     symbol,
		"cancelSize": cancelSize,
	}
	req := NewRequest(http.MethodDelete, "/api/v1/hf/orders/cancel/"+orderId, p)
	return as.Call(req)
}

type HfSyncCancelOrderWithSizeRes struct {
	OrderId    string `json:"orderId"`
	CancelSize string `json:"cancelSize"`
}

// HfSyncCancelAllOrders his endpoint allows cancellation of all orders related to a specific trading pair
// with a status of open
// (including all orders pertaining to high-frequency trading accounts and non-high-frequency trading accounts)
func (as *ApiService) HfSyncCancelAllOrders(symbol string) (*ApiResponse, error) {
	p := map[string]string{
		"symbol": symbol,
	}
	req := NewRequest(http.MethodDelete, "/api/v1/hf/orders", p)
	return as.Call(req)
}

// HfObtainActiveOrders This endpoint obtains a list of all active HF orders.
// The return data is sorted in descending order based on the latest update times.
func (as *ApiService) HfObtainActiveOrders(symbol string) (*ApiResponse, error) {
	p := map[string]string{
		"symbol": symbol,
	}
	req := NewRequest(http.MethodGet, "/api/v1/hf/orders/active", p)
	return as.Call(req)
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

// HfObtainActiveSymbols This interface can query all trading pairs that the user has active orders
func (as *ApiService) HfObtainActiveSymbols() (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/hf/orders/active/symbols", nil)
	return as.Call(req)
}

type HfSymbolsModel struct {
	Symbols []string `json:"symbols"`
}

// HfObtainFilledOrders This endpoint obtains a list of filled HF orders and returns paginated data.
// The returned data is sorted in descending order based on the latest order update times.
func (as *ApiService) HfObtainFilledOrders(p map[string]string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/hf/orders/done", p)
	return as.Call(req)
}

type HfFilledOrdersModel struct {
	LastId json.Number     `json:"lastId"`
	Items  []*HfOrderModel `json:"items"`
}

// HfOrderDetail This endpoint can be used to obtain information for a single HF order using the order id.
func (as *ApiService) HfOrderDetail(orderId, symbol string) (*ApiResponse, error) {
	p := map[string]string{
		"symbol": symbol,
	}
	req := NewRequest(http.MethodGet, "/api/v1/hf/orders/"+orderId, p)
	return as.Call(req)
}

// HfOrderDetailByClientOid The endpoint can be used to obtain information about a single order using clientOid.
// If the order does not exist, then there will be a prompt saying that the order does not exist.
func (as *ApiService) HfOrderDetailByClientOid(clientOid, symbol string) (*ApiResponse, error) {
	p := map[string]string{
		"symbol": symbol,
	}
	req := NewRequest(http.MethodGet, "/api/v1/hf/orders/client-order/"+clientOid, p)
	return as.Call(req)
}

// HfAutoCancelSetting  automatically cancel all orders of the set trading pair after the specified time.
// If this interface is not called again for renewal or cancellation before the set time,
// the system will help the user to cancel the order of the corresponding trading pair.
// otherwise it will not.
func (as *ApiService) HfAutoCancelSetting(timeout int64, symbol string) (*ApiResponse, error) {
	p := map[string]interface{}{
		"symbol":  symbol,
		"timeout": timeout,
	}
	req := NewRequest(http.MethodPost, "/api/v1/hf/orders/dead-cancel-all", p)
	return as.Call(req)
}

type HfAutoCancelSettingRes struct {
	CurrentTime json.Number `json:"currentTime"`
	TriggerTime json.Number `json:"triggerTime"`
}

// HfQueryAutoCancelSetting  Through this interface, you can query the settings of automatic order cancellation
func (as *ApiService) HfQueryAutoCancelSetting() (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/hf/orders/dead-cancel-all/query", nil)
	return as.Call(req)
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

// HfTransactionDetails This endpoint can be used to obtain a list of the latest HF transaction details.
// The returned results are paginated. The data is sorted in descending order according to time.
func (as *ApiService) HfTransactionDetails(p map[string]string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/hf/fills", p)
	return as.Call(req)
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
