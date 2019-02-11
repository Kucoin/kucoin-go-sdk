package main

import (
	"log"

	"github.com/hhxsv5/kucoin-go-sdk"
)

func main() {
	// s := kucoin.NewApiServiceFromEnv()
	s := kucoin.NewApiService(
		kucoin.ApiKeyOption("key"),
		kucoin.ApiSecretOption("secret"),
		kucoin.ApiPassPhraseOption("passphrase"),
	)
	serverTime(s)
	accounts(s)
	orders(s)
	websocket(s)
}

func serverTime(s *kucoin.ApiService) {
	rsp, err := s.ServerTime()
	if err != nil {
		log.Printf("Error: %s", err.Error())
		// Handle error
		return
	}

	var ts int64
	if err := rsp.ReadData(&ts); err != nil {
		// Handle error
		return
	}
	log.Printf("The server time: %d", ts)
}

func accounts(s *kucoin.ApiService) {
	rsp, err := s.Accounts("", "")
	if err != nil {
		// Handle error
		return
	}

	as := kucoin.AccountsModel{}
	if err := rsp.ReadData(&as); err != nil {
		// Handle error
		return
	}

	for _, a := range as {
		log.Printf("Available balance: %s %s => %s", a.Type, a.Currency, a.Available)
	}
}

func orders(s *kucoin.ApiService) {
	rsp, err := s.Orders(map[string]string{}, &kucoin.PaginationParam{CurrentPage: 1, PageSize: 10})
	if err != nil {
		// Handle error
		return
	}

	os := kucoin.OrdersModel{}
	pa, err := rsp.ReadPaginationData(&os)
	if err != nil {
		// Handle error
		return
	}
	log.Printf("Total num: %d, total page: %d", pa.TotalNum, pa.TotalPage)
	for _, o := range os {
		log.Printf("Order: %s, %s, %s", o.Id, o.Type, o.Price)
	}
}
func websocket(s *kucoin.ApiService) {
	rsp, err := s.WebSocketPublicToken()
	if err != nil {
		// Handle error
		return
	}

	tk := &kucoin.WebSocketTokenModel{}
	if err := rsp.ReadData(tk); err != nil {
		// Handle error
		return
	}

	ch := kucoin.NewSubscribeMessage("/market/ticker:KCS-BTC", false, true)

	c := s.NewWebSocketClient(tk, ch)

	if err := c.Connect(); err != nil {
		// Handle error
		return
	}

	mc, ec := c.Subscribe()

	type Ticker struct {
		Sequence    string `json:"sequence"`
		BestAsk     string `json:"bestAsk"`
		Size        string `json:"size"`
		BestBidSize string `json:"bestBidSize"`
		Price       string `json:"price"`
		BestAskSize string `json:"bestAskSize"`
		BestBid     string `json:"bestBid"`
	}

	var i = 0
	for {
		select {
		case err := <-ec:
			c.Stop() // Stop subscribing the WebSocket feed
			log.Printf("Error: %s", err.Error())
			// Handle error
			return
		case msg := <-mc:
			// log.Printf("Received: %s", kucoin.ToJsonString(m))
			t := &Ticker{}
			if err := msg.ReadData(t); err != nil {
				log.Printf("Failure to read: %s", err.Error())
				return
			}
			log.Printf("Ticker: %s, %s, %s", t.Sequence, t.Price, t.Size)
			i++
			if i == 3 {
				log.Println("Exit subscription")
				c.Stop() // Stop subscribing the WebSocket feed
				return
			}
		}
	}
}
