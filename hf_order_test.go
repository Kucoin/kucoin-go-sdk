package kucoin

import (
	"context"
	"testing"
	"time"
)

func TestApiService_HfPlaceOrder(t *testing.T) {
	s := NewApiServiceFromEnv()
	clientOid := IntToString(time.Now().Unix())
	p := map[string]string{
		"clientOid": clientOid,
		"symbol":    "KCS-USDT",
		"type":      "limit",
		"side":      "buy",
		"stp":       "CN",
		"size":      "1",
		"price":     "0.1",
	}
	rsp, err := s.HfPlaceOrder(context.Background(), p)
	if err != nil {
		t.Fatal(err)
	}
	v := &HfPlaceOrderRes{}
	if err := rsp.ReadData(v); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(v))
	if v.OrderId == "" {
		t.Error("Empty key 'orderId'")
	}
}

func TestApiService_HfSyncPlaceOrder(t *testing.T) {
	t.SkipNow()
	s := NewApiServiceFromEnv()
	clientOid := IntToString(time.Now().Unix())
	p := map[string]string{
		"clientOid": clientOid,
		"symbol":    "MATIC-USDT",
		"type":      "market",
		"side":      "sell",
		"stp":       "CN",
		"tags":      "t",
		"remark":    "r",
		"size":      "0.1",
	}
	rsp, err := s.HfSyncPlaceOrder(context.Background(), p)
	if err != nil {
		t.Fatal(err)
	}
	v := &HfSyncPlaceOrderRes{}
	if err := rsp.ReadData(v); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(v))
	if v.OrderId == "" {
		t.Error("Empty key 'orderId'")
	}
}

func TestApiService_HfPlaceMultiOrders(t *testing.T) {
	t.SkipNow()
	s := NewApiServiceFromEnv()
	clientOid := IntToString(time.Now().Unix())
	p := make([]*HFCreateMultiOrderModel, 0)
	p = append(p, &HFCreateMultiOrderModel{
		ClientOid: clientOid,
		Symbol:    "MATIC-USDT",
		OrderType: "market",
		Side:      "sell",
		Size:      "0.1",
	})

	clientOid2 := IntToString(time.Now().Unix())
	p = append(p, &HFCreateMultiOrderModel{
		ClientOid: clientOid2,
		Symbol:    "MATIC-USDT",
		OrderType: "market",
		Side:      "sell",
		Size:      "0.1",
	})

	rsp, err := s.HfPlaceMultiOrders(context.Background(), p)
	if err != nil {
		t.Fatal(err)
	}
	v := &HfPlaceMultiOrdersRes{}
	if err := rsp.ReadData(v); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(v))
}

func TestApiService_HfSyncPlaceMultiOrders(t *testing.T) {
	t.SkipNow()
	s := NewApiServiceFromEnv()
	clientOid := IntToString(time.Now().Unix())
	p := make([]*HFCreateMultiOrderModel, 0)
	p = append(p, &HFCreateMultiOrderModel{
		ClientOid: clientOid,
		Symbol:    "MATIC-USDT",
		OrderType: "market",
		Side:      "buy",
		Size:      "0.1",
	})

	clientOid2 := IntToString(time.Now().Unix())
	p = append(p, &HFCreateMultiOrderModel{
		ClientOid: clientOid2,
		Symbol:    "MATIC-USDT",
		OrderType: "market",
		Side:      "buy",
		Size:      "0.2",
	})

	rsp, err := s.HfSyncPlaceMultiOrders(context.Background(), p)
	if err != nil {
		t.Fatal(err)
	}
	v := &HfSyncPlaceMultiOrdersRes{}
	if err := rsp.ReadData(v); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(v))
}

func TestApiService_HfObtainFilledOrders(t *testing.T) {
	s := NewApiServiceFromEnv()
	p := map[string]string{
		"symbol": "MATIC-USDT",
	}
	rsp, err := s.HfObtainFilledOrders(context.Background(), p)
	if err != nil {
		t.Fatal(err)
	}
	v := &HfFilledOrdersModel{}
	if err := rsp.ReadData(v); err != nil {
		t.Fatal(err)
	}
	if v == nil {
		return
	}
	for _, o := range v.Items {
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

func TestApiService_HfObtainActiveOrders(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.HfObtainActiveOrders(context.Background(), "MATIC-USDT")
	if err != nil {
		t.Fatal(err)
	}
	v := HfOrdersModel{}
	if err := rsp.ReadData(&v); err != nil {
		t.Fatal(err)
	}
	for _, o := range v {
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

func TestApiService_HfObtainActiveSymbols(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.HfObtainActiveSymbols(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	v := HfSymbolsModel{}
	if err := rsp.ReadData(&v); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(v))
}

func TestApiService_HfOrderDetail(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.HfOrderDetail(context.Background(), "649a45d576174800019e44b4", "MATIC-USDT")
	if err != nil {
		t.Fatal(err)
	}
	v := &HfOrderModel{}
	if err := rsp.ReadData(v); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(v))
}

func TestApiService_HfOrderDetailByClientOid(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.HfOrderDetailByClientOid(context.Background(), "1687832021", "MATIC-USDT")
	if err != nil {
		t.Fatal(err)
	}
	v := &HfOrderModel{}
	if err := rsp.ReadData(v); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(v))
}

// 649a45d576174800019e44b4

func TestApiService_HfModifyOrder(t *testing.T) {
	t.SkipNow()
	s := NewApiServiceFromEnv()
	p := map[string]string{
		"symbol":   "MATIC-USDT",
		"orderId":  "649a45d576174800019e44b4",
		"newPrice": "2.0",
	}
	rsp, err := s.HfModifyOrder(context.Background(), p)
	if err != nil {
		t.Fatal(err)
	}
	v := &HfModifyOrderRes{}
	if err := rsp.ReadData(v); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(v))
}

func TestApiService_HfQueryAutoCancelSetting(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.HfQueryAutoCancelSetting(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	v := &AUtoCancelSettingModel{}
	if err := rsp.ReadData(v); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(v))
}

func TestApiService_HfAutoCancelSetting(t *testing.T) {
	t.SkipNow()
	s := NewApiServiceFromEnv()
	rsp, err := s.HfAutoCancelSetting(context.Background(), 10000, "MATIC-USDT")
	if err != nil {
		t.Fatal(err)
	}
	v := &HfAutoCancelSettingRes{}
	if err := rsp.ReadData(v); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(v))
}

func TestApiService_HfCancelOrder(t *testing.T) {
	t.SkipNow()
	s := NewApiServiceFromEnv()
	rsp, err := s.HfCancelOrder(context.Background(), "649a49201a39390001adcce8", "MATIC-USDT")
	if err != nil {
		t.Fatal(err)
	}
	v := &HfOrderIdModel{}
	if err := rsp.ReadData(v); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(v))
}

func TestApiService_HfSyncCancelOrder(t *testing.T) {
	t.SkipNow()
	s := NewApiServiceFromEnv()
	rsp, err := s.HfSyncCancelOrder(context.Background(), "649a49201a39390001adcce8", "MATIC-USDT")
	if err != nil {
		t.Fatal(err)
	}
	v := &HfSyncCancelOrderRes{}
	if err := rsp.ReadData(v); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(v))
}

func TestApiService_HfCancelOrderByClientId(t *testing.T) {
	t.SkipNow()
	s := NewApiServiceFromEnv()
	rsp, err := s.HfCancelOrderByClientId(context.Background(), "649a49201a39390001adcce8", "MATIC-USDT")
	if err != nil {
		t.Fatal(err)
	}
	v := &HfClientOidModel{}
	if err := rsp.ReadData(v); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(v))
}

func TestApiService_HfSyncCancelOrderByClientId(t *testing.T) {
	t.SkipNow()
	s := NewApiServiceFromEnv()
	rsp, err := s.HfSyncCancelOrderByClientId(context.Background(), "649a49201a39390001adcce8", "MATIC-USDT")
	if err != nil {
		t.Fatal(err)
	}
	v := &HfSyncCancelOrderRes{}
	if err := rsp.ReadData(v); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(v))
}

func TestApiService_HfSyncCancelOrderWithSize(t *testing.T) {
	t.SkipNow()
	s := NewApiServiceFromEnv()
	rsp, err := s.HfSyncCancelOrderWithSize(context.Background(), "649a49201a39390001adcce8", "MATIC-USDT", "0.3")
	if err != nil {
		t.Fatal(err)
	}
	v := &HfSyncCancelOrderWithSizeRes{}
	if err := rsp.ReadData(v); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(v))
}
func TestApiService_HfSyncCancelAllOrders(t *testing.T) {
	t.SkipNow()
	s := NewApiServiceFromEnv()
	rsp, err := s.HfSyncCancelAllOrders(context.Background(), "MATIC-USDT")
	if err != nil {
		t.Fatal(err)
	}
	data := new(string)
	if err := rsp.ReadData(data); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(data))
}

func TestApiService_HfTransactionDetails(t *testing.T) {
	s := NewApiServiceFromEnv()
	p := map[string]string{
		"symbol": "MATIC-USDT",
	}
	rsp, err := s.HfTransactionDetails(context.Background(), p)
	if err != nil {
		t.Fatal(err)
	}
	v := &HfTransactionDetailsModel{}
	if err := rsp.ReadData(v); err != nil {
		t.Fatal(err)
	}
	if v == nil {
		return
	}
	for _, item := range v.Items {
		t.Log(ToJsonString(item))
	}
}
func TestApiService_HfCancelOrders(t *testing.T) {

	s := NewApiServiceFromEnv()
	rsp, err := s.HfCancelOrders(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	o := &HfCancelOrdersResultModel{}
	if err := rsp.ReadData(o); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(o))

}

func TestApiService_HfMarginActiveSymbols(t *testing.T) {

	s := NewApiServiceFromEnv()
	rsp, err := s.HfMarginActiveSymbols(context.Background(), "MARGIN_ISOLATED_TRADE")
	if err != nil {
		t.Fatal(err)
	}
	o := &HFMarginActiveSymbolsModel{}
	if err := rsp.ReadData(o); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(o))

}

func TestApiService_HfMarginOrderV3(t *testing.T) {
	s := NewApiServiceFromEnv()

	// market order
	req := &HfMarginOrderV3Req{
		ClientOid:  IntToString(time.Now().Unix()),
		Side:       "buy",
		Symbol:     "PEPE-USDT",
		Type:       "market",
		Stp:        "CN",
		IsIsolated: false,
		AutoBorrow: true,
		AutoRepay:  true,
		Funds:      "8",
	}

	rsp, err := s.HfCreateMarinOrderV3(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	o := &HfMarginOrderV3Resp{}
	if err := rsp.ReadData(o); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(o))

	reqSell := &HfMarginOrderV3Req{
		ClientOid:  IntToString(time.Now().Unix()),
		Side:       "sell",
		Symbol:     "PEPE-USDT",
		Type:       "market",
		Stp:        "CN",
		IsIsolated: false,
		AutoBorrow: true,
		AutoRepay:  true,
		Funds:      "100000",
	}

	rsp, err = s.HfCreateMarinOrderV3(context.Background(), reqSell)
	if err != nil {
		t.Fatal(err)
	}
	o = &HfMarginOrderV3Resp{}
	if err := rsp.ReadData(o); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(o))

	// limit order
	reqLimit := &HfMarginOrderV3Req{
		ClientOid:  IntToString(time.Now().Unix()),
		Side:       "buy",
		Symbol:     "SHIB-USDT",
		Type:       "limit",
		Stp:        "CN",
		IsIsolated: false,
		AutoBorrow: true,
		AutoRepay:  true,
		Price:      "0.000001",
		Size:       "1000000",
	}

	rspObj, err := s.HfCreateMarinOrderV3(context.Background(), reqLimit)
	if err != nil {
		t.Fatal(err)
	}
	o = &HfMarginOrderV3Resp{}
	if err := rspObj.ReadData(o); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(o))
}

func TestApiService_HfMarginOrderTestV3(t *testing.T) {
	s := NewApiServiceFromEnv()

	req := &HfMarginOrderV3Req{
		ClientOid:  IntToString(time.Now().Unix()),
		Side:       "buy",
		Symbol:     "PEPE-USDT",
		Type:       "market",
		Stp:        "CN",
		IsIsolated: false,
		AutoBorrow: true,
		AutoRepay:  true,
		Funds:      "8",
	}

	rsp, err := s.HfCreateMarinOrderTestV3(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	o := &HfMarginOrderV3Resp{}
	if err := rsp.ReadData(o); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(o))

}

func TestApiService_HfCancelMarinOrderV3(t *testing.T) {
	s := NewApiServiceFromEnv()

	req := &HfCancelMarinOrderV3Req{
		OrderId: "66ab62c1693a4f000753b464",
		Symbol:  "SHIB-USDT",
	}
	rsp, err := s.HfCancelMarinOrderV3(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	o := &HfCancelMarinOrderV3Resp{}
	if err := rsp.ReadData(o); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(o))
}

func TestApiService_HfCancelClientMarinOrderV3(t *testing.T) {
	s := NewApiServiceFromEnv()

	req := &HfCancelClientMarinOrderV3Req{
		ClientOid: "1722508074",
		Symbol:    "SHIB-USDT",
	}

	rsp, err := s.HfCancelClientMarinOrderV3(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	o := &HfCancelClientMarinOrderV3Resp{}
	if err := rsp.ReadData(o); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(o))

}

func TestApiService_HfCancelAllMarginOrdersV3(t *testing.T) {
	s := NewApiServiceFromEnv()

	req := &HfCancelAllMarginOrdersV3Req{
		TradeType: "MARGIN_TRADE",
		Symbol:    "SHIB-USDT",
	}

	rsp, err := s.HfCancelAllMarginOrdersV3(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	var o HfCancelAllMarginOrdersV3Resp
	if err := rsp.ReadData(&o); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(o))

}

func TestApiService_HfMarinActiveOrdersV3(t *testing.T) {
	s := NewApiServiceFromEnv()

	req := &HfMarinActiveOrdersV3Req{
		TradeType: "MARGIN_TRADE",
		Symbol:    "SHIB-USDT",
	}

	rsp, err := s.HfMarinActiveOrdersV3(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	o := &HfMarinActiveOrdersV3Resp{}
	if err := rsp.ReadData(o); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(o))

}

func TestApiService_HfMarinDoneOrdersV3(t *testing.T) {
	s := NewApiServiceFromEnv()

	{
		req := &HfMarinDoneOrdersV3Req{
			TradeType: "MARGIN_TRADE",
			Symbol:    "PEPE-USDT",
		}

		rsp, err := s.HfMarinDoneOrdersV3(context.Background(), req)
		if err != nil {
			t.Fatal(err)
		}
		o := &HfMarinDoneOrdersV3Resp{}
		if err := rsp.ReadData(o); err != nil {
			t.Fatal(err)
		}
		t.Log(ToJsonString(o))
	}
	{
		req := &HfMarinDoneOrdersV3Req{
			TradeType: "MARGIN_TRADE",
			Symbol:    "PEPE-USDT",
			Side:      "buy",
			Type:      "market",
			StartAt:   1722482940355,
		}

		rsp, err := s.HfMarinDoneOrdersV3(context.Background(), req)
		if err != nil {
			t.Fatal(err)
		}
		o := &HfMarinDoneOrdersV3Resp{}
		if err := rsp.ReadData(o); err != nil {
			t.Fatal(err)
		}
		t.Log(ToJsonString(o))
	}

}

func TestApiService_HfMarinOrderV3(t *testing.T) {
	s := NewApiServiceFromEnv()

	req := &HfMarinOrderV3Req{
		OrderId: "66ab00fc693a4f0007ac03db",
		Symbol:  "PEPE-USDT",
	}

	rsp, err := s.HfMarinOrderV3(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	o := &HfMarinOrderV3Resp{}
	if err := rsp.ReadData(o); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(o))

}

func TestApiService_HfMarinClientOrderV3(t *testing.T) {
	s := NewApiServiceFromEnv()

	req := &HfMarinClientOrderV3Req{
		ClientOid: "1722482939",
		Symbol:    "PEPE-USDT",
	}

	rsp, err := s.HfMarinClientOrderV3(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	o := &HfMarinClientOrderV3Resp{}
	if err := rsp.ReadData(o); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(o))
}

func TestApiService_HfMarinFillsV3(t *testing.T) {
	s := NewApiServiceFromEnv()

	{
		req := &HfMarinFillsV3Req{
			Symbol:    "PEPE-USDT",
			TradeType: "MARGIN_TRADE",
			Side:      "buy",
		}

		rsp, err := s.HfMarinFillsV3(context.Background(), req)
		if err != nil {
			t.Fatal(err)
		}
		o := &HfMarinFillsV3Resp{}
		if err := rsp.ReadData(o); err != nil {
			t.Fatal(err)
		}
		t.Log(ToJsonString(o))
	}

	{
		req := &HfMarinFillsV3Req{
			Symbol:    "PEPE-USDT",
			TradeType: "MARGIN_TRADE",
			Side:      "buy",
			OrderId:   "66ab00fc693a4f0007ac03db",
		}

		rsp, err := s.HfMarinFillsV3(context.Background(), req)
		if err != nil {
			t.Fatal(err)
		}
		o := &HfMarinFillsV3Resp{}
		if err := rsp.ReadData(o); err != nil {
			t.Fatal(err)
		}
		t.Log(ToJsonString(o))
	}
}
