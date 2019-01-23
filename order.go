package kucoin

import "net/http"

type CreateOrderResultModel struct {
	OrderId string `json:"orderId"`
}

func (as *ApiService) CreateOrder(params map[string]string) (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v1/orders", params)
	return as.call(req)
}

type CancelOrderResultModel struct {
	CancelledOrderIds []string `json:"cancelledOrderIds"`
}

func (as *ApiService) CancelOrder(orderId string) (*ApiResponse, error) {
	req := NewRequest(http.MethodDelete, "/api/v1/orders/"+orderId, nil)
	return as.call(req)
}

func (as *ApiService) CancelOrders(symbol string) (*ApiResponse, error) {
	p := map[string]string{}
	if symbol != "" {
		p["symbol"] = symbol
	}
	req := NewRequest(http.MethodDelete, "/api/v1/orders", p)
	return as.call(req)
}

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
	IceBerge      bool   `json:"iceberge"`
	VisibleSize   string `json:"visibleSize"`
	CancelAfter   string `json:"cancelAfter"`
	Channel       string `json:"channel"`
	ClientOid     string `json:"clientOid"`
	Remark        string `json:"remark"`
	Tags          string `json:"tags"`
	IsActive      bool   `json:"isActive"`
	CancelExist   bool   `json:"cancelExist"`
	CreatedAt     int64  `json:"createdAt"`
}

type OrdersModel []OrderModel

func (as *ApiService) Orders(params map[string]string, pagination *PaginationParam) (*ApiResponse, error) {
	pagination.ReadParam(&params)
	req := NewRequest(http.MethodGet, "/api/v1/orders", params)
	return as.call(req)
}

func (as *ApiService) Order(orderId string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/orders/"+orderId, nil)
	return as.call(req)
}
