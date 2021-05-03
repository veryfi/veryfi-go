[![Go Reference](https://pkg.go.dev/badge/github.com/hoanhan101/veryfi-go/veryfi.svg)](https://pkg.go.dev/github.com/hoanhan101/veryfi-go/veryfi)
[![Go Report Card](https://goreportcard.com/badge/github.com/hoanhan101/veryfi-go)](https://goreportcard.com/report/github.com/hoanhan101/veryfi-go)

# veryfi-go

The Go client for communicating with the Veryfi OCR API.

## Installing 

This package can be installed by cloning this directory:

```
git clone https://github.com/hoanhan101/veryfi-go.git
```

## Developing

### Obtaining Client ID and user keys

If you don't have an account with Veryfi, please go ahead and register here: [https://hub.veryfi.com/signup/api/](https://hub.veryfi.com/signup/api/)

### Setting up the project

For the ease of developing, a `Makefile` is included for several handy commands. Simply issue `make help` for more informations.
```
clean                Remove temporary files and build artifacts
cover                Run unit tests and open the coverage report
fmt                  Run gofmt on all files
github-tag           Create and push a tag with the current client version
help                 Print usage information
lint                 Lint project source files
test-integration     Run integration tests
test-unit            Run unit tests
version              Print the version
```

### Go API Client Library

The **Veryfi** library can be used to communicate with Veryfi API. All available functionality is described here <https://pkg.go.dev/github.com/hoanhan101/veryfi-go/veryfi>.

Below is the sample script using **Veryfi** to OCR and extract data from a document:

```go
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/hoanhan101/veryfi-go/veryfi"
	"github.com/hoanhan101/veryfi-go/veryfi/scheme"
)

func main() {
	timeout, err := time.ParseDuration("10s")
	if err != nil {
		log.Fatal(err)
	}

	client, err := veryfi.NewClientV7(&veryfi.Options{
		ClientID: "YOUR_CLIENT_ID",
		Username: "YOUR_USERNAME",
		APIKey:   "YOUR_API_KEY",
		HTTP: veryfi.HTTPOptions{
			Timeout: timeout,
			Retry: veryfi.RetryOptions{
				Count: 1,
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.ProcessDocumentURL(scheme.DocumentURLOptions{
		FileURL: "https://templates.invoicehome.com/invoice-template-us-neat-750px.png",
		DocumentSharedOptions: scheme.DocumentSharedOptions{
			Tags: []string{"electric", "repair", "ny"},
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", resp)
}
```

For more examples about different methods to process documents, refer to `examples` directory, or [examples/process/main.go](examples/process/main.go) specifically.

### Testing

To run unit tests:
```
make test-unit
```

To run integration tests, supply your `CLIENT_ID`, `USERNAME`, and `API_KEY` environment variables in [Makefile](Makefile) and run `make test-integration`:
```
.PHONY: test-integration
test-integration:  ## Run integration tests
	CLIENT_ID=FIXME USERNAME=FIXME API_KEY=FIXME go test -race -cover -run Integration -coverprofile=coverage.out -covermode=atomic ./...
```
