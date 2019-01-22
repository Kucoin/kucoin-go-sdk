package main

import (
	"github.com/Kucoin/kucoin-go-sdk"
	"log"
)

func main() {
	// s := kucoin.NewApiServiceFromEnv()
	s := kucoin.NewApiService(
		kucoin.ApiBaseURIOption("https://openapi-v2.kucoin.com"), // Set the base uri, default "https://openapi-v2.kucoin.com" for production environment.
		kucoin.ApiKeyOption("key"),
		kucoin.ApiSecretOption("secret"),
		kucoin.ApiPassPhraseOption("passphrase"),
	)
	rsp, err := s.Accounts("", "")
	if err != nil {
		// Handle error
	}

	as := kucoin.AccountsModel{}
	if err := rsp.ReadData(&as); err != nil {
		// Handle error
	}

	for _, a := range as {
		log.Printf("The available balance: %s => %s", a.Currency, a.Available)
	}
}
