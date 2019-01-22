package kucoin

import "testing"

func doPaginationTest(t *testing.T, response *ApiResponse, v interface{}) {
	p := &PaginationModel{}
	if err := response.ReadData(p); err != nil {
		t.Fatal(err)
	}
	p.ReadItems(v)
}
