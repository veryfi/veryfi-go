# veryfi-go

The Go client for communicating with the Veryfi OCR API.

## Installing 

This package can be installed via go get

```
go get github.com/hoanhan101/veryfi-go
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
help                 Print usage information
lint                 Lint project source files
test-integration     Run integration tests
test-unit            Run unit tests
version              Print the version
```

### Go API Client Library

The **veryfi** library can be used to communicate with Veryfi API. All available functionality is described here <https://veryfi.github.io/veryfi-python/reference/veryfi/#client>

Below is the sample script using **veryfi** to OCR and extract data from a document:

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
