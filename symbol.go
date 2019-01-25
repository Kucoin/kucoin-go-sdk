package kucoin

import (
	"net/http"
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

// Symbols returns a list of available currency pairs for trading.
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

// Ticker returns the ticker include only the inside (i.e. best) bid and ask data, last price and last trade size.
func (as *ApiService) Ticker(symbol string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/market/orderbook/level1", map[string]string{"symbol": symbol})
	return as.Call(req)
}

type Stats24hrModel struct {
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

// Stats24hr returns 24 hr stats for the symbol. volume is in base currency units. open, high, low are in quote currency units.
func (as *ApiService) Stats24hr(symbol string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/market/stats", map[string]string{"symbol": symbol})
	return as.Call(req)
}

// Markets returns the transaction currencies for the entire trading market.
func (as *ApiService) Markets() (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/market", nil)
	return as.Call(req)
}

type PartOrderBookModel struct {
	Sequence string     `json:"sequence"`
	Bids     [][]string `json:"bids"`
	Asks     [][]string `json:"asks"`
}

// PartOrderBook returns a list of open orders(aggregated) for a symbol.
func (as *ApiService) PartOrderBook(symbol string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/market/orderbook/level2_100", map[string]string{"symbol": symbol})
	return as.Call(req)
}

type FullOrderBookModel struct {
	Sequence string     `json:"sequence"`
	Bids     [][]string `json:"bids"`
	Asks     [][]string `json:"asks"`
}

// AggregatedFullOrderBook returns a list of open orders(aggregated) for a symbol.
func (as *ApiService) AggregatedFullOrderBook(symbol string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/market/orderbook/level2", map[string]string{"symbol": symbol})
	return as.Call(req)
}

// AtomicFullOrderBook returns a list of open orders for a symbol.
// Level-3 order book includes all bids and asks (non-aggregated, each item in Level-3 means a single order).
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

// TradeHistories returns a list the latest trades for a symbol.
func (as *ApiService) TradeHistories(symbol string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/market/histories", map[string]string{"symbol": symbol})
	return as.Call(req)
}

type HistoricRateModel []string
type HistoricRatesModel []*HistoricRateModel

// HistoricRates returns historic rates for a symbol.
// Rates are returned in grouped buckets based on requested type.
func (as *ApiService) HistoricRates(symbol string, startAt, endAt int64, typo string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/market/candles", map[string]string{
		"symbol":  symbol,
		"startAt": IntToString(startAt),
		"endAt":   IntToString(endAt),
		"type":    typo,
	})
	return as.Call(req)
}
