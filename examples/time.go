package main

import (
	"github.com/Kucoin/kucoin-go-sdk"
	"log"
)

func main() {
	// s := kucoin.NewApiServiceFromEnv()
	s := kucoin.NewApiService(
		kucoin.ApiBaseURIOption("https://openapi-v2.kucoin.com"), // Set the base uri, default "https://openapi-v2.kucoin.com" for production environment.
	)
	rsp, err := s.ServerTime()
	if err != nil {
		// Handle error
	}

	var ts int64
	if err := rsp.ReadData(&ts); err != nil {
		// Handle error
	}
	log.Printf("The server time: %d", ts)
}
