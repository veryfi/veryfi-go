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

	documentID := "35239377"

	// Update Document
	resp, err := client.UpdateDocument(documentID, scheme.DocumentUpdateOptions{
		BillToName:    "Hoanh An",
		BillToAddress: "NY",
		Vendor: scheme.VendorUpdateOptions{
			Name:    "Hoanh An",
			Address: "NY",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID: %v\tBill To Name: %v\n", resp.ID, resp.BillToName)

	// Search Documents
	docs, err := client.SearchDocuments(scheme.DocumentSearchOptions{
		Tag: "example",
	})
	if err != nil {
		log.Fatal(err)
	}
	for _, doc := range *docs {
		fmt.Printf("ID: %v\tBill To Name: %v\n", doc.ID, doc.BillToName)
	}
}
