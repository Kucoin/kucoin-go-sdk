package kucoin

import "net/http"

// A FillModel represents the structure of fill.
type FillModel struct {
	Symbol         string `json:"symbol"`
	TradeId        string `json:"tradeId"`
	OrderId        string `json:"orderId"`
	CounterOrderId string `json:"counterOrderId"`
	Side           string `json:"side"`
	Liquidity      string `json:"liquidity"`
	ForceTaker     bool   `json:"forceTaker"`
	Price          string `json:"price"`
	Size           string `json:"size"`
	Funds          string `json:"funds"`
	Fee            string `json:"fee"`
	FeeRate        string `json:"feeRate"`
	FeeCurrency    string `json:"feeCurrency"`
	Stop           string `json:"stop"`
	Type           string `json:"type"`
	CreatedAt      int64  `json:"createdAt"`
	TradeType      string `json:"tradeType"`
}

// A FillsModel is the set of *FillModel.
type FillsModel []*FillModel

// Fills returns a list of recent fills.
func (as *ApiService) Fills(params map[string]string, pagination *PaginationParam) (*ApiResponse, error) {
	pagination.ReadParam(params)
	req := NewRequest(http.MethodGet, "/api/v1/fills", params)
	return as.Call(req)
}

// RecentFills returns the recent fills of the latest transactions within 24 hours.
func (as *ApiService) RecentFills() (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/limit/fills", nil)
	return as.Call(req)
}
