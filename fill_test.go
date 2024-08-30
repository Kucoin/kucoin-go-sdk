package kucoin

import (
	"context"
	"testing"
)

func TestApiService_Fills(t *testing.T) {
	s := NewApiServiceFromEnv()
	p := &PaginationParam{CurrentPage: 1, PageSize: 10}
	rsp, err := s.Fills(context.Background(), map[string]string{}, p)
	if err != nil {
		t.Fatal(err)
	}

	fs := FillsModel{}
	if _, err := rsp.ReadPaginationData(&fs); err != nil {
		t.Fatal(err)
	}
	for _, f := range fs {
		t.Log(ToJsonString(f))
		switch {
		case f.Symbol == "":
			t.Error("Empty key 'symbol'")
		case f.TradeId == "":
			t.Error("Empty key 'tradeId'")
		case f.OrderId == "":
			t.Error("Empty key 'orderId'")
		case f.Type == "":
			t.Error("Empty key 'type'")
		case f.Side == "":
			t.Error("Empty key 'side'")
		}
	}
}

func TestApiService_RecentFills(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.RecentFills(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	fs := FillsModel{}
	if err := rsp.ReadData(&fs); err != nil {
		t.Fatal(err)
	}
	for _, f := range fs {
		t.Log(ToJsonString(f))
		switch {
		case f.Symbol == "":
			t.Error("Empty key 'symbol'")
		case f.TradeId == "":
			t.Error("Empty key 'tradeId'")
		case f.OrderId == "":
			t.Error("Empty key 'orderId'")
		case f.Type == "":
			t.Error("Empty key 'type'")
		case f.Side == "":
			t.Error("Empty key 'side'")
		}
	}
}
