package kucoin

import "encoding/json"

// A PaginationParam represents the pagination parameters `currentPage` `pageSize` in a request .
type PaginationParam struct {
	CurrentPage int64
	PageSize    int64
}

// ReadParam read pagination parameters into params.
func (p *PaginationParam) ReadParam(params map[string]string) {
	params["currentPage"], params["pageSize"] = IntToString(p.CurrentPage), IntToString(p.PageSize)
}

// A PaginationModel represents the pagination in a response.
type PaginationModel struct {
	CurrentPage int64           `json:"currentPage"`
	PageSize    int64           `json:"pageSize"`
	TotalNum    int64           `json:"totalNum"`
	TotalPage   int64           `json:"totalPage"`
	RawItems    json.RawMessage `json:"items"` // delay parsing
}

// ReadItems read the `items` into v.
func (p *PaginationModel) ReadItems(v interface{}) error {
	return json.Unmarshal(p.RawItems, v)
}
