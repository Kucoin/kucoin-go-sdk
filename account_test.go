package kucoin

import (
	"encoding/json"
	"testing"
)

func TestApiService_Accounts(t *testing.T) {
	s := NewApiServiceFromEnv()
	cl, err := s.Accounts("", "")
	if err != nil {
		t.Error(err)
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
	cl, err := s.Accounts("", "")
	if err != nil {
		t.Error(err)
	}
	if len(cl) == 0 {
		return
	}
	a, err := s.Account(cl[0].Id)
	if err != nil {
		t.Error(err)
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
