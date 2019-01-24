package main

import (
	"log"

	"github.com/Kucoin/kucoin-go-sdk"
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
	for {
		select {
		case m := <-mc:
			log.Printf("Received: %s", kucoin.ToJsonString(m))
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
