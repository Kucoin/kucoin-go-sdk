package kucoin

import (
	"fmt"
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
		t.Log(ToJsonString(c))
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
	t.Log(ToJsonString(a))
	switch {
	case a.Currency == "":
		t.Error("Empty key 'currency'")
	case a.Holds == "":
		t.Error("Empty key 'holds'")
	case a.Balance == "":
		t.Error("Empty key 'balance'")
	case a.Available == "":
		t.Error("Empty key 'available'")
	}
}

func TestApiService_SubAccountUsers(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.SubAccountUsers()
	if err != nil {
		t.Fatal(err)
	}
	cl := SubAccountUsersModel{}
	if err := rsp.ReadData(&cl); err != nil {
		t.Fatal(err)
	}
	if len(cl) == 0 {
		return
	}
	for _, c := range cl {
		t.Log(ToJsonString(c))
		switch {
		case c.UserId == "":
			t.Error("Empty key 'userId'")
		case c.SubName == "":
			t.Error("Empty key 'subName'")
		}
	}
}

func TestApiService_SubAccounts(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.SubAccounts()
	if err != nil {
		t.Fatal(err)
	}
	cl := SubAccountsModel{}
	if err := rsp.ReadData(&cl); err != nil {
		t.Fatal(err)
	}
	if len(cl) == 0 {
		return
	}
	for _, c := range cl {
		t.Log(ToJsonString(c))
		switch {
		case c.SubUserId == "":
			t.Error("Empty key 'subUserId'")
		case c.SubName == "":
			t.Error("Empty key 'subName'")
		}
		for _, b := range c.MainAccounts {
			switch {
			case b.Currency == "":
				t.Error("Empty key 'currency'")
			}
		}
		for _, b := range c.TradeAccounts {
			switch {
			case b.Currency == "":
				t.Error("Empty key 'currency'")
			}
		}
	}
}

func TestApiService_SubAccount(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.SubAccounts()
	if err != nil {
		t.Fatal(err)
	}
	cl := SubAccountsModel{}
	if err := rsp.ReadData(&cl); err != nil {
		t.Fatal(err)
	}
	if len(cl) == 0 {
		return
	}
	rsp, err = s.SubAccount(cl[0].SubUserId)
	if err != nil {
		t.Fatal(err)
	}
	a := SubAccountModel{}
	if err := rsp.ReadData(&a); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(a))
	switch {
	case a.SubUserId == "":
		t.Error("Empty key 'subUserId'")
	case a.SubName == "":
		t.Error("Empty key 'subName'")
	}
	for _, b := range a.MainAccounts {
		switch {
		case b.Currency == "":
			t.Error("Empty key 'currency'")
		}
	}
	for _, b := range a.TradeAccounts {
		switch {
		case b.Currency == "":
			t.Error("Empty key 'currency'")
		}
	}
}

func TestApiService_CreateAccount(t *testing.T) {
	t.SkipNow()

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

func TestApiService_AccountLedgers(t *testing.T) {
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
	p := &PaginationParam{CurrentPage: 1, PageSize: 10}
	rsp, err = s.AccountLedgers(l[0].Id, 0, 0, map[string]string{}, p)
	if err != nil {
		t.Fatal(err)
	}
	hs := AccountLedgersModel{}
	if _, err := rsp.ReadPaginationData(&hs); err != nil {
		t.Fatal(err)
	}
	for _, h := range hs {
		t.Log(ToJsonString(h))
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
	p := &PaginationParam{CurrentPage: 1, PageSize: 10}
	rsp, err = s.AccountHolds(l[0].Id, p)
	if err != nil {
		t.Fatal(err)
	}
	hs := AccountHoldsModel{}
	if _, err := rsp.ReadPaginationData(&hs); err != nil {
		t.Fatal(err)
	}
	for _, h := range hs {
		t.Log(ToJsonString(h))
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
	t.SkipNow()

	s := NewApiServiceFromEnv()
	clientOid := IntToString(time.Now().Unix())
	rsp, err := s.InnerTransfer(clientOid, "xx", "yy", "0.001")
	if err != nil {
		t.Fatal(err)
	}
	v := &InnerTransferResultModel{}
	if err := rsp.ReadData(v); err != nil {
		t.Fatal(err)
	}
	if v.OrderId == "" {
		t.Error("Empty key 'orderId'")
	}
}

func TestApiService_InnerTransferV2(t *testing.T) {
	t.SkipNow()

	s := NewApiServiceFromEnv()
	clientOid := IntToString(time.Now().Unix())
	rsp, err := s.InnerTransferV2(clientOid, "KCS", "main", "trade", "2")
	if err != nil {
		t.Fatal(err)
	}
	v := &InnerTransferResultModel{}
	if err := rsp.ReadData(v); err != nil {
		t.Fatal(err)
	}
	if v.OrderId == "" {
		t.Error("Empty key 'orderId'")
	}
}

func TestApiService_SubTransfer(t *testing.T) {
	t.SkipNow()

	s := NewApiServiceFromEnv()
	clientOid := IntToString(time.Now().Unix())
	p := map[string]string{
		"clientOid":      clientOid,
		"currency":       "KCS",
		"amount":         "1.0",
		"direction":      "IN",
		"accountType":    "main",
		"subAccountType": "trade",
		"subUserId":      "5cc5b31c38300c336230d071",
	}
	rsp, err := s.SubTransfer(p)
	if err != nil {
		t.Fatal(err)
	}
	v := &SubTransferResultModel{}
	if err := rsp.ReadData(v); err != nil {
		t.Fatal(err)
	}
	if v.OrderId == "" {
		t.Error("Empty key 'orderId'")
	}
}
