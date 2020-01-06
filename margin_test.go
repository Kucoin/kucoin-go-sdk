package kucoin

import (
	"math/rand"
	"testing"
	"time"
)

func TestApiService_CurrentMarkPrice(t *testing.T) {
	symbols := []string{"USDT-BTC", "ETH-BTC", "LTC-BTC", "EOS-BTC", "XRP-BTC", "KCS-BTC"}
	rand.Seed(time.Now().UnixNano())
	symbol := symbols[rand.Intn(len(symbols))]

	s := NewApiServiceFromEnv()
	rsp, err := s.CurrentMarkPrice(symbol)

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
	rsp, err := s.MarginConfig()

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
	rsp, err := s.MarginAccount()

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
	rsp, err := s.CreateBorrowOrder(params)

	if err != err {
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
	rsp, err := s.BorrowOrder("5e12f29dd43f8d0008c87981")

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

	rsp, err := s.BorrowOutstandingRecords("", pagination)
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

	rsp, err := s.BorrowRepaidRecords("", pagination)
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
	rsp, err := s.RepayAll(params)

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
	rsp, err := s.RepaySingle(params)

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
	rsp, err := s.CreateLendOrder(params)

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

	rsp, err := s.CancelLendOrder("5e0ebc406817f00008a314f3")
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
	rsp, err := s.ToggleAutoLend(params)
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

	rsp, err := s.LendActiveOrders("BTC", pagination)
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

	rsp, err := s.LendDoneOrders("BTC", pagination)
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

	rsp, err := s.LendTradeUnsettledRecords("BTC", pagination)
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

	rsp, err := s.LendTradeSettledRecords("BTC", pagination)
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

	rsp, err := s.LendAssets("BTC")
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
	rsp, err := s.MarginMarkets(params)
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

	rsp, err := s.MarginTradeLast("BTC")
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
