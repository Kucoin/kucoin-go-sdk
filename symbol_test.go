package kucoin

import (
	"encoding/json"
	"testing"
)

func TestSymbolList(t *testing.T) {
	pa := NewPublicApiFromEnv()
	l, err := pa.Symbols()
	if err != nil {
		t.Error(err)
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

func TestTickerDetail(t *testing.T) {
	pa := NewPublicApiFromEnv()
	c, err := pa.Ticker("ETH-BTC")
	if err != nil {
		t.Error(err)
	}
	b, _ := json.Marshal(c)
	t.Log(string(b))
	switch {
	case c.Sequence == "":
		t.Error("Missing key 'sequence'")
	case c.Price == "":
		t.Error("Missing key 'price'")
	case c.Size == "":
		t.Error("Missing key 'size'")
	case c.BestBid == "":
		t.Error("Missing key 'bestBid'")
	case c.BestBidSize == "":
		t.Error("Missing key 'bestBidSize'")
	case c.BestAsk == "":
		t.Error("Missing key 'bestAsk'")
	case c.BestAskSize == "":
		t.Error("Missing key 'bestAskSize'")
	}
}

func TestPartOrderBook(t *testing.T) {
	pa := NewPublicApiFromEnv()
	c, err := pa.PartOrderBook("ETH-BTC")
	if err != nil {
		t.Error(err)
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

func TestAggregatedFullOrderBook(t *testing.T) {
	pa := NewPublicApiFromEnv()
	c, err := pa.AggregatedFullOrderBook("ETH-BTC")
	if err != nil {
		t.Error(err)
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

func TestAtomicFullOrderBook(t *testing.T) {
	pa := NewPublicApiFromEnv()
	c, err := pa.AtomicFullOrderBook("ETH-BTC")
	if err != nil {
		t.Error(err)
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

func TestTradeHistories(t *testing.T) {
	pa := NewPublicApiFromEnv()
	l, err := pa.TradeHistories("ETH-BTC")
	if err != nil {
		t.Error(err)
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

func TestHistoricRates(t *testing.T) {
	pa := NewPublicApiFromEnv()
	l, err := pa.HistoricRates("ETH-BTC", 0, 0, "")
	if err != nil {
		t.Error(err)
	}
	for _, c := range l {
		b, _ := json.Marshal(c)
		t.Log(string(b))
		if len(*c) != 7 {
			t.Error("Invalid length of rate")
		}
	}
}
