package veryfi

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/hoanhan101/veryfi-go/veryfi/scheme"
	"github.com/stretchr/testify/assert"
)

func TestIntegrationFailNoAuth(t *testing.T) {
	client, err := NewClientV7(&Options{})
	assert.NotNil(t, client)
	assert.NoError(t, err)

	resp, err := client.ProcessDocumentURL(scheme.DocumentURLOptions{
		FileURL: "foo.com",
	})
	assert.Nil(t, resp)
	assert.Error(t, err)
}

func TestIntegrationFailInvalidClientID(t *testing.T) {
	client, err := NewClientV7(&Options{
		ClientID: "foo",
		Username: os.Getenv("USERNAME"),
		APIKey:   os.Getenv("API_KEY"),
	})
	assert.NotNil(t, client)
	assert.NoError(t, err)

	resp, err := client.ProcessDocumentURL(scheme.DocumentURLOptions{
		FileURL: "foo.com",
	})
	assert.Nil(t, resp)
	assert.Error(t, err)
}

func TestIntegrationFailInvalidUsername(t *testing.T) {
	client, err := NewClientV7(&Options{
		ClientID: os.Getenv("CLIENT_ID"),
		Username: "foo",
		APIKey:   os.Getenv("API_KEY"),
	})
	assert.NotNil(t, client)
	assert.NoError(t, err)

	resp, err := client.ProcessDocumentURL(scheme.DocumentURLOptions{
		FileURL: "foo.com",
	})
	assert.Nil(t, resp)
	assert.Error(t, err)
}

func TestIntegrationFailInvalidAPIKey(t *testing.T) {
	client, err := NewClientV7(&Options{
		ClientID: os.Getenv("CLIENT_ID"),
		Username: os.Getenv("USERNAME"),
		APIKey:   "foo",
	})
	assert.NotNil(t, client)
	assert.NoError(t, err)

	resp, err := client.ProcessDocumentURL(scheme.DocumentURLOptions{
		FileURL: "foo.com",
	})
	assert.Nil(t, resp)
	assert.Error(t, err)
}

func TestIntegrationFailInvalidDocument(t *testing.T) {
	timeout, _ := time.ParseDuration("10s")
	client, err := NewClientV7(&Options{
		ClientID: os.Getenv("CLIENT_ID"),
		Username: os.Getenv("USERNAME"),
		APIKey:   os.Getenv("API_KEY"),
		HTTP: HTTPOptions{
			Timeout: timeout,
		},
	})
	assert.NotNil(t, client)
	assert.NoError(t, err)

	resp, err := client.ProcessDocumentURL(scheme.DocumentURLOptions{
		FileURL: "foo.com",
	})
	assert.Nil(t, resp)
	assert.Error(t, err)
}

func TestIntegrationSuccessProcessDocumentURL(t *testing.T) {
	timeout, _ := time.ParseDuration("10s")
	client, err := NewClientV7(&Options{
		ClientID: os.Getenv("CLIENT_ID"),
		Username: os.Getenv("USERNAME"),
		APIKey:   os.Getenv("API_KEY"),
		HTTP: HTTPOptions{
			Timeout: timeout,
		},
	})
	assert.NotNil(t, client)
	assert.NoError(t, err)

	resp, err := client.ProcessDocumentURL(scheme.DocumentURLOptions{
		FileURL: "https://templates.invoicehome.com/invoice-template-us-neat-750px.png",
		DocumentSharedOptions: scheme.DocumentSharedOptions{
			Tags: []string{"integration", "test", "url"},
		},
	})
	assert.NotNil(t, resp)
	assert.NoError(t, err)
}

func TestIntegrationSuccessProcessDocumentUpload(t *testing.T) {
	timeout, _ := time.ParseDuration("10s")
	client, err := NewClientV7(&Options{
		ClientID: os.Getenv("CLIENT_ID"),
		Username: os.Getenv("USERNAME"),
		APIKey:   os.Getenv("API_KEY"),
		HTTP: HTTPOptions{
			Timeout: timeout,
		},
	})
	assert.NotNil(t, client)
	assert.NoError(t, err)

	testpath, _ := os.Getwd()
	resp, err := client.ProcessDocumentUpload(scheme.DocumentUploadOptions{
		FilePath: fmt.Sprintf("%s/testdata/invoice1.png", testpath),
		DocumentSharedOptions: scheme.DocumentSharedOptions{
			FileName: "test_invoice",
			Tags:     []string{"integration", "test", "upload"},
		},
	})
	assert.NotNil(t, resp)
	assert.NoError(t, err)
}
