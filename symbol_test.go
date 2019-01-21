package kucoin

import (
	"encoding/json"
	"testing"
)

func TestSymbolList(t *testing.T) {
	l, err := SymbolList()
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
	c, err := TickerDetail("ETH-BTC")
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
