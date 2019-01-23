package kucoin

import "testing"

func doPaginationTest(t *testing.T, response *ApiResponse, v interface{}) {
	p := &PaginationModel{}
	if err := response.ReadData(p); err != nil {
		t.Fatal(err)
	}
	if err := p.ReadItems(v); err != nil {
		t.Fatal(err)
	}
}
