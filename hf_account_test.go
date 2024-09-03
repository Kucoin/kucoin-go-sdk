package kucoin

import (
	"context"
	"testing"
	"time"
)

func TestApiService_HfAccountInnerTransfer(t *testing.T) {
	s := NewApiServiceFromEnv()
	clientOid := IntToString(time.Now().Unix())
	p := map[string]string{
		"clientOid": clientOid,
		"currency":  "USDT",
		"from":      "trade",
		"to":        "margin_v2",
		"amount":    "1",
	}
	rsp, err := s.HfAccountInnerTransfer(context.Background(), p)
	if err != nil {
		t.Fatal(err)
	}
	v := &InnerTransferResultModel{}
	if err := rsp.ReadData(v); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(v))
	if v.OrderId == "" {
		t.Error("Empty key 'orderId'")
	}
}

func TestApiService_HfAccounts(t *testing.T) {
	s := NewApiServiceFromEnv()

	rsp, err := s.HfAccounts(context.Background(), "", "trade_hf")
	if err != nil {
		t.Fatal(err)
	}
	v := &HfAccountsModel{}
	if err := rsp.ReadData(v); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(v))
}

func TestApiService_HfAccount(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.HfAccount(context.Background(), "2969860516868")
	if err != nil {
		t.Fatal(err)
	}
	v := &HfAccountModel{}
	if err := rsp.ReadData(v); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(v))
}

func TestApiService_HfAccountTransferable(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.HfAccountTransferable(context.Background(), "USDT")
	if err != nil {
		t.Fatal(err)
	}
	a := &HfAccountTransferableModel{}
	if err := rsp.ReadData(a); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(a))
}

func TestApiService_HfAccountLedgers(t *testing.T) {
	s := NewApiServiceFromEnv()
	p := map[string]string{}
	rsp, err := s.HfAccountLedgers(context.Background(), p)
	if err != nil {
		t.Fatal(err)
	}
	a := &HfAccountLedgersModel{}
	if err := rsp.ReadData(a); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(a))
}
