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
	lineItemID := "59874645"
	deleteLineItemID := "59993860"

	// Add a Line Item
	resp, err := client.AddLineItem(documentID, scheme.LineItemOptions{
		Order:       1,
		Description: "Example",
		Total:       1.0,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID: %v\t Description: %v\t Total: %v\n", resp.ID, resp.Description, resp.Total)

	// Get All Line Items
	items, err := client.GetLineItems(documentID)
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range items.LineItems {
		fmt.Printf("ID: %v\t Description: %v\t Total: %v\n", item.ID, item.Description, item.Total)
	}

	// Get a Line Item
	resp, err = client.GetLineItem(documentID, lineItemID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID: %v\t Description: %v\t Total: %v\n", resp.ID, resp.Description, resp.Total)

	// Update a Line Item
	resp, err = client.UpdateLineItem(documentID, lineItemID, scheme.LineItemOptions{
		Order:       6,
		Description: "Example",
		Total:       6.6,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID: %v\t Description: %v\t Total: %v\n", resp.ID, resp.Description, resp.Total)

	// Delete a Line Item
	err = client.DeleteLineItem(documentID, deleteLineItemID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Successfully deleted %s\n", deleteLineItemID)
}
