package veryfi

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/veryfi/veryfi-go/veryfi/scheme"
)

func cleanUp(t *testing.T, documentID int) {
	client, err := NewClientV7(&Options{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Username:     os.Getenv("USERNAME"),
		APIKey:       os.Getenv("API_KEY"),
	})
	assert.NotNil(t, client)
	assert.NoError(t, err)

	err = client.DeleteDocument(strconv.Itoa(documentID))
	assert.NoError(t, err)
}

func TestIntegrationFailNoAuth(t *testing.T) {
	client, err := NewClientV7(&Options{})
	assert.NotNil(t, client)
	assert.NoError(t, err)

	resp, err := client.ProcessDocumentURL(scheme.DocumentURLOptions{
		FileURL: "foo.com",
	})
	assert.Error(t, err)
	assert.Nil(t, resp)
}

func TestIntegrationFailInvalidClientID(t *testing.T) {
	client, err := NewClientV7(&Options{
		ClientID:     "foo",
		ClientSecret: "bar",
		Username:     os.Getenv("USERNAME"),
		APIKey:       os.Getenv("API_KEY"),
	})
	assert.NotNil(t, client)
	assert.NoError(t, err)

	resp, err := client.ProcessDocumentURL(scheme.DocumentURLOptions{
		FileURL: "foo.com",
	})
	assert.Error(t, err)
	assert.Nil(t, resp)
}

func TestIntegrationFailInvalidUsername(t *testing.T) {
	client, err := NewClientV7(&Options{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Username:     "foo",
		APIKey:       os.Getenv("API_KEY"),
	})
	assert.NoError(t, err)
	assert.NotNil(t, client)

	resp, err := client.ProcessDocumentURL(scheme.DocumentURLOptions{
		FileURL: "foo.com",
	})
	assert.Error(t, err)
	assert.Nil(t, resp)
}

func TestIntegrationFailInvalidAPIKey(t *testing.T) {
	client, err := NewClientV7(&Options{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Username:     os.Getenv("USERNAME"),
		APIKey:       "foo",
	})
	assert.NoError(t, err)
	assert.NotNil(t, client)

	resp, err := client.ProcessDocumentURL(scheme.DocumentURLOptions{
		FileURL: "foo.com",
	})
	assert.Error(t, err)
	assert.Nil(t, resp)
}

func TestIntegrationFailInvalidDocument(t *testing.T) {
	client, err := NewClientV7(&Options{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Username:     os.Getenv("USERNAME"),
		APIKey:       os.Getenv("API_KEY"),
	})
	assert.NotNil(t, client)
	assert.NoError(t, err)

	resp, err := client.ProcessDocumentURL(scheme.DocumentURLOptions{
		FileURL: "foo.com",
	})
	assert.Error(t, err)
	assert.Nil(t, resp)
}

func TestIntegrationSuccessProcessDocumentURL(t *testing.T) {
	client, err := NewClientV7(&Options{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Username:     os.Getenv("USERNAME"),
		APIKey:       os.Getenv("API_KEY"),
	})
	assert.NotNil(t, client)
	assert.NoError(t, err)

	resp, err := client.ProcessDocumentURL(scheme.DocumentURLOptions{
		FileURL: "http://cdn-dev.veryfi.com/testing/veryfi-python/receipt_public.jpg",
		DocumentSharedOptions: scheme.DocumentSharedOptions{
			Tags: []string{"integration", "test", "url"},
		},
	})
	defer cleanUp(t, resp.ID)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestIntegrationSuccessProcessDocumentUpload(t *testing.T) {
	client, err := NewClientV7(&Options{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Username:     os.Getenv("USERNAME"),
		APIKey:       os.Getenv("API_KEY"),
	})
	assert.NotNil(t, client)
	assert.NoError(t, err)

	testpath, err := os.Getwd()
	assert.NoError(t, err)

	resp, err := client.ProcessDocumentUpload(scheme.DocumentUploadOptions{
		FilePath: fmt.Sprintf("%s/testdata/receipt_public.jpeg", testpath),
		DocumentSharedOptions: scheme.DocumentSharedOptions{
			FileName: "receipt_public.jpeg",
			Tags:     []string{"integration", "test", "upload"},
		},
	})
	defer cleanUp(t, resp.ID)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}
