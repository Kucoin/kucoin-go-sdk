package kucoin

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestApiService_Accounts(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.Accounts("", "")
	if err != nil {
		t.Fatal(err)
	}
	cl := AccountsModel{}
	rsp.ReadData(&cl)
	for _, c := range cl {
		b, _ := json.Marshal(c)
		t.Log(string(b))
		switch {
		case c.Id == "":
			t.Error("Missing key 'id'")
		case c.Currency == "":
			t.Error("Missing key 'currency'")
		case c.Type == "":
			t.Error("Missing key 'type'")
		case c.Balance == "":
			t.Error("Missing key 'balance'")
		case c.Available == "":
			t.Error("Missing key 'available'")
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
	rsp.ReadData(cl)
	if len(cl) == 0 {
		return
	}
	rsp, err = s.Account(cl[0].Id)
	if err != nil {
		t.Fatal(err)
	}
	a := &AccountModel{}
	rsp.ReadData(a)
	b, _ := json.Marshal(a)
	t.Log(string(b))
	switch {
	case a.Currency == "":
		t.Error("Missing key 'currency'")
	case a.Type == "":
		t.Error("Missing key 'type'")
	case a.Balance == "":
		t.Error("Missing key 'balance'")
	case a.Available == "":
		t.Error("Missing key 'available'")
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
	rsp.ReadData(a)
	t.Log(a.Id)
	switch {
	case a.Id == "":
		t.Error("Missing key 'id'")
	}
}
