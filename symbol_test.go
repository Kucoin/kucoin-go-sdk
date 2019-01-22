package kucoin

import (
	"encoding/json"
	"testing"
	"time"
)

func TestApiService_Symbols(t *testing.T) {
	s := NewApiServiceFromEnv()
	l := SymbolModels{}
	if _, err := s.Symbols(&l); err != nil {
		t.Fatal(err)
	}
	for _, c := range l {
		b, _ := json.Marshal(c)
		t.Log(string(b))
		switch {
		case c.Name == "":
			t.Error("Missing key 'name'")
		case c.Symbol == "":
			t.Error("Missing key 'symbol'")
		case c.BaseCurrency == "":
			t.Error("Missing key 'baseCurrency'")
		case c.QuoteCurrency == "":
			t.Error("Missing key 'quoteCurrency'")
		case c.BaseMinSize == "":
			t.Error("Missing key 'baseMinSize'")
		case c.QuoteMinSize == "":
			t.Error("Missing key 'quoteMinSize'")
		case c.BaseMaxSize == "":
			t.Error("Missing key 'baseMaxSize'")
		case c.QuoteMaxSize == "":
			t.Error("Missing key 'quoteMaxSize'")
		case c.BaseIncrement == "":
			t.Error("Missing key 'baseIncrement'")
		case c.QuoteIncrement == "":
			t.Error("Missing key 'quoteIncrement'")
		case c.PriceIncrement == "":
			t.Error("Missing key 'priceIncrement'")
		}
	}
}

func TestApiService_Ticker(t *testing.T) {
	s := NewApiServiceFromEnv()
	tk := &TickerModel{}
	if _, err := s.Ticker(tk, "ETH-BTC"); err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(tk)
	t.Log(string(b))
	switch {
	case tk.Sequence == "":
		t.Error("Missing key 'sequence'")
	case tk.Price == "":
		t.Error("Missing key 'price'")
	case tk.Size == "":
		t.Error("Missing key 'size'")
	case tk.BestBid == "":
		t.Error("Missing key 'bestBid'")
	case tk.BestBidSize == "":
		t.Error("Missing key 'bestBidSize'")
	case tk.BestAsk == "":
		t.Error("Missing key 'bestAsk'")
	case tk.BestAskSize == "":
		t.Error("Missing key 'bestAskSize'")
	}
}

func TestApiService_PartOrderBook(t *testing.T) {
	s := NewApiServiceFromEnv()
	c := &PartOrderBookModel{}
	if _, err := s.PartOrderBook(c, "ETH-BTC"); err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(c)
	t.Log(string(b))
	switch {
	case c.Symbol == "":
		t.Error("Missing key 'symbol'")
	case c.ChangeRate == "":
		t.Error("Missing key 'changeRate'")
	case c.ChangePrice == "":
		t.Error("Missing key 'changePrice'")
	case c.Open == "":
		t.Error("Missing key 'open'")
	case c.Close == "":
		t.Error("Missing key 'close'")
	case c.Low == "":
		t.Error("Missing key 'low'")
	case c.Vol == "":
		t.Error("Missing key 'vol'")
	case c.VolValue == "":
		t.Error("Missing key 'volValue'")
	}
}

func TestApiService_AggregatedFullOrderBook(t *testing.T) {
	s := NewApiServiceFromEnv()
	c := &FullOrderBookModel{}
	_, err := s.AggregatedFullOrderBook(c, "ETH-BTC")
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(c)
	t.Log(string(b))
	switch {
	case c.Sequence == "":
		t.Error("Missing key 'sequence'")
	case len(c.Bids) == 0:
		t.Error("Empty key 'bids'")
	case len(c.Asks) == 0:
		t.Error("Empty key 'asks'")
	}
}

func TestApiService_AtomicFullOrderBook(t *testing.T) {
	s := NewApiServiceFromEnv()
	c := &FullOrderBookModel{}
	if _, err := s.AtomicFullOrderBook(c, "ETH-BTC"); err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(c)
	t.Log(string(b))
	switch {
	case c.Sequence == "":
		t.Error("Missing key 'sequence'")
	case len(c.Bids) == 0:
		t.Error("Empty key 'bids'")
	case len(c.Asks) == 0:
		t.Error("Empty key 'asks'")
	}
}

func TestApiService_TradeHistories(t *testing.T) {
	s := NewApiServiceFromEnv()
	l := TradeHistoriesModel{}
	if _, err := s.TradeHistories(&l, "ETH-BTC"); err != nil {
		t.Fatal(err)
	}
	for _, c := range l {
		b, _ := json.Marshal(c)
		t.Log(string(b))
		switch {
		case c.Sequence == "":
			t.Error("Missing key 'sequence'")
		case c.Price == "":
			t.Error("Missing key 'price'")
		case c.Size == "":
			t.Error("Missing key 'size'")
		case c.Side == "":
			t.Error("Missing key 'side'")
		case c.Time == 0:
			t.Error("Missing key 'time'")
		}
	}
}

func TestApiService_HistoricRates(t *testing.T) {
	s := NewApiServiceFromEnv()
	l := HistoricRatesModel{}
	if _, err := s.HistoricRates(&l, "ETH-BTC", time.Now().Unix()-7*24*3600, time.Now().Unix(), "30min"); err != nil {
		t.Fatal(err)
	}
	for _, c := range l {
		b, _ := json.Marshal(c)
		t.Log(string(b))
		if len(*c) != 7 {
			t.Error("Invalid length of rate")
		}
	}
}
