package veryfi

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/veryfi/veryfi-go/v3/veryfi/scheme"
)

// newCapturingClient returns a client whose request bodies are recorded in sent.
func newCapturingClient(t *testing.T, sent *[]map[string]interface{}) (*Client, func()) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		assert.NoError(t, err)

		var got map[string]interface{}
		assert.NoError(t, json.Unmarshal(body, &got))
		*sent = append(*sent, got)

		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write([]byte(`{"id": 1}`))
		assert.NoError(t, err)
	}))

	client, err := NewClientV8(&Options{
		EnvironmentURL: strings.TrimPrefix(server.URL, "https://"),
		ClientID:       "testClientID",
		Username:       "testUsername",
		APIKey:         "testAPIKey",
	})
	assert.NoError(t, err)
	client.SetTLSConfig(&tls.Config{InsecureSkipVerify: true})

	return client, server.Close
}

func TestUnitClientV8_ProcessDocumentUpload_Thinking(t *testing.T) {
	pwd, err := os.Getwd()
	assert.NoError(t, err)
	receiptPath := fmt.Sprintf("%v/testdata/%v", pwd, "receipt_public.jpg")

	for _, mode := range []scheme.ThinkingMode{scheme.ThinkingLight, scheme.ThinkingDeep} {
		t.Run(string(mode), func(t *testing.T) {
			var sent []map[string]interface{}
			client, closeServer := newCapturingClient(t, &sent)
			defer closeServer()

			_, err := client.ProcessDocumentUpload(scheme.DocumentUploadOptions{
				FilePath:              receiptPath,
				DocumentSharedOptions: scheme.DocumentSharedOptions{Thinking: mode},
			})
			assert.NoError(t, err)
			assert.Len(t, sent, 1)
			assert.Equal(t, string(mode), sent[0]["thinking"])
		})
	}
}

func TestUnitClientV8_ProcessDocumentUpload_NoThinking(t *testing.T) {
	pwd, err := os.Getwd()
	assert.NoError(t, err)
	receiptPath := fmt.Sprintf("%v/testdata/%v", pwd, "receipt_public.jpg")

	var sent []map[string]interface{}
	client, closeServer := newCapturingClient(t, &sent)
	defer closeServer()

	_, err = client.ProcessDocumentUpload(scheme.DocumentUploadOptions{FilePath: receiptPath})
	assert.NoError(t, err)
	assert.Len(t, sent, 1)
	assert.NotContains(t, sent[0], "thinking")
}

func TestUnitClientV8_ProcessDocumentURL_Thinking(t *testing.T) {
	for _, mode := range []scheme.ThinkingMode{scheme.ThinkingLight, scheme.ThinkingDeep} {
		t.Run(string(mode), func(t *testing.T) {
			var sent []map[string]interface{}
			client, closeServer := newCapturingClient(t, &sent)
			defer closeServer()

			_, err := client.ProcessDocumentURL(scheme.DocumentURLOptions{
				FileURL:               "http://cdn-dev.veryfi.com/testing/veryfi-python/receipt_public.jpg",
				DocumentSharedOptions: scheme.DocumentSharedOptions{Thinking: mode},
			})
			assert.NoError(t, err)
			assert.Len(t, sent, 1)
			assert.Equal(t, string(mode), sent[0]["thinking"])
		})
	}
}

func TestUnitClientV8_ProcessDocumentURL_NoThinking(t *testing.T) {
	var sent []map[string]interface{}
	client, closeServer := newCapturingClient(t, &sent)
	defer closeServer()

	_, err := client.ProcessDocumentURL(scheme.DocumentURLOptions{
		FileURL: "http://cdn-dev.veryfi.com/testing/veryfi-python/receipt_public.jpg",
	})
	assert.NoError(t, err)
	assert.Len(t, sent, 1)
	assert.NotContains(t, sent[0], "thinking")
}
