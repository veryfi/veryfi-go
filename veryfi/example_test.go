package veryfi_test

import (
	"fmt"
	"log"
	"time"

	"github.com/hoanhan101/veryfi-go/veryfi"
	"github.com/hoanhan101/veryfi-go/veryfi/scheme"
)

func ExampleProcessDocumentURL() {
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

	resp, err = client.ProcessDocumentURL(scheme.DocumentURLOptions{
		FileURL: "YOUR_INVOICE_URL",
		DocumentSharedOptions: scheme.DocumentSharedOptions{
			Tags: []string{"example", "test", "url"},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID: %v\tTotal: %v %v\n", resp.ID, resp.Total, resp.CurrencyCode)
}
