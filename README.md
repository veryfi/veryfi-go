<img src="https://user-images.githubusercontent.com/30125790/212157461-58bdc714-2f89-44c2-8e4d-d42bee74854e.png#gh-dark-mode-only" width="200">
<img src="https://user-images.githubusercontent.com/30125790/212157486-bfd08c5d-9337-4b78-be6f-230dc63838ba.png#gh-light-mode-only" width="200">

[![Go Reference](https://pkg.go.dev/badge/github.com/veryfi/veryfi-go/veryfi.svg)](https://pkg.go.dev/github.com/veryfi/veryfi-go/veryfi)
[![Go Report Card](https://goreportcard.com/badge/github.com/veryfi/veryfi-go)](https://goreportcard.com/report/github.com/veryfi/veryfi-go)
[![Test](https://github.com/veryfi/veryfi-go/actions/workflows/test.yml/badge.svg?branch=master)](https://github.com/veryfi/veryfi-go/actions/workflows/test.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)


**veryfi-go** is a Go module for communicating with the [Veryfi OCR API](https://veryfi.com/api/).


## Table of Contents
- [Installing](#installing)
- [Getting Started](#getting-started)
  - [Obtaining Client ID and user keys](#obtaining-client-id-and-user-keys)
  - [Setting up the project](#setting-up-the-project)
  - [Go API Client Library](#go-api-client-library)
  - [Testing](#testing)
- [Need Help?](#need-help)
- [Tutorial](#tutorial)


## Installing 


This package can be installed by cloning this directory:

```
git clone https://github.com/veryfi/veryfi-go.git
```

or using `go get`:
```
go get github.com/veryfi/veryfi-go
```


## Getting Started


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


The **Veryfi** library can be used to communicate with Veryfi API. All available functionality is described here <https://pkg.go.dev/github.com/veryfi/veryfi-go/veryfi>.

Below is the sample script using **Veryfi** to OCR and extract data from a document:

```go
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/veryfi/veryfi-go/veryfi"
	"github.com/veryfi/veryfi-go/veryfi/scheme"
)

func main() {
	timeout, err := time.ParseDuration("10s")
	if err != nil {
		log.Fatal(err)
	}

	client, err := veryfi.NewClientV8(&veryfi.Options{
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
		FileURL: "YOUR_HOSTED_FILE_URL",
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

A successful response will look something like this:
```
&{ABNNumber: AccountNumber: BillToAddress: BillToName: BillToVATNumber: CardNumber: Category: Created:2021-05-20 19:21:38 CurrencyCode:USD Date:2019-02-26 00:00:00 DeliveryDate: Discount:0 DocumentReferenceNumber: DueDate:2019-02-26 ExternalID: ID:23002226 ImgFileName:3947f571-a41b-4b79-abc7-c0d9805c8610.png ImgThumbnailURL:https://scdn.veryfi.com/receipts/3947f571-a41b-4b79-abc7-c0d9805c8610_1_t.png?Expires=1621538559&Signature=BokBYv9jyJcXbCXu49DqxHwRdAWEgG8xfMw7LHujXSCA5y4kGd-QaBDwMzMCgCuM0Ezdrv3lgAZa0Cr8A5DKAzymXxnfdEiV46w~iy1zGPRgx6IkqvllB4bWqHFdwuu88CJarfIjvkcaygcECiFHg3RSKuuN4eGUYDP~fK8ER~Awb9Cr5FpTbTMc9kOfyc~vii2Mikg3TBiTbcdshhjgD2oRI4nFh1fpwRpfHAArIR-ijYAetjFEOQycUiu6WnzWAyEV9RCP9KcrKOnY5eKD-mm5mKuGQGXX1OT2AGw80klF1epx7XppeER9kALF1s8Dq87s8gdnnVsrstEF3~e8Yg__&Key-Pair-Id=APKAJCILBXEJFZF4DCHQ ImgURL:https://scdn.veryfi.com/receipts/3947f571-a41b-4b79-abc7-c0d9805c8610.png?Expires=1621538559&Signature=G7T6n7~Gpr1Pi5rfPRn1GoOeTlKZnVxLbWSZf~svnNpytILXvN9tg7y-Ib39lcifHeM6vjVfm4Pa4k63-ri~SySGFq-RWtF4IjQGM3Hw4~8wHB-sPhorn4JeVd~e~CpaUgFJbGSRnbb1cmBDFdkuBMbLkdC7m5ifwE10kanUU87Q~vpDYLkQINzfylHJk21rwtSPvIiEX8rudLK1F1BGl7TWvx-o7BT~PTCJ-RsA~j4eGuOprDXpt5Achpf-LMUa-iRCpMFupWVOZFPGln8rDqp-TcpryTawTbNlajg0nFDtF1eqBlbfoEycb-ZECtV4KECZtle5T7rBqhGQsmUxNQ__&Key-Pair-Id=APKAJCILBXEJFZF4DCHQ Insurance: InvoiceNumber:   US-001 IsDuplicate:0 LineItems:[] OCRText:

INVOICE                         LOGO
East Repair Inc.
1912 Harvest Lane
New York, NY 12210
BILL TO SHIP TO INVOICE #       US-001
John Smith      John Smith      INVOICE DATE    11/02/2019
2 Court Square  3787 Pineview Drive
New York, NY 12210      Cambridge, MA 12210     P.O.    2312/2019
        DUE DATE        26/02/2019

QTY     DESCRIPTION     UNIT PRICE      AMOUNT

1       Front and rear brake cables             100.00  100.00
2 New set of pedal arms         15.00   30.00
3       Labor 3hrs                      5.00    15.00

        Subtotal        145.00
        Sales Tax 6.25% 9.06
        TOTAL   $154.06

        TERMS & CONDITIONS
Thank you       , Please make checks payable to: East Repair Inc.
        Payment is due within 15 days
        John Smith
 OrderDate: PaymentDisplayName:No Payment, PaymentTerms: PaymentType:no_payment, PhoneNumber: PurchaseOrderNumber: Rounding:0 ServiceEndDate: ServiceStartDate: ShipDate: ShipToAddress: ShipToName: Shipping:0 StoreNumber: Subtotal:145 Tax:9.06 TaxLines:[] Tip:0 Total:154.06 TotalWeight: TrackingNumber: Updated:2021-05-20 19:21:39 VATNumber: Vendor:{Address:1912 harvest lane new york, ny 12210 2 court square    3787 pineview drive Category: Email: FaxNumber: Name: PhoneNumber: RawName: VendorLogo: VendorRegNumber: VendorType: Web:} VendorAccountNumber: VendorBankName: VendorBankNumber: VendorBankSwift: VendorIban:}%
```

For more examples about different methods to process documents, refer to the [documentation's examples](https://pkg.go.dev/github.com/veryfi/veryfi-go/veryfi#pkg-examples).


### Testing


To run unit tests:
```
make test-unit
```

To run integration tests, supply your `CLIENT_ID`, `CLIENT_SECRET`, `USERNAME`, and `API_KEY` environment variables in [Makefile](Makefile) and run `make test-integration`:
```
.PHONY: test-integration
test-integration:  ## Run integration tests
	CLIENT_ID=FIXME CLIENT_SECRET=FIXME USERNAME=FIXME API_KEY=FIXME go test -race -cover -run Integration -coverprofile=coverage.out -covermode=atomic ./...
```


## Need Help?

Visit https://docs.veryfi.com/ to access integration guides and usage notes in the Veryfi API Documentation Portal

If you run into any issue or need help installing or using the library, please contact <support@veryfi.com>.

If you found a bug in this library or would like new features added, then open an issue or pull requests against this repo!

To learn more about Veryfi visit <https://www.veryfi.com/>.


## Tutorial


Below is a introduction to the Go SDK. We're gonna walkthrough a problem and solve it together:

> Let’s say we are faced with a challenge where we have to capture and extract data from thousands (or even millions) of backlogged receipts. We don’t want to do it manually because copying and pasting from files to files takes a lot of time and is often error-prone. Also, different receipts have different forms and styles, finding the meaningful data across a huge pile of files can be a tedious task. What is even better is if we could extract the receipts and see the processing in real-time where the results can be returned almost immediately in seconds, not hours. So now, what is the best way that we can solve this?
> 
If you prefer a video format, here is the link to our Youtube channel →
[Link to blog post →](https://www.veryfi.com/go/)


## Releases

For a patch or minor version release, create a tagged release.

For a major version release, update `github.com/veryfi/veryfi-go/v` in the code, create a tagged release.
