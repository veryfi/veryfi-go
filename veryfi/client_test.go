package veryfi

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/veryfi/veryfi-go/veryfi/scheme"
	"github.com/veryfi/veryfi-go/veryfi/test"
)

func setUp(t *testing.T) (test.HTTPServer, *Client, string, *scheme.Document) {
	server := test.NewHTTPServer()
	assert.NotNil(t, server)

	pwd, err := os.Getwd()
	assert.NoError(t, err)

	mockReceiptPath := fmt.Sprintf("%v/testdata/%v", pwd, "receipt_public.jpeg")
	mockReceiptData := fmt.Sprintf("%v/testdata/%v", pwd, "receipt_public.json")
	mockResp, err := ioutil.ReadFile(mockReceiptData)
	assert.NoError(t, err)

	mockRespStr := string(mockResp)
	server.Serve(t, "/api/v8/partner/documents/36966934/", 200, mockRespStr)
	server.Serve(t, "/api/v8/partner/documents/", 200, mockRespStr)

	expected := &scheme.Document{
		Created:       "2021-06-22 20:11:10",
		CurrencyCode:  "USD",
		Date:          "2021-06-22 16:11:10",
		ID:            36966934,
		ImgFileName:   "7a0371f1-f695-4f9b-9e2b-da54cdf189fc.jpg",
		InvoiceNumber: "98",
		LineItems: []scheme.LineItem{
			{
				Description: "98 Meat Pty Xchz",
				ID:          67185481,
				Price:       0.0,
				Quantity:    1.0,
				Total:       90.85,
				Type:        "food",
			},
		},
		OCRText: "\n\\x0c2004-10-31\n\t8:21 PM\nYOUR GUEST NUMBER IS\n98\nIN-N-OUT BURGER LAS VEGAS EASTERN\n2004-10-31\t\t8:21 PM\n165 1 5 98\nCashier: SAM\nGUEST #: 98\nCounter-Eat in\n\t2.65\nDbDb\t\t88.20\n98 Meat Pty Xchz\n\t90.85\nCounter-Eat In\t\t6.81\nTAX 7.50%\t\t97.66\nAmount Due\n\t$97.66\nCASH TENDER\t\t$.00\nChange\n2004-10-31\t\t8:21 PM\nTHANK YOU!\n",
		Payment: scheme.PaymentsInfo{
			CardNumber:  "1234",
			DisplayName: "Cash",
			Terms:       "",
			Type:        "cash",
		},
		ReferenceNumber: "VBIJG-6934",
		Tax:             97.66,
		TaxLines: []scheme.TaxLine{
			{
				Rate:  7.5,
				Total: 97.66,
			},
		},
		Total:   97.66,
		Updated: "2021-06-22 20:11:11",
		Vendor: scheme.Vendor{
			Name:    "In-N-Out Burger",
			RawName: "In-N-Out Burger",
			Logo:    "",
			Type:    "",
		},
	}

	client, err := NewClientV8(&Options{
		EnvironmentURL: server.URL,
		ClientID:       "testClientID",
		Username:       "testUsername",
		APIKey:         "testAPIKey",
	})
	client.SetTLSConfig(&tls.Config{InsecureSkipVerify: true})
	assert.NotNil(t, client)
	assert.NoError(t, err)

	return server, client, mockReceiptPath, expected
}

func TestUnitNewClientV8_NilConfig(t *testing.T) {
	client, err := NewClientV8(nil)
	assert.Nil(t, client)
	assert.Error(t, err)
}

func TestUnitNewClientV8_Config(t *testing.T) {
	client, err := NewClientV8(&Options{
		ClientID: "testClientID",
		Username: "testUsername",
		APIKey:   "testAPIKey",
	})
	assert.NotNil(t, client)
	assert.NoError(t, err)

	timeout, err := time.ParseDuration("120s")
	assert.NoError(t, err)

	waitTime, err := time.ParseDuration("100ms")
	assert.NoError(t, err)

	maxWaitTime, err := time.ParseDuration("360s")
	assert.NoError(t, err)

	expected := &Options{
		EnvironmentURL: "api.veryfi.com",
		ClientID:       "testClientID",
		Username:       "testUsername",
		APIKey:         "testAPIKey",
		HTTP: HTTPOptions{
			Timeout: timeout,
			Retry: RetryOptions{
				Count:       3,
				WaitTime:    waitTime,
				MaxWaitTime: maxWaitTime,
			},
		},
	}

	resp := client.Config()
	assert.NotNil(t, resp)
	assert.Equal(t, expected, resp)
}

func TestUnitClientV8_GetDocument(t *testing.T) {
	server, client, _, expected := setUp(t)
	defer server.Close()

	resp, err := client.GetDocument("36966934", scheme.DocumentGetOptions{})
	assert.NotNil(t, resp)
	assert.NoError(t, err)
	assert.Equal(t, expected, resp)
}

func TestUnitClientV8_ProcessDocumentUpload(t *testing.T) {
	server, client, mockReceiptPath, expected := setUp(t)
	defer server.Close()

	resp, err := client.ProcessDocumentUpload(scheme.DocumentUploadOptions{
		FilePath: mockReceiptPath,
	})
	assert.NotNil(t, resp)
	assert.NoError(t, err)
	assert.Equal(t, expected, resp)
}

func TestUnitClientV8_ProcessDocumentURL(t *testing.T) {
	server, client, _, expected := setUp(t)
	defer server.Close()

	resp, err := client.ProcessDocumentURL(scheme.DocumentURLOptions{
		FileURL: "http://cdn-dev.veryfi.com/testing/veryfi-python/receipt_public.jpg",
	})
	assert.NotNil(t, resp)
	assert.NoError(t, err)
	assert.Equal(t, expected, resp)
}
