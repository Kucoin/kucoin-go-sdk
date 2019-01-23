package main

import (
	"log"

	"github.com/Kucoin/kucoin-go-sdk"
)

func main() {
	// s := kucoin.NewApiServiceFromEnv()
	s := kucoin.NewApiService()
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
