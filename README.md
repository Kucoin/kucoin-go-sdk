# GO SDK for KuCoin API
> The detailed document [https://docs.kucoin.com](https://docs.kucoin.com).


## Usage

- Choose environment

| Environment | BaseUri |
| -------- | -------- |
| *Production* `DEFAULT` | https://openapi-v2.kucoin.com |
| *Sandbox* | https://openapi-sandbox.kucoin.com |

```go
// Switch to the sandbox environment
s := kucoin.NewApiService(
    kucoin.ApiBaseURIOption("https://openapi-v2.kucoin.com"),
)
```

- Example of API `without` authentication

```go
package main
import (
	"github.com/Kucoin/kucoin-go-sdk"
	"log"
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
```

- Example of API `with` authentication

```go
package main
import (
	"github.com/Kucoin/kucoin-go-sdk"
	"log"
)
func main() {
	// s := kucoin.NewApiServiceFromEnv()
	s := kucoin.NewApiService(
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
```

- Example of WebSocket feed

```go
// Todo
```

- API list

| API | Authentication |
| -------- | -------- |
| Accounts() | YES |
| Account() | YES |
| CreateAccount() | YES |
| AccountHistories() | YES |
| AccountHolds() | YES |
| InnerTransfer() | YES |
| Currencies() | NO |
| Currency() | NO |
| Symbols() | NO |
| Ticker() | NO |
| PartOrderBook() | NO |
| AggregatedFullOrderBook() | NO |
| AtomicFullOrderBook() | NO |
| TradeHistories() | NO |
| HistoricRates() | NO |
| ServerTime() | NO |


## Run tests

```shell
# Add your API key into environmental variable first.
export API_BASE_URI=https://openapi-sandbox.kucoin.com
export API_KEY=key
export API_SECRET=secret
export API_PASSPHRASE=passphrase

// Run tests
go test -v
```

## License

[MIT](LICENSE)
