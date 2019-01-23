package kucoin

import "encoding/json"

type PaginationModel struct {
	CurrentPage uint64          `json:"currentPage"`
	PageSize    uint64          `json:"pageSize"`
	TotalNum    uint64          `json:"totalNum"`
	TotalPage   uint64          `json:"totalPage"`
	RawItems    json.RawMessage `json:"items"` // delay parsing
}

func (p *PaginationModel) ReadItems(v interface{}) error {
	if err := json.Unmarshal(p.RawItems, v); err != nil {
		return err
	}
	return nil
}
