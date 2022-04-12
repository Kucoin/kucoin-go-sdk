package main

import (
	"log"

	"github.com/Mreyna3/kucoin-go-sdk"
)

func main() {
	//s := kucoin.NewApiServiceFromEnv()
	s := kucoin.NewApiService(
		kucoin.ApiKeyOption("key"),
		kucoin.ApiSecretOption("secret"),
		kucoin.ApiPassPhraseOption("passphrase"),
	)
	serverTime(s)
	accounts(s)
	orders(s)
	publicWebsocket(s)
	privateWebsocket(s)
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
func publicWebsocket(s *kucoin.ApiService) {
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

	c := s.NewWebSocketClient(tk)

	mc, ec, err := c.Connect()
	if err != nil {
		// Handle error
		return
	}

	ch1 := kucoin.NewSubscribeMessage("/market/ticker:KCS-BTC", false)
	ch2 := kucoin.NewSubscribeMessage("/market/ticker:ETH-BTC", false)
	uch := kucoin.NewUnsubscribeMessage("/market/ticker:ETH-BTC", false)

	if err := c.Subscribe(ch1, ch2); err != nil {
		// Handle error
		return
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
			log.Printf("Received: %s", kucoin.ToJsonString(msg))
			i++
			if i == 5 {
				log.Println("Unsubscribe ETH-BTC")
				if err = c.Unsubscribe(uch); err != nil {
					log.Printf("Error: %s", err.Error())
					// Handle error
					return
				}
			}
			if i == 10 {
				log.Println("Subscribe ETH-BTC")
				if err = c.Subscribe(ch2); err != nil {
					log.Printf("Error: %s", err.Error())
					// Handle error
					return
				}
			}
			if i == 15 {
				log.Println("Exit subscription")
				c.Stop() // Stop subscribing the WebSocket feed
				return
			}
		}
	}
}

func privateWebsocket(s *kucoin.ApiService) {
	rsp, err := s.WebSocketPrivateToken()
	if err != nil {
		// Handle error
		return
	}

	tk := &kucoin.WebSocketTokenModel{}
	//tk.AcceptUserMessage = true
	if err := rsp.ReadData(tk); err != nil {
		// Handle error
		return
	}

	c := s.NewWebSocketClient(tk)

	mc, ec, err := c.Connect()
	if err != nil {
		// Handle error
		return
	}

	ch1 := kucoin.NewSubscribeMessage("/market/level3:BTC-USDT", false)
	ch2 := kucoin.NewSubscribeMessage("/account/balance", false)

	if err := c.Subscribe(ch1, ch2); err != nil {
		// Handle error
		return
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
			log.Printf("Received: %s", kucoin.ToJsonString(msg))
			i++
			if i == 10 {
				log.Println("Subscribe ETH-BTC")
				if err = c.Subscribe(ch2); err != nil {
					log.Printf("Error: %s", err.Error())
					// Handle error
					return
				}
			}
			if i == 15 {
				log.Println("Exit subscription")
				c.Stop() // Stop subscribing the WebSocket feed
				return
			}
		}
	}
}
