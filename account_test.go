package kucoin

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestApiService_Accounts(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.Accounts("", "")
	if err != nil {
		t.Fatal(err)
	}
	cl := AccountsModel{}
	if err := rsp.ReadData(&cl); err != nil {
		t.Fatal(err)
	}
	for _, c := range cl {
		t.Log(JsonSting(c))
		switch {
		case c.Id == "":
			t.Error("Empty key 'id'")
		case c.Currency == "":
			t.Error("Empty key 'currency'")
		case c.Type == "":
			t.Error("Empty key 'type'")
		case c.Balance == "":
			t.Error("Empty key 'balance'")
		case c.Available == "":
			t.Error("Empty key 'available'")
		}
	}
}

func TestApiService_Account(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.Accounts("", "")
	if err != nil {
		t.Fatal(err)
	}
	cl := AccountsModel{}
	if err := rsp.ReadData(&cl); err != nil {
		t.Fatal(err)
	}
	if len(cl) == 0 {
		return
	}
	rsp, err = s.Account(cl[0].Id)
	if err != nil {
		t.Fatal(err)
	}
	a := &AccountModel{}
	if err := rsp.ReadData(a); err != nil {
		t.Fatal(err)
	}
	t.Log(JsonSting(a))
	switch {
	case a.Currency == "":
		t.Error("Empty key 'currency'")
	case a.Type == "":
		t.Error("Empty key 'type'")
	case a.Balance == "":
		t.Error("Empty key 'balance'")
	case a.Available == "":
		t.Error("Empty key 'available'")
	}
}

func TestApiService_CreateAccount(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.CreateAccount("trade", "BTC")
	if err != nil {
		t.Log(fmt.Sprintf("Create account failed: %s, %s", rsp.Code, rsp.Message))
		t.Fatal(err)
	}
	if rsp.Code == "230005" {
		t.Log(fmt.Sprintf("Account exits: %s, %s", rsp.Code, rsp.Message))
		return
	}
	a := &AccountModel{}
	if err := rsp.ReadData(a); err != nil {
		t.Fatal(err)
	}
	t.Log(a.Id)
	switch {
	case a.Id == "":
		t.Error("Empty key 'id'")
	}
}

func TestApiService_AccountHistories(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.Accounts("", "")
	if err != nil {
		t.Fatal(err)
	}
	l := AccountsModel{}
	if err := rsp.ReadData(&l); err != nil {
		t.Fatal(err)
	}
	if len(l) == 0 {
		return
	}
	rsp, err = s.AccountHistories(l[0].Id, 0, 0)
	if err != nil {
		t.Fatal(err)
	}
	log.Println(string(rsp.RawData))
	hs := AccountHistoriesModel{}
	doPaginationTest(t, rsp, &hs)
	for _, h := range hs {
		t.Log(JsonSting(h))
		switch {
		case h.Currency == "":
			t.Error("Empty key 'currency'")
		case h.Amount == "":
			t.Error("Empty key 'amount'")
		case h.Fee == "":
			t.Error("Empty key 'fee'")
		case h.Balance == "":
			t.Error("Empty key 'balance'")
		case h.BizType == "":
			t.Error("Empty key 'bizType'")
		case h.Direction == "":
			t.Error("Empty key 'direction'")
		case h.CreatedAt == 0:
			t.Error("Empty key 'createdAt'")
		case len(h.Context) == 0:
			t.Error("Empty key 'context'")
		}
	}
}

func TestApiService_AccountHolds(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.Accounts("", "")
	if err != nil {
		t.Fatal(err)
	}
	l := AccountsModel{}
	if err := rsp.ReadData(&l); err != nil {
		t.Fatal(err)
	}
	if len(l) == 0 {
		return
	}
	rsp, err = s.AccountHolds(l[0].Id)
	if err != nil {
		t.Fatal(err)
	}
	log.Println(string(rsp.RawData))
	hs := AccountHoldsModel{}
	doPaginationTest(t, rsp, &hs)
	for _, h := range hs {
		t.Log(JsonSting(h))
		switch {
		case h.Currency == "":
			t.Error("Empty key 'currency'")
		case h.HoldAmount == "":
			t.Error("Empty key 'holdAmount'")
		case h.BizType == "":
			t.Error("Empty key 'bizType'")
		case h.OrderId == "":
			t.Error("Empty key 'orderId'")
		case h.CreatedAt == 0:
			t.Error("Empty key 'createdAt'")
		case h.UpdatedAt == 0:
			t.Error("Empty key 'updatedAt'")
		}
	}
}

func TestApiService_InnerTransfer(t *testing.T) {
	// Skip this tests
	t.SkipNow()

	s := NewApiServiceFromEnv()
	clientOid := IntToString(time.Now().Unix())
	rsp, err := s.InnerTransfer(clientOid, "xx", "yy", "0.001")
	if err != nil {
		t.Fatal(err)
	}
	v := &InterTransferResultModel{}
	if err := rsp.ReadData(v); err != nil {
		t.Fatal(err)
	}
	if v.OrderId == "" {
		t.Error("Empty key 'orderId'")
	}
}
