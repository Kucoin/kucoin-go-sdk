package kucoin

import (
	"encoding/json"
	"testing"
	"time"
)

func TestApiService_Symbols(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.Symbols()
	if err != nil {
		t.Fatal(err)
	}
	l := SymbolsModel{}
	rsp.ReadData(&l)
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
	rsp, err := s.Ticker("ETH-BTC")
	if err != nil {
		t.Fatal(err)
	}
	tk := &TickerModel{}
	rsp.ReadData(tk)
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
	rsp, err := s.PartOrderBook("ETH-BTC")
	if err != nil {
		t.Fatal(err)
	}
	c := &PartOrderBookModel{}
	rsp.ReadData(c)
	b, _ := json.Marshal(c)
	t.Log(string(b))
	switch {
	case c.Sequence == "":
		t.Error("Missing key 'sequence'")
	case len(c.Asks) == 0:
		t.Error("Empty key 'asks'")
	case len(c.Bids) == 0:
		t.Error("Empty key 'bids'")
	}
}

func TestApiService_AggregatedFullOrderBook(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.AggregatedFullOrderBook("ETH-BTC")
	if err != nil {
		t.Fatal(err)
	}
	c := &FullOrderBookModel{}
	rsp.ReadData(c)
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
	rsp, err := s.AtomicFullOrderBook("ETH-BTC")
	if err != nil {
		t.Fatal(err)
	}
	c := &FullOrderBookModel{}
	rsp.ReadData(c)
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
	rsp, err := s.TradeHistories("ETH-BTC")
	if err != nil {
		t.Fatal(err)
	}
	l := TradeHistoriesModel{}
	rsp.ReadData(&l)
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
	rsp, err := s.HistoricRates("ETH-BTC", time.Now().Unix()-7*24*3600, time.Now().Unix(), "30min")
	if err != nil {
		t.Fatal(err)
	}
	l := HistoricRatesModel{}
	rsp.ReadData(&l)
	for _, c := range l {
		b, _ := json.Marshal(c)
		t.Log(string(b))
		if len(*c) != 7 {
			t.Error("Invalid length of rate")
		}
	}
}
