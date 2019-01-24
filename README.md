# GO SDK for KuCoin API
> The detailed document [https://docs.kucoin.com](https://docs.kucoin.com).

## Install

```bash
go get github.com/Kucoin/kucoin-go-sdk
```

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

```

- Example of API `with` authentication

```go
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
	rsp, err := s.Accounts("", "")
	if err != nil {
		// Handle error
	}

	as := kucoin.AccountsModel{}
	if err := rsp.ReadData(&as); err != nil {
		// Handle error
	}

	for _, a := range as {
		log.Printf("Available balance: %s %s => %s", a.Type, a.Currency, a.Available)
	}
}
```

```go
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
	rsp, err := s.Orders(map[string]string{}, &kucoin.PaginationParam{CurrentPage: 1, PageSize: 10})
	if err != nil {
		// Handle error
	}

	os := kucoin.OrdersModel{}
	pa, err := rsp.ReadPaginationData(&os)
	if err != nil {
		// Handle error
	}
	log.Printf("Total num: %d, total page: %d", pa.TotalNum, pa.TotalPage)
	for _, o := range os {
		log.Printf("Order: %s, %s, %s", o.Id, o.Type, o.Price)
	}
}
```

- Example of WebSocket feed
> Require [gorilla/websocket](https://github.com/gorilla/websocket)

```bash
go get github.com/gorilla/websocket
```

```go
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
```

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
