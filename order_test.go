package kucoin

import (
	"context"
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
	rsp, err := s.CreateOrder(context.Background(), p)
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
	rsp, err := s.CreateMultiOrder(context.Background(), "KCS-ETH", orders)
	if err != nil {
		t.Fatal(err)
	}
	r := &CreateMultiOrderResultModel{}
	if err := rsp.ReadData(r); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(r))
}

func TestApiService_CancelOrder(t *testing.T) {
	t.SkipNow()

	s := NewApiServiceFromEnv()
	rsp, err := s.CancelOrder(context.Background(), "order id")
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

func TestApiService_CancelOrderByClient(t *testing.T) {
	t.SkipNow()

	s := NewApiServiceFromEnv()
	rsp, err := s.CancelOrderByClient(context.Background(), "client id")
	if err != nil {
		t.Fatal(err)
	}
	o := &CancelOrderByClientResultModel{}
	if err := rsp.ReadData(o); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(o))
	switch {
	case len(o.CancelledOrderId) == 0:
		t.Error("Empty key 'cancelledOrderId'")
	case len(o.ClientOid) == 0:
		t.Error("Empty key 'clientOid'")
	}
}

func TestApiService_CancelOrders(t *testing.T) {
	t.SkipNow()

	s := NewApiServiceFromEnv()
	rsp, err := s.CancelOrders(context.Background(), map[string]string{
		"symbol":    "ETH-BTC",
		"tradeType": "TRADE",
	})
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
	rsp, err := s.Orders(context.Background(), map[string]string{}, p)
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
	rsp, err := s.V1Orders(context.Background(), map[string]string{}, p)
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
	rsp, err := s.Orders(context.Background(), map[string]string{}, p)
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

	rsp, err = s.Order(context.Background(), os[0].Id)
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
	rsp, err := s.RecentOrders(context.Background())
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

func TestApiService_OrderByClient(t *testing.T) {
	t.SkipNow()

	s := NewApiServiceFromEnv()
	rsp, err := s.OrderByClient(context.Background(), "client id")
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

func TestApiService_CreatMarginOrder(t *testing.T) {
	t.SkipNow()

	s := NewApiServiceFromEnv()
	p := &CreateOrderModel{
		ClientOid: IntToString(time.Now().UnixNano()),
		Side:      "buy",
		Symbol:    "BTC-USDT",
		Price:     "1",
		Size:      "1",
	}
	rsp, err := s.CreateMarginOrder(context.Background(), p)
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

func TestApiService_CreateStopOrder(t *testing.T) {
	t.SkipNow()

	s := NewApiServiceFromEnv()
	p := &CreateOrderModel{
		ClientOid: IntToString(time.Now().UnixNano()),
		Side:      "buy",
		Symbol:    "BTC-USDT",
		Price:     "1",
		Size:      "1",
		StopPrice: "10.0",
	}
	rsp, err := s.CreateStopOrder(context.Background(), p)
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
func TestApiService_CancelStopOrder(t *testing.T) {
	t.SkipNow()

	s := NewApiServiceFromEnv()
	rsp, err := s.CancelStopOrder(context.Background(), "xxxxx")
	if err != nil {
		t.Fatal(err)
	}
	o := &CancelOrderResultModel{}
	if err := rsp.ReadData(o); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(o))
}

func TestApiService_CancelStopOrderBy(t *testing.T) {
	t.SkipNow()

	s := NewApiServiceFromEnv()
	rsp, err := s.CancelStopOrderBy(context.Background(), map[string]string{"orderId": "xxxx"})
	if err != nil {
		t.Fatal(err)
	}
	o := &CancelOrderResultModel{}
	if err := rsp.ReadData(o); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(o))
}

func TestApiService_StopOrder(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.StopOrder(context.Background(), "vs8hoo98rathe2ak003ag5t9")
	if err != nil {
		t.Fatal(err)
	}
	o := &StopOrderModel{}
	if err := rsp.ReadData(o); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(o))
}

func TestApiService_StopOrderByClient(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.StopOrderByClient(context.Background(), "1112", map[string]string{})
	if err != nil {
		t.Fatal(err)
	}
	o := &StopOrderListModel{}
	t.Log(ToJsonString(rsp))
	if err := rsp.ReadData(o); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(o))
}

func TestApiService_CancelStopOrderByClient(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.CancelStopOrderByClient(context.Background(), "1112", map[string]string{})
	if err != nil {
		t.Fatal(err)
	}
	o := &CancelStopOrderByClientModel{}
	t.Log(ToJsonString(rsp))
	if err := rsp.ReadData(o); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(o))
}

func TestApiService_CreateOcoOrderModel(t *testing.T) {

	s := NewApiServiceFromEnv()
	p := &CreateOcoOrderModel{
		Side:       "buy",
		Symbol:     "BTC-USDT",
		Price:      "1",
		Size:       "1",
		StopPrice:  "100000",
		LimitPrice: "100002",
		TradeType:  "TRADE",
		ClientOid:  IntToString(time.Now().UnixNano()),
		Remark:     "xx",
	}
	rsp, err := s.CreateOcoOrder(context.Background(), p)
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

func TestApiService_DeleteOcoOrder(t *testing.T) {

	s := NewApiServiceFromEnv()
	rsp, err := s.DeleteOcoOrder(context.Background(), "65d1c7042e6db70007e639b2")
	if err != nil {
		t.Fatal(err)
	}
	o := &CancelledOcoOrderResModel{}
	if err := rsp.ReadData(o); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(o))
	switch {
	case len(o.CancelledOrderIds) == 0:
		t.Error("Empty key 'cancelledOrderIds'")
	}
}

func TestApiService_DeleteOcoOrderClientId(t *testing.T) {

	s := NewApiServiceFromEnv()
	rsp, err := s.DeleteOcoOrderClientId(context.Background(), "order client id")
	if err != nil {
		t.Fatal(err)
	}
	o := &CancelledOcoOrderResModel{}
	if err := rsp.ReadData(o); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(o))
	switch {
	case len(o.CancelledOrderIds) == 0:
		t.Error("Empty key 'cancelledOrderIds'")
	}
}

func TestApiService_DeleteOcoOrders(t *testing.T) {

	s := NewApiServiceFromEnv()
	rsp, err := s.DeleteOcoOrders(context.Background(), "BTC-USDT", "")
	if err != nil {
		t.Fatal(err)
	}
	o := &CancelledOcoOrderResModel{}
	if err := rsp.ReadData(o); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(o))
	switch {
	case len(o.CancelledOrderIds) == 0:
		t.Error("Empty key 'cancelledOrderIds'")
	}
}

func TestApiService_OcoOrderDetail(t *testing.T) {

	s := NewApiServiceFromEnv()
	rsp, err := s.OcoOrderDetail(context.Background(), "65d1c7042e6db70007e639b2")
	if err != nil {
		t.Fatal(err)
	}

	o := &OrderDetailModel{}
	if err := rsp.ReadData(&o); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(o))
}

func TestApiService_OcoOrder(t *testing.T) {

	s := NewApiServiceFromEnv()
	rsp, err := s.OcoOrder(context.Background(), "65d1c7042e6db70007e639b2")
	if err != nil {
		t.Fatal(err)
	}

	o := &OcoOrderResModel{}
	if err := rsp.ReadData(&o); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(o))
}

func TestApiService_OcoClientOrder(t *testing.T) {

	s := NewApiServiceFromEnv()
	rsp, err := s.OcoClientOrder(context.Background(), "1708246787246002000")
	if err != nil {
		t.Fatal(err)
	}

	o := &OcoOrderResModel{}
	if err := rsp.ReadData(&o); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(o))
}

func TestApiService_OcoOrders(t *testing.T) {

	s := NewApiServiceFromEnv()
	p2 := &PaginationParam{CurrentPage: 1, PageSize: 10}
	p1 := map[string]string{
		"symbol": "BTC-USDT",
	}
	rsp, err := s.OcoOrders(context.Background(), p1, p2)
	if err != nil {
		t.Fatal(err)
	}

	o := &OcoOrdersModel{}
	if _, err := rsp.ReadPaginationData(&o); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(o))
}
