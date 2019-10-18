package kucoin

import (
	"testing"
)

func TestApiService_CreateDepositAddress(t *testing.T) {
	t.SkipNow()

	s := NewApiServiceFromEnv()
	rsp, err := s.CreateDepositAddress("KCS", "")
	if err != nil {
		t.Fatal(err)
	}
	a := &DepositAddressModel{}
	if err := rsp.ReadData(a); err != nil {
		t.Fatal(err)
	}

	t.Log(ToJsonString(a))

	switch {
	case a.Address == "":
		t.Error("Empty key 'address'")
	case a.Memo == "":
		t.Error("Empty key 'memo'")
	}
}

func TestApiService_DepositAddresses(t *testing.T) {
	t.SkipNow()

	s := NewApiServiceFromEnv()
	rsp, err := s.DepositAddresses("BTC", "")
	if err != nil {
		t.Fatal(err)
	}
	if rsp.Code == "260200" {
		// Ignore deposit.disabled
		return
	}
	as := DepositAddressesModel{}
	if err := rsp.ReadData(&as); err != nil {
		t.Fatal(err)
	}

	for _, a := range as {
		t.Log(ToJsonString(a))
		switch {
		case a.Address == "":
			t.Error("Empty key 'address'")
		case a.Memo == "":
			t.Error("Empty key 'memo'")
		}
	}
}

func TestApiService_Deposits(t *testing.T) {
	s := NewApiServiceFromEnv()
	p := map[string]string{}
	pp := &PaginationParam{CurrentPage: 1, PageSize: 10}
	rsp, err := s.Deposits(p, pp)
	if err != nil {
		t.Fatal(err)
	}
	ds := DepositsModel{}
	if _, err := rsp.ReadPaginationData(&ds); err != nil {
		t.Fatal(err)
	}

	for _, d := range ds {
		t.Log(ToJsonString(d))
		switch {
		case d.Address == "":
			t.Error("Empty key 'address'")
		case d.Amount == "":
			t.Error("Empty key 'amount'")
		case d.Fee == "":
			t.Error("Empty key 'fee'")
		case d.Currency == "":
			t.Error("Empty key 'currency'")
		case d.WalletTxId == "":
			t.Error("Empty key 'walletTxId'")
		case d.Status == "":
			t.Error("Empty key 'status'")
		case d.CreatedAt == 0:
			t.Error("Empty key 'createdAt'")
		case d.UpdatedAt == 0:
			t.Error("Empty key 'updatedAt'")
		}
	}
}

func TestApiService_V1Deposits(t *testing.T) {
	t.SkipNow()
	s := NewApiServiceFromEnv()
	p := map[string]string{}
	pp := &PaginationParam{CurrentPage: 1, PageSize: 10}
	rsp, err := s.V1Deposits(p, pp)
	if err != nil {
		t.Fatal(err)
	}
	ds := V1DepositsModel{}
	if _, err := rsp.ReadPaginationData(&ds); err != nil {
		t.Fatal(err)
	}

	for _, d := range ds {
		t.Log(ToJsonString(d))
		switch {
		case d.Amount == "":
			t.Error("Empty key 'amount'")
		case d.Currency == "":
			t.Error("Empty key 'currency'")
		case d.WalletTxId == "":
			t.Error("Empty key 'walletTxId'")
		case d.Status == "":
			t.Error("Empty key 'status'")
		case d.CreateAt == 0:
			t.Error("Empty key 'createAt'")
		}
	}
}
