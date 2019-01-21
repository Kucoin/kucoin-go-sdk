package kucoin

import (
	"net/http"
)

type Symbol struct {
	Symbol         string `json:"symbol"`
	Name           string `json:"name"`
	BaseCurrency   string `json:"baseCurrency"`
	QuoteCurrency  string `json:"quoteCurrency"`
	BaseMinSize    string `json:"baseMinSize"`
	QuoteMinSize   string `json:"quoteMinSize"`
	BaseMaxSize    string `json:"baseMaxSize"`
	QuoteMaxSize   string `json:"quoteMaxSize"`
	BaseIncrement  string `json:"baseIncrement"`
	QuoteIncrement string `json:"quoteIncrement"`
	PriceIncrement string `json:"priceIncrement"`
	EnableTrading  bool   `json:"enableTrading"`
}

type Symbols []*Symbol

func SymbolList() (Symbols, error) {
	req := NewRequest(http.MethodGet, "/api/v1/symbols", nil)
	rsp, err := Api.Call(req)
	if err != nil {
		return nil, err
	}
	type Data struct {
		Data Symbols `json:"data"`
	}
	v := &Data{}
	rsp.ApiData(v)
	return v.Data, nil
}

type Ticker struct {
	Sequence    string `json:"sequence"`
	Price       string `json:"price"`
	Size        string `json:"size"`
	BestBid     string `json:"bestBid"`
	BestBidSize string `json:"bestBidSize"`
	BestAsk     string `json:"bestAsk"`
	BestAskSize string `json:"bestAskSize"`
}

func TickerDetail(symbol string) (*Ticker, error) {
	req := NewRequest(http.MethodGet, "/api/v1/market/orderbook/level1", map[string]string{"symbol": symbol})
	rsp, err := Api.Call(req)
	if err != nil {
		return nil, err
	}
	type Data struct {
		Data *Ticker `json:"data"`
	}
	v := &Data{}
	rsp.ApiData(v)
	return v.Data, nil
}
