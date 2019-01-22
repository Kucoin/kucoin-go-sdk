package kucoin

import (
	"net/http"
	"strconv"
)

type SymbolModel struct {
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

type SymbolModels []*SymbolModel

func (as *ApiService) Symbols() (SymbolModels, error) {
	req := NewRequest(http.MethodGet, "/api/v1/symbols", nil)
	rsp, err := as.Call(req)
	if err != nil {
		return nil, err
	}
	type Data struct {
		Data SymbolModels `json:"data"`
	}
	v := &Data{}
	rsp.ApiData(v)
	return v.Data, nil
}

type TickerModel struct {
	Sequence    string `json:"sequence"`
	Price       string `json:"price"`
	Size        string `json:"size"`
	BestBid     string `json:"bestBid"`
	BestBidSize string `json:"bestBidSize"`
	BestAsk     string `json:"bestAsk"`
	BestAskSize string `json:"bestAskSize"`
}

func (as *ApiService) Ticker(symbol string) (*TickerModel, error) {
	req := NewRequest(http.MethodGet, "/api/v1/market/orderbook/level1", map[string]string{"symbol": symbol})
	rsp, err := as.Call(req)
	if err != nil {
		return nil, err
	}
	type Data struct {
		Data *TickerModel `json:"data"`
	}
	v := &Data{}
	rsp.ApiData(v)
	return v.Data, nil
}

type PartOrderBookModel struct {
	Symbol      string `json:"symbol"`
	ChangeRate  string `json:"changeRate"`
	ChangePrice string `json:"changePrice"`
	Open        string `json:"open"`
	Close       string `json:"close"`
	High        string `json:"high"`
	Low         string `json:"low"`
	Vol         string `json:"vol"`
	VolValue    string `json:"volValue"`
}

func (as *ApiService) PartOrderBook(symbol string) (*PartOrderBookModel, error) {
	req := NewRequest(http.MethodGet, "/api/v1/market/orderbook/level2_100", map[string]string{"symbol": symbol})
	rsp, err := as.Call(req)
	if err != nil {
		return nil, err
	}
	type Data struct {
		Data *PartOrderBookModel `json:"data"`
	}
	v := &Data{}
	rsp.ApiData(v)
	return v.Data, nil
}

type FullOrderBookModel struct {
	Sequence string     `json:"sequence"`
	Bids     [][]string `json:"bids"`
	Asks     [][]string `json:"asks"`
}

func (as *ApiService) AggregatedFullOrderBook(symbol string) (*FullOrderBookModel, error) {
	req := NewRequest(http.MethodGet, "/api/v1/market/orderbook/level2", map[string]string{"symbol": symbol})
	rsp, err := as.Call(req)
	if err != nil {
		return nil, err
	}
	type Data struct {
		Data *FullOrderBookModel `json:"data"`
	}
	v := &Data{}
	rsp.ApiData(v)
	return v.Data, nil
}

func (as *ApiService) AtomicFullOrderBook(symbol string) (*FullOrderBookModel, error) {
	req := NewRequest(http.MethodGet, "/api/v1/market/orderbook/level3", map[string]string{"symbol": symbol})
	rsp, err := as.Call(req)
	if err != nil {
		return nil, err
	}
	type Data struct {
		Data *FullOrderBookModel `json:"data"`
	}
	v := &Data{}
	rsp.ApiData(v)
	return v.Data, nil
}

type TradeHistoryModel struct {
	Sequence string `json:"sequence"`
	Price    string `json:"price"`
	Size     string `json:"size"`
	Side     string `json:"side"`
	Time     int64  `json:"time"`
}

type TradeHistoriesModel []*TradeHistoryModel

func (as *ApiService) TradeHistories(symbol string) (TradeHistoriesModel, error) {
	req := NewRequest(http.MethodGet, "/api/v1/market/histories", map[string]string{"symbol": symbol})
	rsp, err := as.Call(req)
	if err != nil {
		return nil, err
	}
	type Data struct {
		Data TradeHistoriesModel `json:"data"`
	}
	v := &Data{}
	rsp.ApiData(v)
	return v.Data, nil
}

type HistoricRateModel []string
type HistoricRatesModel []*HistoricRateModel

func (as *ApiService) HistoricRates(symbol string, startAt, endAt int64, typ string) (HistoricRatesModel, error) {
	params := map[string]string{"symbol": symbol}
	if startAt > 0 {
		params["startAt"] = strconv.FormatInt(startAt, 10)
	}
	if endAt > 0 {
		params["endAt"] = strconv.FormatInt(endAt, 10)
	}
	if typ != "" {
		params["type"] = typ
	}
	req := NewRequest(http.MethodGet, "/api/v1/market/candles", params)
	rsp, err := as.Call(req)
	if err != nil {
		return nil, err
	}
	type Data struct {
		Data HistoricRatesModel `json:"data"`
	}
	v := &Data{}
	rsp.ApiData(v)
	return v.Data, nil
}
