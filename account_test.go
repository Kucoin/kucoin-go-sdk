package kucoin

import (
	"encoding/json"
	"testing"
)

func TestAccounts(t *testing.T) {
	pa := NewPrivateApiFromEnv()
	cl, err := pa.Accounts("", "")
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
