package kucoin

import (
	"context"
	"testing"
	"time"
)

func TestApiService_Symbols(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.Symbols(context.Background(), "")
	if err != nil {
		t.Fatal(err)
	}
	l := SymbolsModel{}
	if err := rsp.ReadData(&l); err != nil {
		t.Fatal(err)
	}
	for _, c := range l {
		t.Log(ToJsonString(c))
		switch {
		case c.Name == "":
			t.Error("Empty key 'name'")
		case c.Symbol == "":
			t.Error("Empty key 'symbol'")
		case c.BaseCurrency == "":
			t.Error("Empty key 'baseCurrency'")
		case c.QuoteCurrency == "":
			t.Error("Empty key 'quoteCurrency'")
		case c.BaseMinSize == "":
			t.Error("Empty key 'baseMinSize'")
		case c.QuoteMinSize == "":
			t.Error("Empty key 'quoteMinSize'")
		case c.BaseMaxSize == "":
			t.Error("Empty key 'baseMaxSize'")
		case c.QuoteMaxSize == "":
			t.Error("Empty key 'quoteMaxSize'")
		case c.BaseIncrement == "":
			t.Error("Empty key 'baseIncrement'")
		case c.QuoteIncrement == "":
			t.Error("Empty key 'quoteIncrement'")
		case c.FeeCurrency == "":
			t.Error("Empty key 'feeCurrency'")
		case c.PriceIncrement == "":
			t.Error("Empty key 'priceIncrement'")
		}
	}
}

func TestApiService_TickerLevel1(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.TickerLevel1(context.Background(), "ETH-BTC")
	if err != nil {
		t.Fatal(err)
	}
	tk := &TickerLevel1Model{}
	if err := rsp.ReadData(tk); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(tk))
	switch {
	case tk.Sequence == "":
		t.Error("Empty key 'sequence'")
	case tk.Price == "":
		t.Error("Empty key 'price'")
	case tk.Size == "":
		t.Error("Empty key 'size'")
	case tk.BestBid == "":
		t.Error("Empty key 'bestBid'")
	case tk.BestBidSize == "":
		t.Error("Empty key 'bestBidSize'")
	case tk.BestAsk == "":
		t.Error("Empty key 'bestAsk'")
	case tk.BestAskSize == "":
		t.Error("Empty key 'bestAskSize'")
	}
}

func TestApiService_Tickers(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.Tickers(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	ts := &TickersResponseModel{}
	if err := rsp.ReadData(ts); err != nil {
		t.Fatal(err)
	}
	if ts.Time == 0 {
		t.Error("Empty key 'time'")
	}
	for _, tk := range ts.Tickers {
		switch {
		case tk.Symbol == "":
			t.Error("Empty key 'symbol'")
		case tk.Vol == "":
			t.Error("Empty key 'vol'")
		case tk.ChangeRate == "":
			t.Error("Empty key 'changeRate'")
			//case tk.Buy == "":
			//	t.Error("Empty key 'buy'")
			//case tk.Sell == "":
			//	t.Error("Empty key 'sell'")
			//case tk.Last == "":
			//	t.Error("Empty key 'last'")
		}
	}
}

func TestApiService_Stats24hr(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.Stats24hr(context.Background(), "ETH-BTC")
	if err != nil {
		t.Fatal(err)
	}
	st := &Stats24hrModel{}
	if err := rsp.ReadData(st); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(st))
	switch {
	case st.Symbol == "":
		t.Error("Empty key 'symbol'")
	case st.ChangeRate == "":
		t.Error("Empty key 'changRate'")
	}
}

func TestApiService_Markets(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.Markets(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	ms := MarketsModel{}

	if err := rsp.ReadData(&ms); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(ms))
	if len(ms) == 0 {
		t.Error("Empty markets")
	}
}

func TestApiService_AggregatedPartOrderBook(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.AggregatedPartOrderBook(context.Background(), "ETH-BTC", 100)
	if err != nil {
		t.Fatal(err)
	}
	c := &PartOrderBookModel{}
	if err := rsp.ReadData(c); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(c))
	switch {
	case c.Sequence == "":
		t.Error("Empty key 'sequence'")
	case len(c.Asks) == 0:
		t.Error("Empty key 'asks'")
	case len(c.Asks[0]) != 2:
		t.Error("Invalid ask length")
	case len(c.Bids) == 0:
		t.Error("Empty key 'bids'")
	case len(c.Bids[0]) != 2:
		t.Error("Invalid bid length")
	}
}

func TestApiService_AggregatedFullOrderBook(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.AggregatedFullOrderBook(context.Background(), "ETH-BTC")
	if err != nil {
		t.Fatal(err)
	}
	c := &FullOrderBookModel{}
	if err := rsp.ReadData(c); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(c))
	switch {
	case c.Sequence == "":
		t.Error("Empty key 'sequence'")
	case len(c.Asks) == 0:
		t.Error("Empty key 'asks'")
	case len(c.Asks[0]) != 2:
		t.Error("Invalid ask length")
	case len(c.Bids) == 0:
		t.Error("Empty key 'bids'")
	case len(c.Bids[0]) != 2:
		t.Error("Invalid bid length")
	}
}
func TestApiService_AggregatedFullOrderBookV3(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.AggregatedFullOrderBookV3(context.Background(), "BTC-USDT")
	if err != nil {
		t.Fatal(err)
	}
	c := &FullOrderBookModel{}
	if err := rsp.ReadData(c); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(c))
	switch {
	case c.Sequence == "":
		t.Error("Empty key 'sequence'")
	case len(c.Asks) == 0:
		t.Error("Empty key 'asks'")
	case len(c.Asks[0]) != 2:
		t.Error("Invalid ask length")
	case len(c.Bids) == 0:
		t.Error("Empty key 'bids'")
	case len(c.Bids[0]) != 2:
		t.Error("Invalid bid length")
	}
}

func TestApiService_AtomicFullOrderBook(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.AtomicFullOrderBook(context.Background(), "ETH-BTC")
	if err != nil {
		t.Fatal(err)
	}
	c := &FullOrderBookModel{}
	if err := rsp.ReadData(c); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(c))
	switch {
	case c.Sequence == "":
		t.Error("Empty key 'sequence'")
	case len(c.Asks) == 0:
		t.Error("Empty key 'asks'")
	case len(c.Asks[0]) != 4:
		t.Error("Invalid ask length")
	case len(c.Bids) == 0:
		t.Error("Empty key 'bids'")
	case len(c.Bids[0]) != 4:
		t.Error("Invalid bid length")
	}
}

func TestApiService_AtomicFullOrderBookV2(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.AtomicFullOrderBookV2(context.Background(), "ETH-BTC")
	if err != nil {
		t.Fatal(err)
	}
	c := &FullOrderBookV2Model{}
	if err := rsp.ReadData(c); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(c))
	switch {
	case c.Sequence == 0:
		t.Error("Empty key 'sequence'")
	case len(c.Asks) == 0:
		t.Error("Empty key 'asks'")
	case len(c.Asks[0]) != 4:
		t.Error("Invalid ask length")
	case len(c.Bids) == 0:
		t.Error("Empty key 'bids'")
	case len(c.Bids[0]) != 4:
		t.Error("Invalid bid length")
	}
}

func TestApiService_TradeHistories(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.TradeHistories(context.Background(), "ETH-BTC")
	if err != nil {
		t.Fatal(err)
	}
	l := TradeHistoriesModel{}
	if err := rsp.ReadData(&l); err != nil {
		t.Fatal(err)
	}
	for _, c := range l {
		t.Log(ToJsonString(c))
		switch {
		case c.Sequence == "":
			t.Error("Empty key 'sequence'")
		case c.Price == "":
			t.Error("Empty key 'price'")
		case c.Size == "":
			t.Error("Empty key 'size'")
		case c.Side == "":
			t.Error("Empty key 'side'")
		case c.Time == 0:
			t.Error("Empty key 'time'")
		}
	}
}

func TestApiService_KLines(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.KLines(context.Background(), "ETH-BTC", "30min", time.Now().Unix()-7*24*3600, time.Now().Unix())
	if err != nil {
		t.Fatal(err)
	}
	l := KLinesModel{}
	if err := rsp.ReadData(&l); err != nil {
		t.Fatal(err)
	}
	for _, c := range l {
		t.Log(ToJsonString(c))
		if len(*c) != 7 {
			t.Error("Invalid length of rate")
		}
	}
}

func TestApiService_SymbolsV2(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.SymbolsV2(context.Background(), "ETF")
	if err != nil {
		t.Fatal(err)
	}
	l := SymbolsModelV2{}
	if err := rsp.ReadData(&l); err != nil {
		t.Fatal(err)
	}
	for _, c := range l {
		t.Log(ToJsonString(c))
		switch {
		case c.Name == "":
			t.Error("Empty key 'name'")
		case c.Symbol == "":
			t.Error("Empty key 'symbol'")
		case c.BaseCurrency == "":
			t.Error("Empty key 'baseCurrency'")
		case c.QuoteCurrency == "":
			t.Error("Empty key 'quoteCurrency'")
		case c.BaseMinSize == "":
			t.Error("Empty key 'baseMinSize'")
		case c.QuoteMinSize == "":
			t.Error("Empty key 'quoteMinSize'")
		case c.BaseMaxSize == "":
			t.Error("Empty key 'baseMaxSize'")
		case c.QuoteMaxSize == "":
			t.Error("Empty key 'quoteMaxSize'")
		case c.BaseIncrement == "":
			t.Error("Empty key 'baseIncrement'")
		case c.QuoteIncrement == "":
			t.Error("Empty key 'quoteIncrement'")
		case c.FeeCurrency == "":
			t.Error("Empty key 'feeCurrency'")
		case c.PriceIncrement == "":
			t.Error("Empty key 'priceIncrement'")
		case c.MinFunds == "":
			t.Error("Empty key 'feeCurrency'")
		}

	}
}

func TestApiService_SymbolsV2Detail(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.SymbolsDetail(context.Background(), "BTC-USDT")
	if err != nil {
		t.Fatal(err)
	}
	c := SymbolModelV2{}
	if err := rsp.ReadData(&c); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(c))
}
