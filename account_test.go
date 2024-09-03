package kucoin

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestApiService_Accounts(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.Accounts(context.Background(), "", "")
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
	rsp, err := s.Accounts(context.Background(), "", "")
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
	rsp, err = s.Account(context.Background(), cl[0].Id)
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
	rsp, err := s.SubAccountUsers(context.Background())
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
	rsp, err := s.SubAccounts(context.Background())
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

	ctx := context.Background()

	rsp, err := s.SubAccounts(context.Background())
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
	rsp, err = s.SubAccount(ctx, cl[0].SubUserId)
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

func TestApiService_AccountsTransferable(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.AccountsTransferable(context.Background(), "MATIC", "MAIN")
	if err != nil {
		t.Fatal(err)
	}
	a := &AccountsTransferableModel{}
	if err := rsp.ReadData(a); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(a))
}

func TestApiService_CreateAccount(t *testing.T) {
	t.SkipNow()

	s := NewApiServiceFromEnv()
	rsp, err := s.CreateAccount(context.Background(), "trade", "BTC")
	if err != nil {
		t.Log(fmt.Sprintf("Create account failed: %s, %s", rsp.Code, rsp.Message))
		t.Fatal(err)
	}
	if rsp.Code == "230005" {
		t.Log(fmt.Sprintf("Account exits: %s, %s", rsp.Code, rsp.Message))
		return
	}
	a := &CreateAccountModel{}
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

	ctx := context.Background()

	rsp, err := s.Accounts(ctx, "", "")
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
	rsp, err = s.AccountLedgersV2(ctx, map[string]string{}, p)
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

	ctx := context.Background()

	rsp, err := s.Accounts(ctx, "", "")
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
	rsp, err = s.AccountHolds(ctx, l[0].Id, p)
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

func TestApiService_InnerTransferV2(t *testing.T) {
	t.SkipNow()

	s := NewApiServiceFromEnv()
	clientOid := IntToString(time.Now().Unix())
	rsp, err := s.InnerTransferV2(context.Background(), clientOid, "KCS", "main", "trade", "2")
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

func TestApiService_SubTransferV2(t *testing.T) {
	t.SkipNow()

	s := NewApiServiceFromEnv()
	clientOid := IntToString(time.Now().Unix())
	p := map[string]string{
		"clientOid":      clientOid,
		"currency":       "MATIC",
		"amount":         "9",
		"direction":      "OUT",
		"accountType":    "MAIN",
		"subAccountType": "MAIN",
		"subUserId":      "6482f1e32ba86200010eb03e",
	}
	rsp, err := s.SubTransferV2(context.Background(), p)
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

func TestBaseFee(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.BaseFee(context.Background(), "1")
	if err != nil {
		t.Fatal(err)
	}

	v := &BaseFeeModel{}
	if err := rsp.ReadData(v); err != nil {
		t.Fatal(err)
	}

	t.Log(v)
	t.Log(ToJsonString(v))
}

func TestActualFee(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.ActualFee(context.Background(), "BTC-USDT")
	if err != nil {
		t.Fatal(err)
	}

	v := &TradeFeesResultModel{}
	if err := rsp.ReadData(v); err != nil {
		t.Fatal(err)
	}

	t.Log(v)
	t.Log(ToJsonString(v))
}

func TestApiService_SubAccountUsersV2(t *testing.T) {
	s := NewApiServiceFromEnv()
	pp := PaginationParam{
		CurrentPage: 1,
		PageSize:    2,
	}
	rsp, err := s.SubAccountUsersV2(context.Background(), &pp)
	if err != nil {
		t.Fatal(err)
	}
	cl := SubAccountUsersModelV2{}
	if _, err := rsp.ReadPaginationData(&cl); err != nil {
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

func TestApiService_UserInfoV2(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.UserSummaryInfoV2(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	cl := UserSummaryInfoModelV2{}
	if err := rsp.ReadData(&cl); err != nil {
		t.Fatal(err)
	}
	t.Log(cl)
	t.Log(ToJsonString(cl))
}

func TestApiService_CreateSubAccountV2(t *testing.T) {
	subName := "marginFen1991"
	s := NewApiServiceFromEnv()
	rsp, err := s.CreateSubAccountV2(context.Background(), "Youaremine1314.", "", subName, "Margin")
	if err != nil {
		t.Fatal(err)
	}
	cl := CreateSubAccountV2Res{}
	if err := rsp.ReadData(&cl); err != nil {
		t.Fatal(err)
	}
	t.Log(cl)
	t.Log(ToJsonString(cl))

	if cl.SubName != subName {
		t.Error("Create sub account v2 fail")
	}
}

func TestApiService_SubApiKey(t *testing.T) {
	subName := "TestSubAccount1Fen"
	s := NewApiServiceFromEnv()
	rsp, err := s.SubApiKey(context.Background(), subName, "")
	if err != nil {
		t.Fatal(err)
	}
	cl := SubApiKeyRes{}
	if err := rsp.ReadData(&cl); err != nil {
		t.Fatal(err)
	}
	t.Log(cl)
	t.Log(ToJsonString(cl))
}

func TestApiService_CreateSubApiKey(t *testing.T) {
	t.SkipNow()
	s := NewApiServiceFromEnv()
	rsp, err := s.CreateSubApiKey(context.Background(), "TestSubAccount3Fen", "123abcABC", "3", "General", "", "")
	if err != nil {
		t.Fatal(err)
	}
	cl := CreateSubApiKeyRes{}
	if err := rsp.ReadData(&cl); err != nil {
		t.Fatal(err)
	}
	t.Log(cl)
	t.Log(ToJsonString(cl))
}

func TestApiService_UpdateSubApiKey(t *testing.T) {
	t.SkipNow()
	s := NewApiServiceFromEnv()
	rsp, err := s.UpdateSubApiKey(context.Background(), "TestSubAccount1Fen", "123abcABC", "648804c835848e0001690fb9", "Trade", "", "30")
	if err != nil {
		t.Fatal(err)
	}
	cl := UpdateSubApiKeyRes{}
	if err := rsp.ReadData(&cl); err != nil {
		t.Fatal(err)
	}
	t.Log(cl)
	t.Log(ToJsonString(cl))
}

func TestApiService_DeleteSubApiKey(t *testing.T) {
	t.SkipNow()
	s := NewApiServiceFromEnv()
	rsp, err := s.DeleteSubApiKey(context.Background(), "TestSubAccount3Fen", "123abcABC", "6497fc7c19a9ea0001d7ac46")
	if err != nil {
		t.Fatal(err)
	}
	cl := UpdateSubApiKeyRes{}
	if err := rsp.ReadData(&cl); err != nil {
		t.Fatal(err)
	}
	t.Log(cl)
	t.Log(ToJsonString(cl))
}

func TestApiService_SubAccountsV2(t *testing.T) {
	s := NewApiServiceFromEnv()

	pp := PaginationParam{
		CurrentPage: 1,
		PageSize:    10,
	}
	rsp, err := s.SubAccountsV2(context.Background(), &pp)
	if err != nil {
		t.Fatal(err)
	}
	cl := SubAccountsModel{}
	if _, err := rsp.ReadPaginationData(&cl); err != nil {
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

func TestApiService_UniversalTransfer(t *testing.T) {
	p := &UniversalTransferReq{
		ClientOid:       IntToString(time.Now().Unix()),
		Type:            "INTERNAL",
		Currency:        "USDT",
		Amount:          "5",
		FromAccountType: "TRADE",
		ToAccountType:   "CONTRACT",
	}
	s := NewApiServiceFromEnv()
	rsp, err := s.UniversalTransfer(context.Background(), p)
	if err != nil {
		t.Fatal(err)
	}
	v := &UniversalTransferRes{}
	if err := rsp.ReadData(v); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(v))
	if v.OrderId == "" {
		t.Error("Empty key 'orderId'")
	}
}
