package kucoin

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// A MarkPriceModel represents mark price of a symbol
type MarkPriceModel struct {
	Symbol      string      `json:"symbol"`
	Granularity json.Number `json:"granularity"`
	TimePoint   json.Number `json:"timePoint"`
	Value       json.Number `json:"value"`
}

// CurrentMarkPrice returns current mark price of the input symbol
func (as *ApiService) CurrentMarkPrice(symbol string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/mark-price/%s/current", symbol), nil)
	return as.Call(req)
}

// A MarginConfigModel represents a margin configuration
type MarginConfigModel struct {
	CurrencyList     []string    `json:"currencyList"`
	WarningDebtRatio json.Number `json:"warningDebtRatio"`
	LiqDebtRatio     json.Number `json:"liqDebtRatio"`
	MaxLeverage      json.Number `json:"maxLeverage"`
}

// MarginConfig returns a margin configuration
func (as *ApiService) MarginConfig() (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/margin/config", nil)
	return as.Call(req)
}

// A MarginAccountModel represents a margin account information
type MarginAccountModel struct {
	Accounts []struct {
		AvailableBalance json.Number `json:"availableBalance"`
		Currency         string      `json:"currency"`
		HoldBalance      json.Number `json:"holdBalance"`
		Liability        json.Number `json:"liability"`
		MaxBorrowSize    json.Number `json:"maxBorrowSize"`
		TotalBalance     json.Number `json:"totalBalance"`
	} `json:"accounts"`
	DebtRatio json.Number `json:"debtRatio"`
}

// MarginAccount returns a margin account information
func (as *ApiService) MarginAccount() (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/margin/account", nil)
	return as.Call(req)
}
