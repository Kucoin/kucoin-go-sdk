package kucoin

import (
	"testing"
	"time"
)

func TestApiService_CreateOrder(t *testing.T) {
	t.SkipNow()

	s := NewApiServiceFromEnv()
	p := map[string]string{
		"clientOid": IntToString(time.Now().UnixNano()),
		"side":      "buy",
		"symbol":    "KCS-ETH",
		"price":     "0.0036",
		"size":      "1",
	}
	rsp, err := s.CreateOrder(p)
	if err != nil {
		t.Fatal(err)
	}
	o := &CreateOrderResultModel{}
	if err := rsp.ReadData(o); err != nil {
		t.Fatal(err)
	}
	t.Log(JsonString(o))
	switch {
	case o.OrderId == "":
		t.Error("Empty key 'OrderId'")
	}
}

func TestApiService_CancelOrder(t *testing.T) {
	t.SkipNow()

	s := NewApiServiceFromEnv()
	rsp, err := s.CancelOrder("order id")
	if err != nil {
		t.Fatal(err)
	}
	o := &CancelOrderResultModel{}
	if err := rsp.ReadData(o); err != nil {
		t.Fatal(err)
	}
	t.Log(JsonString(o))
	switch {
	case len(o.CancelledOrderIds) == 0:
		t.Error("Empty key 'cancelledOrderIds'")
	}
}

func TestApiService_CancelOrders(t *testing.T) {
	t.SkipNow()

	s := NewApiServiceFromEnv()
	rsp, err := s.CancelOrders("BTC")
	if err != nil {
		t.Fatal(err)
	}
	o := &CancelOrderResultModel{}
	if err := rsp.ReadData(o); err != nil {
		t.Fatal(err)
	}
	t.Log(JsonString(o))
	switch {
	case len(o.CancelledOrderIds) == 0:
		t.Error("Empty key 'cancelledOrderIds'")
	}
}

func TestApiService_Orders(t *testing.T) {
	s := NewApiServiceFromEnv()
	p := &PaginationParam{CurrentPage: 1, PageSize: 10}
	rsp, err := s.Orders(map[string]string{}, p)
	if err != nil {
		t.Fatal(err)
	}

	os := OrdersModel{}
	if _, err := rsp.ReadPaginationData(&os); err != nil {
		t.Fatal(err)
	}
	for _, o := range os {
		t.Log(JsonString(o))
		switch {
		case o.Id == "":
			t.Error("Empty key 'id'")
		case o.Symbol == "":
			t.Error("Empty key 'symbol'")
		case o.OpType == "":
			t.Error("Empty key 'opType'")
		case o.Type == "":
			t.Error("Empty key 'type'")
		case o.Side == "":
			t.Error("Empty key 'side'")
		}
	}
}

func TestApiService_Order(t *testing.T) {
	s := NewApiServiceFromEnv()

	p := &PaginationParam{CurrentPage: 1, PageSize: 1}
	rsp, err := s.Orders(map[string]string{}, p)
	if err != nil {
		t.Fatal(err)
	}

	os := OrdersModel{}
	if _, err := rsp.ReadPaginationData(&os); err != nil {
		t.Fatal(err)
	}
	if len(os) == 0 {
		t.SkipNow()
	}

	rsp, err = s.Order(os[0].Id)
	if err != nil {
		t.Fatal(err)
	}

	o := &OrderModel{}
	if err := rsp.ReadData(&o); err != nil {
		t.Fatal(err)
	}
	t.Log(JsonString(o))
	switch {
	case o.Id == "":
		t.Error("Empty key 'id'")
	case o.Symbol == "":
		t.Error("Empty key 'symbol'")
	case o.OpType == "":
		t.Error("Empty key 'opType'")
	case o.Type == "":
		t.Error("Empty key 'type'")
	case o.Side == "":
		t.Error("Empty key 'side'")
	}
}
