package kucoin

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestApiService_Accounts(t *testing.T) {
	s := NewApiServiceFromEnv()
	cl := AccountsModel{}
	if _, err := s.Accounts(&cl, "", ""); err != nil {
		t.Fatal(err)
	}
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
	cl := AccountsModel{}
	if _, err := s.Accounts(&cl, "", ""); err != nil {
		t.Fatal(err)
	}
	if len(cl) == 0 {
		return
	}
	a := &AccountModel{}
	if _, err := s.Account(a, cl[0].Id); err != nil {
		t.Fatal(err)
	}
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
	a := &AccountModel{}
	if rsp, err := s.CreateAccount(a, "trade", "BTC"); err != nil {
		t.Log(fmt.Sprintf("Create account failed: %s, %s", rsp.Code, rsp.Message))
		t.Fatal(err)
	}
	t.Log(a.Id)
	switch {
	case a.Id == "":
		t.Error("Missing key 'id'")
	}
}
