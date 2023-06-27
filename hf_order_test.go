package kucoin

import (
	"testing"
	"time"
)

func TestApiService_HfPlaceOrder(t *testing.T) {
	t.SkipNow()
	s := NewApiServiceFromEnv()
	clientOid := IntToString(time.Now().Unix())
	p := map[string]string{
		"clientOid": clientOid,
		"symbol":    "MATIC-USDT",
		"type":      "limit",
		"side":      "sell",
		"stp":       "CN",
		"size":      "0.1",
		"price":     "3.0",
	}
	rsp, err := s.HfPlaceOrder(p)
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
	rsp, err := s.HfSyncPlaceOrder(p)
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

	rsp, err := s.HfPlaceMultiOrders(p)
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

	rsp, err := s.HfSyncPlaceMultiOrders(p)
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
	rsp, err := s.HfObtainFilledOrders(p)
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
	rsp, err := s.HfObtainActiveOrders("MATIC-USDT")
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
	rsp, err := s.HfObtainActiveSymbols()
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
	rsp, err := s.HfOrderDetail("649a45d576174800019e44b4", "MATIC-USDT")
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
	rsp, err := s.HfOrderDetailByClientOid("1687832021", "MATIC-USDT")
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
	rsp, err := s.HfModifyOrder(p)
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
	rsp, err := s.HfQueryAutoCancelSetting()
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
	rsp, err := s.HfAutoCancelSetting(10000, "MATIC-USDT")
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
	rsp, err := s.HfCancelOrder("649a49201a39390001adcce8", "MATIC-USDT")
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
	rsp, err := s.HfSyncCancelOrder("649a49201a39390001adcce8", "MATIC-USDT")
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
	rsp, err := s.HfCancelOrderByClientId("649a49201a39390001adcce8", "MATIC-USDT")
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
	rsp, err := s.HfSyncCancelOrderByClientId("649a49201a39390001adcce8", "MATIC-USDT")
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
	rsp, err := s.HfSyncCancelOrderWithSize("649a49201a39390001adcce8", "MATIC-USDT", "0.3")
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
	rsp, err := s.HfSyncCancelAllOrders("MATIC-USDT")
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
	rsp, err := s.HfTransactionDetails(p)
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
