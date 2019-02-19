package kucoin

import (
	"net/http"
)

// A SymbolModel represents an available currency pairs for trading.
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

// A SymbolsModel is the set of *SymbolModel.
type SymbolsModel []*SymbolModel

// Symbols returns a list of available currency pairs for trading.
func (as *ApiService) Symbols() (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/symbols", nil)
	return as.Call(req)
}

// A TickerModel represents ticker include only the inside (i.e. best) bid and ask data, last price and last trade size.
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

// A MarketTickerModel represents a market ticker for all trading pairs in the market (including 24h volume).
type MarketTickerModel struct {
	Symbol      string `json:"symbol"`
	High        string `json:"high"`
	Vol         string `json:"vol"`
	Low         string `json:"low"`
	ChangePrice string `json:"changePrice"`
	ChangeRate  string `json:"changeRate"`
	Close       string `json:"close"`
	VolValue    string `json:"volValue"`
	Open        string `json:"open"`
}

// A MarketTickersModel is the set of *MarketTickerModel.
type MarketTickersModel []*MarketTickerModel

// MarketTickers returns all tickers as MarketTickersModel for all trading pairs in the market (including 24h volume).
func (as *ApiService) MarketTickers() (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/market/allTickers", nil)
	return as.Call(req)
}

// A Stats24hrModel represents 24 hr stats for the symbol.
// Volume is in base currency units.
// Open, high, low are in quote currency units.
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
	req := NewRequest(http.MethodGet, "/api/v1/markets", nil)
	return as.Call(req)
}

// A PartOrderBookModel represents a list of open orders for a symbol, a part of Order Book within 100 depth for each side(ask or bid).
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

// A FullOrderBookModel represents a list of open orders for a symbol, with full depth.
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

// A TradeHistoryModel represents a the latest trades for a symbol.
type TradeHistoryModel struct {
	Sequence string `json:"sequence"`
	Price    string `json:"price"`
	Size     string `json:"size"`
	Side     string `json:"side"`
	Time     int64  `json:"time"`
}

// A TradeHistoriesModel is the set of *TradeHistoryModel.
type TradeHistoriesModel []*TradeHistoryModel

// TradeHistories returns a list the latest trades for a symbol.
func (as *ApiService) TradeHistories(symbol string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/market/histories", map[string]string{"symbol": symbol})
	return as.Call(req)
}

// HistoricRateModel represents the historic rates for a symbol.
// Rates are returned in grouped buckets based on requested type.
type HistoricRateModel []string

// A HistoricRatesModel is the set of *HistoricRateModel.
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
