package veryfi

import (
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
