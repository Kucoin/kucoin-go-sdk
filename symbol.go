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

type SymbolsModel []*SymbolModel

func (as *ApiService) Symbols() (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/symbols", nil)
	return as.Call(req)
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

func (as *ApiService) Ticker(symbol string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/market/orderbook/level1", map[string]string{"symbol": symbol})
	return as.Call(req)
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

func (as *ApiService) PartOrderBook(symbol string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/market/orderbook/level2_100", map[string]string{"symbol": symbol})
	return as.Call(req)
}

type FullOrderBookModel struct {
	Sequence string     `json:"sequence"`
	Bids     [][]string `json:"bids"`
	Asks     [][]string `json:"asks"`
}

func (as *ApiService) AggregatedFullOrderBook(symbol string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/market/orderbook/level2", map[string]string{"symbol": symbol})
	return as.Call(req)
}

func (as *ApiService) AtomicFullOrderBook(symbol string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/market/orderbook/level3", map[string]string{"symbol": symbol})
	return as.Call(req)
}

type TradeHistoryModel struct {
	Sequence string `json:"sequence"`
	Price    string `json:"price"`
	Size     string `json:"size"`
	Side     string `json:"side"`
	Time     int64  `json:"time"`
}

type TradeHistoriesModel []*TradeHistoryModel

func (as *ApiService) TradeHistories(symbol string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/market/histories", map[string]string{"symbol": symbol})
	return as.Call(req)
}

type HistoricRateModel []string
type HistoricRatesModel []*HistoricRateModel

func (as *ApiService) HistoricRates(symbol string, startAt, endAt int64, typo string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/market/candles", map[string]string{
		"symbol":  symbol,
		"startAt": strconv.FormatInt(startAt, 10),
		"endAt":   strconv.FormatInt(endAt, 10),
		"type":    typo,
	})
	return as.Call(req)
}
