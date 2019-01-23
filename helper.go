package kucoin

import (
	"encoding/json"
	"strconv"
)

func IntToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

func ToJsonString(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(b)
}
