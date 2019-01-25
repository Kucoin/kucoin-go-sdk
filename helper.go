package kucoin

import (
	"encoding/json"
	"strconv"
)

// IntToString converts int64 to string.
func IntToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

// ToJsonString converts any value to JSON string.
func ToJsonString(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(b)
}
