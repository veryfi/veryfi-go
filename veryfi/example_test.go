package veryfi_test

import (
	"fmt"
	"log"
	"time"

	"github.com/veryfi/veryfi-go/veryfi"
	"github.com/veryfi/veryfi-go/veryfi/scheme"
)

func ExampleClient_processDocument() {
	timeout, err := time.ParseDuration("10s")
	if err != nil {
		log.Fatal(err)
	}

	// Initialize a Veryfi Client for v7 API.
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

	// Specify your document's filepath.
	testfile := "YOUR_TEST_FILEPATH"

	// Method 1: Process an uploaded document.
	resp, err := client.ProcessDocumentUpload(scheme.DocumentUploadOptions{
		FilePath: testfile,
		DocumentSharedOptions: scheme.DocumentSharedOptions{
			FileName: "invoice1.png",
			Tags:     []string{"example", "test", "upload"},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Method 1 - ID: %v\tTotal: %v %v\n", resp.ID, resp.Total, resp.CurrencyCode)

	// Method 2: Process a document via an URL.
	resp, err = client.ProcessDocumentURL(scheme.DocumentURLOptions{
		FileURL: "YOUR_INVOICE_URL",
		DocumentSharedOptions: scheme.DocumentSharedOptions{
			Tags: []string{"example", "test", "url"},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Method 2 - ID: %v\tTotal: %v %v\n", resp.ID, resp.Total, resp.CurrencyCode)
}

func ExampleClient_manageDocument() {
	timeout, err := time.ParseDuration("10s")
	if err != nil {
		log.Fatal(err)
	}

	// Initialize a Veryfi Client for v7 API.
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

	// Specify what documents to update/delete.
	documentID := "YOUR_DOCUMENT_ID"
	deleteDocumentID := "YOUR_DOCUMENT_ID"

	// Update a document.
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

	// Search documents.
	docs, err := client.SearchDocuments(scheme.DocumentSearchOptions{
		Tag: "example",
	})
	if err != nil {
		log.Fatal(err)
	}
	for _, doc := range *docs {
		fmt.Printf("ID: %v\tBill To Name: %v\n", doc.ID, doc.BillToName)
	}

	// Get a document.
	resp, err = client.GetDocument(documentID, scheme.DocumentGetOptions{
		ReturnAuditTrail: "1",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID: %v\tBill To Name: %v\n", resp.ID, resp.BillToName)

	// Delete a document.
	err = client.DeleteDocument(deleteDocumentID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Successfully deleted %s\n", deleteDocumentID)
}
func ExampleClient_manageLineItem() {
	timeout, err := time.ParseDuration("10s")
	if err != nil {
		log.Fatal(err)
	}

	// Initialize a Veryfi Client for v7 API.
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

	// Specify what documents and line items to update/delete.
	documentID := "YOUR_DOCUMENT_ID"
	lineItemID := "YOUR_LINE_ITEM_ID"
	deleteLineItemID := "YOUR_LINE_ITEM_ID"

	// Add a line item.
	resp, err := client.AddLineItem(documentID, scheme.LineItemOptions{
		Order:       1,
		Description: "Example",
		Total:       1.0,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID: %v\t Description: %v\t Total: %v\n", resp.ID, resp.Description, resp.Total)

	// Get all line items.
	items, err := client.GetLineItems(documentID)
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range items.LineItems {
		fmt.Printf("ID: %v\t Description: %v\t Total: %v\n", item.ID, item.Description, item.Total)
	}

	// Get a line item.
	resp, err = client.GetLineItem(documentID, lineItemID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID: %v\t Description: %v\t Total: %v\n", resp.ID, resp.Description, resp.Total)

	// Update a line item.
	resp, err = client.UpdateLineItem(documentID, lineItemID, scheme.LineItemOptions{
		Order:       6,
		Description: "Example",
		Total:       6.6,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID: %v\t Description: %v\t Total: %v\n", resp.ID, resp.Description, resp.Total)

	// Delete a line item.
	err = client.DeleteLineItem(documentID, deleteLineItemID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Successfully deleted %s\n", deleteLineItemID)
}

func ExampleClient_manageTag() {
	timeout, err := time.ParseDuration("10s")
	if err != nil {
		log.Fatal(err)
	}

	// Initialize a Veryfi Client for v7 API.
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

	// Specify what documents and tag to update/delete.
	documentID := "YOUR_DOCUMENT_ID"
	deleteTagID := "YOUR_TAG_ID"

	// Add a tag.
	resp, err := client.AddTag(documentID, scheme.TagOptions{
		Name: "example3",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID: %v\t Name: %v\n", resp.ID, resp.Name)

	// Get all tags.
	tags, err := client.GetTags(documentID)
	if err != nil {
		log.Fatal(err)
	}
	for _, tag := range tags.Tags {
		fmt.Printf("ID: %v\t Name: %v\n", tag.ID, tag.Name)
	}

	// Get all global tags.
	tags, err = client.GetGlobalTags()
	if err != nil {
		log.Fatal(err)
	}
	for _, tag := range tags.Tags {
		fmt.Printf("ID: %v\t Name: %v\n", tag.ID, tag.Name)
	}

	// Delete a tag.
	err = client.DeleteTag(documentID, deleteTagID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Successfully deleted %s\n", deleteTagID)

	// Delete a tag globally.
	err = client.DeleteGlobalTag(deleteTagID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Successfully deleted %s\n", deleteTagID)
}
