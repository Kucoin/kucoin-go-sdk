package kucoin

import (
	"testing"
	"time"
)

func TestApiService_HfAccountInnerTransfer(t *testing.T) {
	t.SkipNow()
	s := NewApiServiceFromEnv()
	clientOid := IntToString(time.Now().Unix())
	p := map[string]string{
		"clientOid": clientOid,
		"currency":  "USDT",
		"from":      "main",
		"to":        "trade_hf",
		"amount":    "0.3",
	}
	rsp, err := s.HfAccountInnerTransfer(p)
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
	rsp, err := s.HfAccounts("", "trade_hf")
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
	rsp, err := s.HfAccount("2969860516868")
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
	rsp, err := s.HfAccountTransferable("USDT")
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
	rsp, err := s.HfAccountLedgers(p)
	if err != nil {
		t.Fatal(err)
	}
	a := &HfAccountLedgersModel{}
	if err := rsp.ReadData(a); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(a))
}
