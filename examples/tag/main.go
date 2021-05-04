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
	deleteTagID := "782061"

	// Add a Tag
	resp, err := client.AddTag(documentID, scheme.TagOptions{
		Name: "example3",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID: %v\t Name: %v\n", resp.ID, resp.Name)

	// Get All Tags
	tags, err := client.GetTags(documentID)
	if err != nil {
		log.Fatal(err)
	}
	for _, tag := range tags.Tags {
		fmt.Printf("ID: %v\t Name: %v\n", tag.ID, tag.Name)
	}

	// Get All Global Tags
	tags, err = client.GetGlobalTags()
	if err != nil {
		log.Fatal(err)
	}
	for _, tag := range tags.Tags {
		fmt.Printf("ID: %v\t Name: %v\n", tag.ID, tag.Name)
	}

	// Delete a Tag
	err = client.DeleteTag(documentID, deleteTagID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Successfully deleted %s\n", deleteTagID)

	// Delete a Tag Globally
	err = client.DeleteGlobalTag(deleteTagID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Successfully deleted %s\n", deleteTagID)
}
