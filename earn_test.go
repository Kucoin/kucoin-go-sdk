package kucoin

import (
	"context"
	"testing"
)

func TestApiService_CreateEarnOrder(t *testing.T) {

	s := NewApiServiceFromEnv()
	p := &CreateEarnOrderReq{
		ProductId:   "2212",
		Amount:      "10",
		AccountType: "TRADE",
	}
	rsp, err := s.CreateEarnOrder(context.Background(), p)
	if err != nil {
		t.Fatal(err)
	}
	o := &CreateEarnOrderRes{}
	if err := rsp.ReadData(o); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(o))
}

func TestApiService_DeleteEarnOrder(t *testing.T) {

	s := NewApiServiceFromEnv()
	rsp, err := s.DeleteEarnOrder(context.Background(), "2596986", "10", "TRADE", "1")
	if err != nil {
		t.Fatal(err)
	}
	o := &DeleteEarnOrderRes{}
	if err := rsp.ReadData(o); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(o))
}

func TestApiService_QueryOTCLoanInfo(t *testing.T) {

	s := NewApiServiceFromEnv()
	rsp, err := s.QueryOTCLoanInfo(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	o := &OTCLoanModel{}
	if err := rsp.ReadData(o); err != nil {
		t.Fatal(err)
	}
	for _, order := range o.Orders {
		t.Log(ToJsonString(order))
	}

	t.Log(ToJsonString(o.Ltv))
	t.Log(ToJsonString(o.ParentUid))
	t.Log(ToJsonString(o.TotalMarginAmount))
	t.Log(ToJsonString(o.TransferMarginAmount))
	for _, margin := range o.Margins {
		t.Log(ToJsonString(margin))
	}
}

func TestApiService_RedeemPreview(t *testing.T) {

	s := NewApiServiceFromEnv()
	rsp, err := s.RedeemPreview(context.Background(), "2596986", "TRADE")
	if err != nil {
		t.Fatal(err)
	}

	o := &RedeemPreviewModel{}
	if err := rsp.ReadData(&o); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(o))
}

func TestApiService_QuerySavingProducts(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.QuerySavingProducts(context.Background(), "")
	if err != nil {
		t.Fatal(err)
	}

	os := EarnProductsRes{}
	if err := rsp.ReadData(&os); err != nil {
		t.Fatal(err)
	}
	for _, o := range os {
		t.Log(ToJsonString(o))
		switch {
		case o.Id == "":
			t.Error("Empty key 'id'")
		case o.Currency == "":
			t.Error("Empty key 'Currency'")
		case o.Category == "":
			t.Error("Empty key 'Category'")
		case o.Type == "":
			t.Error("Empty key 'Type'")
		case o.Status == "":
			t.Error("Empty key 'Status'")
		}
	}
}

func TestApiService_QueryPromotionProducts(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.QueryPromotionProducts(context.Background(), "USDT")
	if err != nil {
		t.Fatal(err)
	}

	os := EarnProductsRes{}
	if err := rsp.ReadData(&os); err != nil {
		t.Fatal(err)
	}
	for _, o := range os {
		t.Log(ToJsonString(o))
		switch {
		case o.Id == "":
			t.Error("Empty key 'id'")
		case o.Currency == "":
			t.Error("Empty key 'Currency'")
		case o.Category == "":
			t.Error("Empty key 'Category'")
		case o.Type == "":
			t.Error("Empty key 'Type'")
		case o.Status == "":
			t.Error("Empty key 'Status'")
		}
	}
}

func TestApiService_QueryKCSStakingProducts(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.QueryKCSStakingProducts(context.Background(), "KCS")
	if err != nil {
		t.Fatal(err)
	}

	os := EarnProductsRes{}
	if err := rsp.ReadData(&os); err != nil {
		t.Fatal(err)
	}
	for _, o := range os {
		t.Log(ToJsonString(o))
		switch {
		case o.Id == "":
			t.Error("Empty key 'id'")
		case o.Currency == "":
			t.Error("Empty key 'Currency'")
		case o.Category == "":
			t.Error("Empty key 'Category'")
		case o.Type == "":
			t.Error("Empty key 'Type'")
		case o.Status == "":
			t.Error("Empty key 'Status'")
		}
	}
}

func TestApiService_QueryStakingProducts(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.QueryStakingProducts(context.Background(), "")
	if err != nil {
		t.Fatal(err)
	}

	os := EarnProductsRes{}
	if err := rsp.ReadData(&os); err != nil {
		t.Fatal(err)
	}
	for _, o := range os {
		t.Log(ToJsonString(o))
		switch {
		case o.Id == "":
			t.Error("Empty key 'id'")
		case o.Currency == "":
			t.Error("Empty key 'Currency'")
		case o.Category == "":
			t.Error("Empty key 'Category'")
		case o.Type == "":
			t.Error("Empty key 'Type'")
		case o.Status == "":
			t.Error("Empty key 'Status'")
		}
	}
}

func TestApiService_QueryETHProducts(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.QueryETHProducts(context.Background(), "")
	if err != nil {
		t.Fatal(err)
	}

	os := EarnProductsRes{}
	if err := rsp.ReadData(&os); err != nil {
		t.Fatal(err)
	}
	for _, o := range os {
		t.Log(ToJsonString(o))
		switch {
		case o.Id == "":
			t.Error("Empty key 'id'")
		case o.Currency == "":
			t.Error("Empty key 'Currency'")
		case o.Category == "":
			t.Error("Empty key 'Category'")
		case o.Type == "":
			t.Error("Empty key 'Type'")
		case o.Status == "":
			t.Error("Empty key 'Status'")
		}
	}
}

func TestApiService_QueryHoldAssets(t *testing.T) {
	s := NewApiServiceFromEnv()

	p := &PaginationParam{CurrentPage: 1, PageSize: 10}
	rsp, err := s.QueryHoldAssets(context.Background(), "", "", "", p)
	if err != nil {
		t.Fatal(err)
	}

	os := HoldAssetsRes{}
	if _, err := rsp.ReadPaginationData(&os); err != nil {
		t.Fatal(err)
	}
	if len(os) == 0 {
		t.SkipNow()
	}

	for _, o := range os {
		t.Log(ToJsonString(o))
	}
}

func TestApiService_QueryOTCLoanAccountsInfo(t *testing.T) {
	s := NewApiServiceFromEnv()

	rsp, err := s.QueryOTCLoanAccountsInfo(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	os := OTCAccountsModel{}
	if err := rsp.ReadData(&os); err != nil {
		t.Fatal(err)
	}
	if len(os) == 0 {
		t.SkipNow()
	}

	for _, o := range os {
		t.Log(ToJsonString(o))
	}
}
