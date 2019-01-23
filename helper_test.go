package kucoin

import "testing"

func TestIntToString(t *testing.T) {
	var i int64 = 5200
	if IntToString(i) != "5200" {
		t.Error("Invalid string")
	}
}

func TestToJsonString(t *testing.T) {
	type test struct {
		M1 string            `json:"m1"`
		M2 int64             `json:"m2"`
		M3 bool              `json:"m3"`
		M4 map[string]string `json:"m4"`
		m5 chan string
	}
	var s = test{
		M1: "KuCoin",
		M2: 5200,
		M3: false,
		M4: map[string]string{"KCS": "$300"},
		m5: make(chan string, 10),
	}
	var a = `{"m1":"KuCoin","m2":5200,"m3":false,"m4":{"KCS":"$300"}}`
	if ToJsonString(s) != a {
		t.Error("Invalid string")
	}
}
