package kucoin

import (
	"context"
	"math/rand"
	"testing"
	"time"
)

func TestApiService_CurrentMarkPrice(t *testing.T) {
	symbols := []string{"USDT-BTC", "ETH-BTC", "LTC-BTC", "EOS-BTC", "XRP-BTC", "KCS-BTC"}
	rand.Seed(time.Now().UnixNano())
	symbol := symbols[rand.Intn(len(symbols))]

	s := NewApiServiceFromEnv()
	rsp, err := s.CurrentMarkPrice(context.Background(), symbol)

	if err != nil {
		t.Fatal(err)
	}
	o := &MarkPriceModel{}
	if err := rsp.ReadData(o); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(o))
	switch {
	case o.Symbol == "":
		t.Error("empty key 'Symbol'")
	case o.Granularity == "":
		t.Error("empty key 'granularity'")
	case o.TimePoint == "":
		t.Error("empty key 'TimePoint'")
	case o.Value == "":
		t.Error("empty key 'Value'")
	}
}

func TestApiService_MarginConfig(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.MarginConfig(context.Background())

	if err != nil {
		t.Fatal(err)
	}

	o := &MarginConfigModel{}
	if err := rsp.ReadData(o); err != nil {
		t.Fatal(err)
	}

	t.Log(ToJsonString(o))
	switch {
	case len(o.CurrencyList) == 0:
		t.Error("zero length of 'CurrencyList'")
	case o.WarningDebtRatio == "":
		t.Error("empty key 'WarningDebtRatio'")
	case o.LiqDebtRatio == "":
		t.Error("empty key 'LiqDebtRatio'")
	case o.MaxLeverage == "":
		t.Error("empty key 'MaxLeverage'")
	}
}

func TestApiService_MarginAccount(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.MarginAccount(context.Background())

	if err != nil {
		t.Fatal(err)
	}

	o := &MarginAccountModel{}
	if err := rsp.ReadData(o); err != nil {
		t.Fatal(err)
	}

	t.Log(ToJsonString(o))

	switch {
	case len(o.Accounts) == 0:
		t.Error("zero length of 'Accounts'")
	case o.DebtRatio == "":
		t.Error("empty key 'DebtRatio'")
	}
}

func TestApiService_CreateBorrowOrder(t *testing.T) {
	t.SkipNow()
	s := NewApiServiceFromEnv()

	params := map[string]string{"currency": "BTC", "type": "IOC", "size": "0.003"}
	rsp, err := s.CreateBorrowOrder(context.Background(), params)

	if err != nil {
		t.Fatal(err)
	}

	o := &CreateBorrowOrderResultModel{}
	if err := rsp.ReadData(o); err != nil {
		t.Fatal(err)
	}

	t.Log(ToJsonString(o))

	switch {
	case o.OrderId == "":
		t.Error("empty key 'OrderId'")
	}
}

func TestApiService_BorrowOrder(t *testing.T) {
	t.SkipNow()
	s := NewApiServiceFromEnv()
	rsp, err := s.BorrowOrder(context.Background(), "5e12f29dd43f8d0008c87981")

	if err != nil {
		t.Fatal(err)
	}

	o := &BorrowOrderModel{}
	if err := rsp.ReadData(o); err != nil {
		t.Fatal(err)
	}

	t.Log(ToJsonString(o))

	switch {
	case o.OrderId == "":
		t.Error("empty key 'OrderId'")
	case o.Currency == "":
		t.Error("empty key 'Currency'")
	case o.Size == "":
		t.Error("empty key 'Size'")
	case o.Filled == "":
		t.Error("empty key 'Filled'")
	case o.Status == "":
		t.Error("empty key 'Status'")
	case len(o.MatchList) == 0:
		t.Error("zero length of 'Status'")
	}
}

func TestApiService_BorrowOutstandingRecords(t *testing.T) {
	s := NewApiServiceFromEnv()
	pagination := &PaginationParam{
		CurrentPage: 1,
		PageSize:    10,
	}

	rsp, err := s.BorrowOutstandingRecords(context.Background(), "", pagination)
	if err != nil {
		t.Fatal(err)
	}

	os := BorrowOutstandingRecordsModel{}
	if _, err := rsp.ReadPaginationData(&os); err != nil {
		t.Fatal(err)
	}

	t.Log(ToJsonString(os))

	if len(os) == 0 {
		t.SkipNow()
	}

	for _, o := range os {
		switch {
		case o.Currency == "":
			t.Error("empty key 'Currency'")
		case o.TradeId == "":
			t.Error("empty key 'TradeId'")
		case o.Liability == "":
			t.Error("empty key 'Liability'")
		case o.Principal == "":
			t.Error("empty key 'Principal'")
		case o.AccruedInterest == "":
			t.Error("empty key 'AccruedInterest'")
		case o.CreatedAt == "":
			t.Error("empty key 'createdAt'")
		case o.MaturityTime == "":
			t.Error("empty key 'MaturityTime'")
		case o.Term == "":
			t.Error("empty key 'Term'")
		case o.RepaidSize == "":
			t.Error("empty key 'RepaidSize'")
		case o.DailyIntRate == "":
			t.Error("empty key 'DailyIntRate'")
		}
	}
}

func TestApiService_BorrowRepaidRecords(t *testing.T) {
	s := NewApiServiceFromEnv()
	pagination := &PaginationParam{
		CurrentPage: 1,
		PageSize:    10,
	}

	rsp, err := s.BorrowRepaidRecords(context.Background(), "", pagination)
	if err != nil {
		t.Fatal(err)
	}

	os := BorrowRepaidRecordsModel{}
	if _, err := rsp.ReadPaginationData(&os); err != nil {
		t.Fatal(err)
	}

	t.Log(ToJsonString(os))

	if len(os) == 0 {
		t.SkipNow()
	}

	for _, o := range os {
		switch {
		case o.Currency == "":
			t.Error("empty key 'Currency'")
		case o.DailyIntRate == "":
			t.Error("empty key 'DailyIntRate'")
		case o.Interest == "":
			t.Error("empty key 'Interest'")
		case o.Principal == "":
			t.Error("empty key 'Principal'")
		case o.RepayTime == "":
			t.Error("empty key 'RepayTime'")
		case o.Term == "":
			t.Error("empty key 'Term'")
		case o.TradeId == "":
			t.Error("empty key 'TradeId'")
		}
	}
}

func TestApiService_RepayAll(t *testing.T) {
	t.SkipNow()
	s := NewApiServiceFromEnv()
	params := map[string]string{"currency": "BTC", "sequence": "RECENTLY_EXPIRE_FIRST", "size": "0.00001"}
	rsp, err := s.RepayAll(context.Background(), params)

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("RawData: %+v", string(rsp.RawData))

	if err := rsp.ReadData(nil); err != nil {
		t.Fatal(err)
	}
}

func TestApiService_RepaySingle(t *testing.T) {
	t.SkipNow()
	s := NewApiServiceFromEnv()
	params := map[string]string{"currency": "BTC", "tradeId": "5e0dca1fdd28950009db2530", "size": "0.00001"}
	rsp, err := s.RepaySingle(context.Background(), params)

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("RawData: %+v", string(rsp.RawData))

	if err := rsp.ReadData(nil); err != nil {
		t.Fatal(err)
	}
}

func TestApiService_CreateLendOrder(t *testing.T) {
	t.SkipNow()
	s := NewApiServiceFromEnv()
	params := map[string]string{"currency": "BTC", "size": "0.02", "dailyIntRate": "0.002", "term": "28"}
	rsp, err := s.CreateLendOrder(context.Background(), params)

	if err != nil {
		t.Fatal(err)
	}

	o := &CreateLendOrderResultModel{}
	if err := rsp.ReadData(o); err != nil {
		t.Fatal(err)
	}

	t.Log(ToJsonString(o))

	switch {
	case o.OrderId == "":
		t.Error("empty key 'OrderId'")
	}
}

func TestApiService_CancelLendOrder(t *testing.T) {
	t.SkipNow()
	s := NewApiServiceFromEnv()

	rsp, err := s.CancelLendOrder(context.Background(), "5e0ebc406817f00008a314f3")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("RawData: %+v", string(rsp.RawData))

	if err := rsp.ReadData(nil); err != nil {
		t.Fatal(err)
	}
}

func TestApiService_ToggleAutoLend(t *testing.T) {
	t.SkipNow()
	s := NewApiServiceFromEnv()

	params := map[string]string{"currency": "BTC", "isEnable": "true", "retainSize": "0.4", "dailyIntRate": "0.002", "term": "14"}
	rsp, err := s.ToggleAutoLend(context.Background(), params)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("RawData: %+v", string(rsp.RawData))

	if err := rsp.ReadData(nil); err != nil {
		t.Fatal(err)
	}
}

func TestApiService_LendActiveOrders(t *testing.T) {
	s := NewApiServiceFromEnv()
	pagination := &PaginationParam{
		CurrentPage: 1,
		PageSize:    10,
	}

	rsp, err := s.LendActiveOrders(context.Background(), "BTC", pagination)
	if err != nil {
		t.Fatal(err)
	}

	os := LendActiveOrdersModel{}
	if _, err := rsp.ReadPaginationData(&os); err != nil {
		t.Fatal(err)
	}

	t.Log(ToJsonString(os))

	if len(os) == 0 {
		t.SkipNow()
	}

	for _, o := range os {
		switch {
		case o.OrderId == "":
			t.Error("empty key 'OrderId'")
		case o.Currency == "":
			t.Error("empty key 'Currency'")
		case o.Size == "":
			t.Error("empty key 'Size'")
		case o.FilledSize == "":
			t.Error("empty key 'FilledSize'")
		case o.DailyIntRate == "":
			t.Error("empty key 'DailyIntRate'")
		case o.Term == "":
			t.Error("empty key 'Term'")
		case o.CreatedAt == "":
			t.Error("empty key 'CreatedAt'")
		}
	}
}

func TestApiService_LendDoneOrders(t *testing.T) {
	s := NewApiServiceFromEnv()
	pagination := &PaginationParam{
		CurrentPage: 1,
		PageSize:    10,
	}

	rsp, err := s.LendDoneOrders(context.Background(), "BTC", pagination)
	if err != nil {
		t.Fatal(err)
	}

	os := LendDoneOrdersModel{}
	if _, err := rsp.ReadPaginationData(&os); err != nil {
		t.Fatal(err)
	}

	t.Log(ToJsonString(os))

	if len(os) == 0 {
		t.SkipNow()
	}

	for _, o := range os {
		switch {
		case o.OrderId == "":
			t.Error("empty key 'OrderId'")
		case o.Currency == "":
			t.Error("empty key 'Currency'")
		case o.Size == "":
			t.Error("empty key 'Size'")
		case o.FilledSize == "":
			t.Error("empty key 'FilledSize'")
		case o.DailyIntRate == "":
			t.Error("empty key 'DailyIntRate'")
		case o.Term == "":
			t.Error("empty key 'Term'")
		case o.CreatedAt == "":
			t.Error("empty key 'CreatedAt'")
		case o.Status == "":
			t.Error("empty key 'Status'")
		}
	}
}

func TestApiService_LendTradeUnsettledRecords(t *testing.T) {
	s := NewApiServiceFromEnv()
	pagination := &PaginationParam{
		CurrentPage: 1,
		PageSize:    10,
	}

	rsp, err := s.LendTradeUnsettledRecords(context.Background(), "BTC", pagination)
	if err != nil {
		t.Fatal(err)
	}

	os := LendTradeUnsettledRecordsModel{}
	if _, err := rsp.ReadPaginationData(&os); err != nil {
		t.Fatal(err)
	}

	t.Log(ToJsonString(os))

	if len(os) == 0 {
		t.SkipNow()
	}

	for _, o := range os {
		switch {
		case o.TradeId == "":
			t.Error("empty key 'TradeId'")
		case o.Currency == "":
			t.Error("empty key 'Currency'")
		case o.Size == "":
			t.Error("empty key 'Size'")
		case o.AccruedInterest == "":
			t.Error("empty key 'AccruedInterest'")
		case o.Repaid == "":
			t.Error("empty key 'Repaid'")
		case o.DailyIntRate == "":
			t.Error("empty key 'DailyIntRate'")
		case o.Term == "":
			t.Error("empty key 'Term'")
		case o.MaturityTime == "":
			t.Error("empty key 'MaturityTime'")
		}
	}
}

func TestApiService_LendTradeSettledRecords(t *testing.T) {
	s := NewApiServiceFromEnv()
	pagination := &PaginationParam{
		CurrentPage: 1,
		PageSize:    10,
	}

	rsp, err := s.LendTradeSettledRecords(context.Background(), "BTC", pagination)
	if err != nil {
		t.Fatal(err)
	}

	os := LendTradeSettledRecordsModel{}
	if _, err := rsp.ReadPaginationData(&os); err != nil {
		t.Fatal(err)
	}

	t.Log(ToJsonString(os))

	if len(os) == 0 {
		t.SkipNow()
	}

	for _, o := range os {
		switch {
		case o.TradeId == "":
			t.Error("empty key 'TradeId'")
		case o.Currency == "":
			t.Error("empty key 'Currency'")
		case o.Size == "":
			t.Error("empty key 'Size'")
		case o.Interest == "":
			t.Error("empty key 'Interest'")
		case o.Repaid == "":
			t.Error("empty key 'Repaid'")
		case o.DailyIntRate == "":
			t.Error("empty key 'DailyIntRate'")
		case o.Term == "":
			t.Error("empty key 'Term'")
		case o.SettledAt == "":
			t.Error("empty key 'SettledAt'")
		}
	}
}

func TestApiService_LendAssets(t *testing.T) {
	s := NewApiServiceFromEnv()

	rsp, err := s.LendAssets(context.Background(), "BTC")
	if err != nil {
		t.Fatal(err)
	}

	os := &LendAssetsModel{}
	if err := rsp.ReadData(os); err != nil {
		t.Fatal(err)
	}

	t.Log(ToJsonString(os))

	for _, o := range *os {
		switch {
		case o.Currency == "":
			t.Error("empty key 'Currency'")
		case o.Outstanding == "":
			t.Error("empty key 'Outstanding'")
		case o.FilledSize == "":
			t.Error("empty key 'FilledSize'")
		case o.AccruedInterest == "":
			t.Error("empty key 'AccruedInterest'")
		case o.RealizedProfit == "":
			t.Error("empty key 'RealizedProfit'")
		}
	}
}

func TestApiService_MarginMarkets(t *testing.T) {
	s := NewApiServiceFromEnv()

	params := map[string]string{"currency": "BTC"}
	rsp, err := s.MarginMarkets(context.Background(), params)
	if err != nil {
		t.Fatal(err)
	}

	os := &MarginMarketsModel{}
	if err := rsp.ReadData(os); err != nil {
		t.Fatal(err)
	}

	t.Log(ToJsonString(os))

	for _, o := range *os {
		switch {
		case o.DailyIntRate == "":
			t.Error("empty key 'DailyIntRate'")
		case o.Term == "":
			t.Error("empty key 'Term'")
		case o.Size == "":
			t.Error("empty key 'Size'")
		}
	}
}

func TestApiService_MarginTradeLast(t *testing.T) {
	s := NewApiServiceFromEnv()

	rsp, err := s.MarginTradeLast(context.Background(), "BTC")
	if err != nil {
		t.Fatal(err)
	}

	os := &MarginTradesModel{}
	if err := rsp.ReadData(os); err != nil {
		t.Fatal(err)
	}

	t.Log(ToJsonString(os))

	for _, o := range *os {
		switch {
		case o.TradeId == "":
			t.Error("empty key 'TradeId'")
		case o.Currency == "":
			t.Error("empty key 'Currency'")
		case o.Size == "":
			t.Error("empty key 'Size'")
		case o.DailyIntRate == "":
			t.Error("empty key 'DailyIntRate'")
		case o.Term == "":
			t.Error("empty key 'Term'")
		case o.Timestamp == "":
			t.Error("empty key 'Timestamp'")
		}
	}
}

func TestApiService_MarginRiskLimit(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.MarginRiskLimit(context.Background(), "cross")
	if err != nil {
		t.Fatal(err)
	}

	os := &MarginRiskLimitModel{}
	if err := rsp.ReadData(os); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(os))
}

func TestApiService_MarginIsolatedSymbols(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.MarginIsolatedSymbols(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	os := &MarginIsolatedSymbolsModel{}
	if err := rsp.ReadData(os); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(os))
}

func TestApiService_MarginIsolatedAccounts(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.MarginIsolatedAccounts(context.Background(), "USDT")
	if err != nil {
		t.Fatal(err)
	}

	os := &MarginIsolatedAccountsModel{}
	if err := rsp.ReadData(os); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(os))
}

func TestApiService_MarginIsolatedAccount(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.IsolatedAccount(context.Background(), "BTC-USDT")
	if err != nil {
		t.Fatal(err)
	}

	os := &MarginIsolatedAccountAssetsModel{}
	if err := rsp.ReadData(os); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(os))
}

func TestApiService_MarginIsolatedBorrow(t *testing.T) {
	t.SkipNow()
	s := NewApiServiceFromEnv()
	p := map[string]string{
		"symbol":         "MATIC-USDT",
		"currency":       "MATIC",
		"size":           "1",
		"borrowStrategy": "FOK",
		"maxRate":        "",
		"period":         "7",
	}
	rsp, err := s.IsolatedBorrow(context.Background(), p)
	if err != nil {
		t.Fatal(err)
	}

	os := &MarginIsolatedBorrowRes{}
	if err := rsp.ReadData(os); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(os))
}

func TestApiService_IsolatedBorrowOutstandingRecord(t *testing.T) {
	s := NewApiServiceFromEnv()
	p := map[string]string{
		"symbol":   "MATIC-USDT",
		"currency": "MATIC",
	}

	page := &PaginationParam{PageSize: 10, CurrentPage: 1}
	rsp, err := s.IsolatedBorrowOutstandingRecord(context.Background(), p, page)
	if err != nil {
		t.Fatal(err)
	}

	os := &IsolatedBorrowOutstandingRecordsModel{}
	if _, err := rsp.ReadPaginationData(os); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(os))
}

func TestApiService_IsolatedBorrowRepaidRecord(t *testing.T) {
	s := NewApiServiceFromEnv()
	p := map[string]string{
		"symbol":   "MATIC-USDT",
		"currency": "MATIC",
	}

	page := &PaginationParam{PageSize: 10, CurrentPage: 1}
	rsp, err := s.IsolatedBorrowRepaidRecord(context.Background(), p, page)
	if err != nil {
		t.Fatal(err)
	}

	os := &IsolatedBorrowRepaidRecordRecordsModel{}
	if _, err := rsp.ReadPaginationData(os); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(os))
}

func TestApiService_IsolatedRepayAll(t *testing.T) {
	t.SkipNow()
	s := NewApiServiceFromEnv()
	p := map[string]string{
		"currency":    "BTC",
		"seqStrategy": "HIGHEST_RATE_FIRST",
		"size":        "1.9",
		"symbol":      "BTC-USDT",
	}
	rsp, err := s.IsolatedRepayAll(context.Background(), p)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("RawData: %+v", string(rsp.RawData))

	if err := rsp.ReadData(nil); err != nil {
		t.Fatal(err)
	}
}

func TestApiService_IsolatedRepaySingle(t *testing.T) {
	t.SkipNow()
	s := NewApiServiceFromEnv()
	p := map[string]string{
		"currency": "MATIC",
		"loanId":   "64993793ad39960001cceeb6",
		"size":     "0.00000833",
		"symbol":   "MATIC-USDT",
	}
	rsp, err := s.IsolatedRepaySingle(context.Background(), p)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("RawData: %+v", string(rsp.RawData))

	if err := rsp.ReadData(nil); err != nil {
		t.Fatal(err)
	}
}

func TestApiService_MarginCurrencyInfo(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.MarginCurrencyInfo(context.Background(), "BTCUP")
	if err != nil {
		t.Fatal(err)
	}
	c := &MarginCurrenciesModel{}
	if err := rsp.ReadData(c); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(c))

}

func TestApiService_MarginCurrencies(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.MarginCurrencies(context.Background(), "BTC", "BTC-USDT", "true")
	if err != nil {
		t.Fatal(err)
	}
	c := &IsolatedCurrenciesRiskLimitModel{}
	if err := rsp.ReadData(c); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(c))
}

func TestApiService_MarginBorrowV3(t *testing.T) {
	s := NewApiServiceFromEnv()

	p := &MarginBorrowV3Req{
		IsIsolated:  false,
		Currency:    "USDT",
		Size:        "10",
		TimeInForce: "FOK",
		IsHf:        false,
	}
	rsp, err := s.MarginBorrowV3(context.Background(), p)

	if err != nil {
		t.Fatal(err)
	}

	o := &MarginBorrowV3Res{}
	if err := rsp.ReadData(o); err != nil {
		t.Fatal(err)
	}

	t.Log(ToJsonString(o))

	switch {
	case o.OrderNo == "":
		t.Error("empty key 'OrderId'")
	}
}

func TestApiService_QueryMarginBorrowV3(t *testing.T) {
	s := NewApiServiceFromEnv()

	p := map[string]string{
		"currency": "USDT",
	}

	pagination := &PaginationParam{
		CurrentPage: 1,
		PageSize:    10,
	}
	rsp, err := s.QueryMarginBorrowV3(context.Background(), p, pagination)

	if err != nil {
		t.Fatal(err)
	}

	o := &MarginBorrowsV3Model{}
	if _, err := rsp.ReadPaginationData(o); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(o))
}

func TestApiService_MarginRepayV3(t *testing.T) {
	s := NewApiServiceFromEnv()

	p := &MarginRepay3VReq{
		IsIsolated: false,
		Currency:   "USDT",
		Size:       "10",
		IsHf:       false,
	}
	rsp, err := s.MarginRepayV3(context.Background(), p)

	if err != nil {
		t.Fatal(err)
	}

	o := &MarginRepayV3Res{}
	if err := rsp.ReadData(o); err != nil {
		t.Fatal(err)
	}

	t.Log(ToJsonString(o))

	switch {
	case o.OrderNo == "":
		t.Error("empty key 'OrderId'")
	}
}

func TestApiService_QueryMarginRepayV3(t *testing.T) {
	s := NewApiServiceFromEnv()

	p := map[string]string{
		"currency": "USDT",
	}
	pagination := &PaginationParam{
		CurrentPage: 1,
		PageSize:    10,
	}
	rsp, err := s.QueryMarginRepayV3(context.Background(), p, pagination)

	if err != nil {
		t.Fatal(err)
	}

	o := &MarginRepaysV3Model{}
	if _, err := rsp.ReadPaginationData(o); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(o))
}

func TestApiService_QueryInterestV3(t *testing.T) {
	s := NewApiServiceFromEnv()

	p := map[string]string{
		"currency": "USDT",
	}
	pagination := &PaginationParam{
		CurrentPage: 1,
		PageSize:    10,
	}
	rsp, err := s.QueryInterestV3(context.Background(), p, pagination)

	if err != nil {
		t.Fatal(err)
	}

	o := &MarginInterestsV3Model{}
	if _, err := rsp.ReadPaginationData(o); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(o))
}

func TestApiService_QueryMarginCurrenciesV3(t *testing.T) {
	s := NewApiServiceFromEnv()

	rsp, err := s.QueryMarginCurrenciesV3(context.Background(), "USDT")

	if err != nil {
		t.Fatal(err)
	}

	o := &MarginCurrenciesV3Model{}
	if err := rsp.ReadData(o); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(o))
}

func TestApiService_QueryMarginInterestRateV3(t *testing.T) {
	s := NewApiServiceFromEnv()

	rsp, err := s.QueryMarginInterestRateV3(context.Background(), "USDT")

	if err != nil {
		t.Fatal(err)
	}

	o := &MarginInterestRatesV3Model{}
	if err := rsp.ReadData(o); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(o))
}

func TestApiService_LendingPurchaseV3(t *testing.T) {
	s := NewApiServiceFromEnv()
	p := &LendingPurchaseV3Req{
		Currency:     "BOME",
		Size:         "100",
		InterestRate: "0.083",
	}
	rsp, err := s.LendingPurchaseV3(context.Background(), p)

	if err != nil {
		t.Fatal(err)
	}

	o := &LendingV3Res{}
	if err := rsp.ReadData(o); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(o))
}

func TestApiService_LendingRedeemV3(t *testing.T) {
	s := NewApiServiceFromEnv()
	p := &LendingRedeemV3Req{
		Currency:        "BOME",
		Size:            "100",
		PurchaseOrderNo: "6698d0ab08b9bb0007052e29",
	}
	rsp, err := s.LendingRedeemV3(context.Background(), p)

	if err != nil {
		t.Fatal(err)
	}

	o := &LendingV3Res{}
	if err := rsp.ReadData(o); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(o))
}

func TestApiService_LendingPurchaseUpdateV3(t *testing.T) {
	s := NewApiServiceFromEnv()
	p := &LendingPurchaseUpdateV3Req{
		Currency:        "BOME",
		PurchaseOrderNo: "6698d0ab08b9bb0007052e29",
		InterestRate:    "0.084",
	}
	_, err := s.LendingPurchaseUpdateV3(context.Background(), p)

	if err != nil {
		t.Fatal(err)
	}
}

func TestApiService_RedemptionOrdersV3(t *testing.T) {
	s := NewApiServiceFromEnv()

	pagination := &PaginationParam{
		CurrentPage: 1,
		PageSize:    10,
	}
	rsp, err := s.RedemptionOrdersV3(context.Background(), "BOME", "DONE", pagination)

	if err != nil {
		t.Fatal(err)
	}

	o := &RedemptionOrdersV3Model{}
	if _, err := rsp.ReadPaginationData(o); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(o))
}

func TestApiService_SubscriptionOrdersV3(t *testing.T) {
	s := NewApiServiceFromEnv()

	pagination := &PaginationParam{
		CurrentPage: 1,
		PageSize:    10,
	}
	rsp, err := s.SubscriptionOrdersV3(context.Background(), "BOME", "PENDING", pagination)

	if err != nil {
		t.Fatal(err)
	}

	o := &SubscriptionOrdersV3Model{}
	if _, err := rsp.ReadPaginationData(o); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(o))
}

func TestApiService_MarginSymbolsV3(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.MarginSymbolsV3(context.Background())

	if err != nil {
		t.Fatal(err)
	}

	o := &MarginSymbolsV3Model{}
	if err := rsp.ReadData(o); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(o))
}

func TestApiService_UpdateUserLeverageV3(t *testing.T) {
	s := NewApiServiceFromEnv()
	p := &UpdateUserLeverageV3Model{
		Symbol:     "BTC-USDT",
		Leverage:   "1",
		IsIsolated: false,
	}
	_, err := s.UpdateUserLeverageV3(context.Background(), p)

	if err != nil {
		t.Fatal(err)
	}
}
