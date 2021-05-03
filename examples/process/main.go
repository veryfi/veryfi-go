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

	testfile := "./veryfi/testdata/invoice1.png"

	// Method 1: Process Uploaded Document
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

	// Method 2: Process Base64 Encoded Document
	encodedFile, err := veryfi.Base64EncodeFile(testfile)
	resp, err = client.ProcessDocumentUploadBase64(scheme.DocumentUploadBase64Options{
		FileData: encodedFile,
		DocumentSharedOptions: scheme.DocumentSharedOptions{
			FileName: "invoice1.png",
			Tags:     []string{"example", "test", "upload", "base64"},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Method 2 - ID: %v\tTotal: %v %v\n", resp.ID, resp.Total, resp.CurrencyCode)

	// Method 3: Process Document via URL
	resp, err = client.ProcessDocumentURL(scheme.DocumentURLOptions{
		FileURL: "https://templates.invoicehome.com/invoice-template-us-neat-750px.png",
		DocumentSharedOptions: scheme.DocumentSharedOptions{
			Tags: []string{"example", "test", "url"},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Method 3 - ID: %v\tTotal: %v %v\n", resp.ID, resp.Total, resp.CurrencyCode)
}
