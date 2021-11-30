package kucoin

import "net/http"

// A CreateOrderModel is the input parameter of CreateOrder().
type CreateOrderModel struct {
	// BASE PARAMETERS
	ClientOid string `json:"clientOid"`
	Side      string `json:"side"`
	Symbol    string `json:"symbol,omitempty"`
	Type      string `json:"type,omitempty"`
	Remark    string `json:"remark,omitempty"`
	Stop      string `json:"stop,omitempty"`
	StopPrice string `json:"stopPrice,omitempty"`
	STP       string `json:"stp,omitempty"`
	TradeType string `json:"tradeType,omitempty"`

	// LIMIT ORDER PARAMETERS
	Price       string `json:"price,omitempty"`
	Size        string `json:"size,omitempty"`
	TimeInForce string `json:"timeInForce,omitempty"`
	CancelAfter int64 `json:"cancelAfter,omitempty"`
	PostOnly    bool   `json:"postOnly,omitempty"`
	Hidden      bool   `json:"hidden,omitempty"`
	IceBerg     bool   `json:"iceberg,omitempty"`
	VisibleSize string `json:"visibleSize,omitempty"`

	// MARKET ORDER PARAMETERS
	// Size  string `json:"size"`
	Funds string `json:"funds,omitempty"`

	// MARGIN ORDER PARAMETERS
	MarginMode string `json:"marginMode,omitempty"`
	AutoBorrow bool   `json:"autoBorrow,omitempty"`
}

// A CreateOrderResultModel represents the result of CreateOrder().
type CreateOrderResultModel struct {
	OrderId string `json:"orderId"`
}

// CreateOrder places a new order.
func (as *ApiService) CreateOrder(o *CreateOrderModel) (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v1/orders", o)
	return as.Call(req)
}

// A CreateMultiOrderResultModel represents the result of CreateMultiOrder().
type CreateMultiOrderResultModel struct {
	Data OrdersModel `json:"data"`
}

// CreateMultiOrder places bulk orders.
func (as *ApiService) CreateMultiOrder(symbol string, orders []*CreateOrderModel) (*ApiResponse, error) {
	params := map[string]interface{}{
		"symbol":    symbol,
		"orderList": orders,
	}
	req := NewRequest(http.MethodPost, "/api/v1/orders/multi", params)
	return as.Call(req)
}

// A CancelOrderResultModel represents the result of CancelOrder().
type CancelOrderResultModel struct {
	CancelledOrderIds []string `json:"cancelledOrderIds"`
}

// CancelOrder cancels a previously placed order.
func (as *ApiService) CancelOrder(orderId string) (*ApiResponse, error) {
	req := NewRequest(http.MethodDelete, "/api/v1/orders/"+orderId, nil)
	return as.Call(req)
}

// A CancelOrderByClientResultModel represents the result of CancelOrderByClient().
type CancelOrderByClientResultModel struct {
	CancelledOrderId string `json:"cancelledOrderId"`
	ClientOid        string `json:"clientOid"`
}

// CancelOrderByClient cancels a previously placed order by client ID.
func (as *ApiService) CancelOrderByClient(clientOid string) (*ApiResponse, error) {
	req := NewRequest(http.MethodDelete, "/api/v1/order/client-order/"+clientOid, nil)
	return as.Call(req)
}

// CancelOrders cancels all orders of the symbol.
// With best effort, cancel all open orders. The response is a list of ids of the canceled orders.
func (as *ApiService) CancelOrders(p map[string]string) (*ApiResponse, error) {
	req := NewRequest(http.MethodDelete, "/api/v1/orders", p)
	return as.Call(req)
}

// An OrderModel represents an order.
type OrderModel struct {
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
	Stop          string `json:"stop"`
	StopTriggered bool   `json:"stopTriggered"`
	StopPrice     string `json:"stopPrice"`
	TimeInForce   string `json:"timeInForce"`
	PostOnly      bool   `json:"postOnly"`
	Hidden        bool   `json:"hidden"`
	IceBerg       bool   `json:"iceberg"`
	VisibleSize   string `json:"visibleSize"`
	CancelAfter   int64 `json:"cancelAfter"`
	Channel       string `json:"channel"`
	ClientOid     string `json:"clientOid"`
	Remark        string `json:"remark"`
	Tags          string `json:"tags"`
	IsActive      bool   `json:"isActive"`
	CancelExist   bool   `json:"cancelExist"`
	CreatedAt     int64  `json:"createdAt"`
	TradeType     string `json:"tradeType"`
}

// A OrdersModel is the set of *OrderModel.
type OrdersModel []*OrderModel

// Orders returns a list your current orders.
func (as *ApiService) Orders(params map[string]string, pagination *PaginationParam) (*ApiResponse, error) {
	pagination.ReadParam(params)
	req := NewRequest(http.MethodGet, "/api/v1/orders", params)
	return as.Call(req)
}

// A V1OrderModel represents a v1 order.
type V1OrderModel struct {
	Symbol    string `json:"symbol"`
	DealPrice string `json:"dealPrice"`
	DealValue string `json:"dealValue"`
	Amount    string `json:"amount"`
	Fee       string `json:"fee"`
	Side      string `json:"side"`
	CreatedAt int64  `json:"createdAt"`
}

// A V1OrdersModel is the set of *V1OrderModel.
type V1OrdersModel []*V1OrderModel

// V1Orders returns a list of v1 historical orders.
func (as *ApiService) V1Orders(params map[string]string, pagination *PaginationParam) (*ApiResponse, error) {
	pagination.ReadParam(params)
	req := NewRequest(http.MethodGet, "/api/v1/hist-orders", params)
	return as.Call(req)
}

// Order returns a single order by order id.
func (as *ApiService) Order(orderId string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/orders/"+orderId, nil)
	return as.Call(req)
}

// OrderByClient returns a single order by client id.
func (as *ApiService) OrderByClient(clientOid string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/order/client-order/"+clientOid, nil)
	return as.Call(req)
}

// RecentOrders returns the recent orders of the latest transactions within 24 hours.
func (as *ApiService) RecentOrders() (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/limit/orders", nil)
	return as.Call(req)
}

// CreateStopOrder places a new stop-order.
func (as *ApiService) CreateStopOrder(o *CreateOrderModel) (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v1/stop-order", o)
	return as.Call(req)
}

// CancelStopOrder cancels a previously placed stop-order.
func (as *ApiService) CancelStopOrder(orderId string) (*ApiResponse, error) {
	req := NewRequest(http.MethodDelete, "/api/v1/stop-order/"+orderId, nil)
	return as.Call(req)
}

// CancelStopOrderByClientModel returns Model of CancelStopOrderByClient API
type CancelStopOrderByClientModel struct {
	CancelledOrderId string `json:"cancelledOrderId"`
	ClientOid        string `json:"clientOid"`
}

// CancelStopOrderByClient cancels a previously placed stop-order by client ID.
func (as *ApiService) CancelStopOrderByClient(clientOid string, p map[string]string) (*ApiResponse, error) {
	p["clientOid"] = clientOid

	req := NewRequest(http.MethodDelete, "/api/v1/stop-order/cancelOrderByClientOid", p)
	return as.Call(req)
}

// StopOrderModel RESPONSES of StopOrder
type StopOrderModel struct {
	Id              string `json:"id"`
	Symbol          string `json:"symbol"`
	UserId          string `json:"userId"`
	Status          string `json:"status"`
	Type            string `json:"type"`
	Side            string `json:"side"`
	Price           string `json:"price"`
	Size            string `json:"size"`
	Funds           string `json:"funds"`
	Stp             string `json:"stp"`
	TimeInForce     string `json:"timeInForce"`
	CancelAfter     int64  `json:"cancelAfter"`
	PostOnly        bool   `json:"postOnly"`
	Hidden          bool   `json:"hidden"`
	IceBerg         bool   `json:"iceberg"`
	VisibleSize     string `json:"visibleSize"`
	Channel         string `json:"channel"`
	ClientOid       string `json:"clientOid"`
	Remark          string `json:"remark"`
	Tags            string `json:"tags"`
	OrderTime       int64  `json:"orderTime"`
	DomainId        string `json:"domainId"`
	TradeSource     string `json:"tradeSource"`
	TradeType       string `json:"tradeType"`
	FeeCurrency     string `json:"feeCurrency"`
	TakerFeeRate    string `json:"takerFeeRate"`
	MakerFeeRate    string `json:"makerFeeRate"`
	CreatedAt       int64  `json:"createdAt"`
	Stop            string `json:"stop"`
	StopTriggerTime string `json:"stopTriggerTime"`
	StopPrice       string `json:"stopPrice"`
}

// StopOrder returns a single order by stop-order id.
func (as *ApiService) StopOrder(orderId string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/stop-order/"+orderId, nil)
	return as.Call(req)
}

// StopOrderListModel StopOrderByClient model
type StopOrderListModel []*StopOrderModel

// StopOrderByClient returns a single stop-order by client id.
func (as *ApiService) StopOrderByClient(clientOid string, p map[string]string) (*ApiResponse, error) {
	p["clientOid"] = clientOid

	req := NewRequest(http.MethodGet, "/api/v1/stop-order/queryOrderByClientOid", p)
	return as.Call(req)
}

// StopOrders returns a list your current orders.
func (as *ApiService) StopOrders(params map[string]string, pagination *PaginationParam) (*ApiResponse, error) {
	pagination.ReadParam(params)
	req := NewRequest(http.MethodGet, "/api/v1/stop-order", params)
	return as.Call(req)
}

// CancelStopOrderBy returns a list your current orders.
func (as *ApiService) CancelStopOrderBy(params map[string]string) (*ApiResponse, error) {
	req := NewRequest(http.MethodDelete, "/api/v1/stop-order/cancel", params)
	return as.Call(req)
}

// CreateMarginOrder places a new margin order.
func (as *ApiService) CreateMarginOrder(o *CreateOrderModel) (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v1/margin/order", o)
	return as.Call(req)
}
