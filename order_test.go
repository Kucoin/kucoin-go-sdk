package kucoin

import (
	"strconv"
	"testing"
	"time"
)

func TestApiService_CreateOrder(t *testing.T) {
	t.SkipNow()

	s := NewApiServiceFromEnv()
	p := &CreateOrderModel{
		ClientOid: IntToString(time.Now().UnixNano()),
		Side:      "buy",
		Symbol:    "KCS-ETH",
		Price:     "0.0036",
		Size:      "1",
	}
	rsp, err := s.CreateOrder(p)
	if err != nil {
		t.Fatal(err)
	}
	o := &CreateOrderResultModel{}
	if err := rsp.ReadData(o); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(o))
	switch {
	case o.OrderId == "":
		t.Error("Empty key 'OrderId'")
	}
}

func TestApiService_CreateMultiOrder(t *testing.T) {
	t.SkipNow()

	s := NewApiServiceFromEnv()

	orders := make([]*CreateOrderModel, 0, 5)
	for i := 0; i < 5; i++ {
		p := &CreateOrderModel{
			ClientOid: IntToString(time.Now().UnixNano() + int64(i)),
			Side:      "buy",
			Price:     "0.0036",
			Size:      "1",
			Remark:    "Multi " + strconv.Itoa(i),
		}
		orders = append(orders, p)
	}
	rsp, err := s.CreateMultiOrder("KCS-ETH", orders)
	if err != nil {
		t.Fatal(err)
	}
	r := &CreateMultiOrderResultModel{}
	if err := rsp.ReadData(r); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(r))
	for _, o := range r.Data {
		switch {
		case o.Status == "":
			t.Error("Empty key 'status'")
		}
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
	t.Log(ToJsonString(o))
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
	t.Log(ToJsonString(o))
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
		t.Log(ToJsonString(o))
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

func TestApiService_V1Orders(t *testing.T) {
	t.SkipNow()
	s := NewApiServiceFromEnv()
	p := &PaginationParam{CurrentPage: 1, PageSize: 10}
	rsp, err := s.V1Orders(map[string]string{}, p)
	if err != nil {
		t.Fatal(err)
	}

	os := V1OrdersModel{}
	if _, err := rsp.ReadPaginationData(&os); err != nil {
		t.Fatal(err)
	}
	for _, o := range os {
		t.Log(ToJsonString(o))
		switch {
		case o.Symbol == "":
			t.Error("Empty key 'symbol'")
		case o.DealPrice == "":
			t.Error("Empty key 'dealPrice'")
		case o.DealValue == "":
			t.Error("Empty key 'dealValue'")
		case o.Amount == "":
			t.Error("Empty key 'amount'")
		case o.Fee == "":
			t.Error("Empty key 'fee'")
		case o.Side == "":
			t.Error("Empty key 'side'")
		}
	}
}

func TestApiService_Order(t *testing.T) {
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
	t.Log(ToJsonString(o))
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

func TestApiService_RecentOrders(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.RecentOrders()
	if err != nil {
		t.Fatal(err)
	}

	os := OrdersModel{}
	if err := rsp.ReadData(&os); err != nil {
		t.Fatal(err)
	}
	for _, o := range os {
		t.Log(ToJsonString(o))
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
