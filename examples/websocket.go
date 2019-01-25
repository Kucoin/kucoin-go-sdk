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

	mc, done, ec := s.WebSocketSubscribePublicChannel("/market/ticker:KCS-BTC", true)
	var i = 0
	type Ticker struct {
		Sequence    string `json:"sequence"`
		BestAsk     string `json:"bestAsk"`
		Size        string `json:"size"`
		BestBidSize string `json:"bestBidSize"`
		Price       string `json:"price"`
		BestAskSize string `json:"bestAskSize"`
		BestBid     string `json:"bestBid"`
	}
	for {
		select {
		case m := <-mc:
			// log.Printf("Received: %s", kucoin.ToJsonString(m))
			t := &Ticker{}
			if err := m.ReadData(t); err != nil {
				log.Printf("Failure to read: %s", err.Error())
				return
			}
			log.Printf("Ticker: %s, %s, %s", t.Sequence, t.Price, t.Size)
			i++
			if i == 3 {
				log.Println("Exit subscription")
				close(done)
				return
			}
		case err := <-ec:
			log.Printf("Error: %s", err.Error())
			close(done)
			return
		}
	}
}
