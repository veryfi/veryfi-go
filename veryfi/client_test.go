package veryfi

import (
	"crypto/tls"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/veryfi/veryfi-go/veryfi/scheme"
	"github.com/veryfi/veryfi-go/veryfi/test"
)

func TestUnitNewClientV7_NilConfig(t *testing.T) {
	client, err := NewClientV7(nil)
	assert.Nil(t, client)
	assert.Error(t, err)
}

func TestUnitClientV7_DocumentProcessing(t *testing.T) {
	in := `
{
   "abn_number":"",
   "account_number":"",
   "barcodes":[],
   "bill_to_address":"",
   "bill_to_name":"",
   "bill_to_vat_number":"",
   "card_number":"",
   "cashback":0.0,
   "category":"",
   "created":"2021-06-22 20:11:10",
   "currency_code":"USD",
   "date":"2021-06-22 16:11:10",
   "discount":0.0,
   "due_date":"",
   "external_id":"",
   "id":36966934,
   "img_file_name":"7a0371f1-f695-4f9b-9e2b-da54cdf189fc.jpg",
   "img_thumbnail_url":"",
   "img_url":"",
   "invoice_number":"98",
   "line_items":[
      {
         "date":"",
         "description":"98 Meat Pty Xchz",
         "discount":0.0,
         "id":67185481,
         "order":0,
         "price":0.0,
         "quantity":1.0,
         "reference":"",
         "sku":"",
         "tax":0.0,
         "tax_rate":0.0,
         "total":90.85,
         "type":"food",
         "unit_of_measure":""
      }
   ],
   "notes":"",
   "ocr_text":"\n\\x0c2004-10-31\n\t8:21 PM\nYOUR GUEST NUMBER IS\n98\nIN-N-OUT BURGER LAS VEGAS EASTERN\n2004-10-31\t\t8:21 PM\n165 1 5 98\nCashier: SAM\nGUEST #: 98\nCounter-Eat in\n\t2.65\nDbDb\t\t88.20\n98 Meat Pty Xchz\n\t90.85\nCounter-Eat In\t\t6.81\nTAX 7.50%\t\t97.66\nAmount Due\n\t$97.66\nCASH TENDER\t\t$.00\nChange\n2004-10-31\t\t8:21 PM\nTHANK YOU!\n",
   "payment_display_name":"Cash",
   "payment_terms":"",
   "payment_type":"cash",
   "phone_number":"",
   "purchase_order_number":"",
   "reference_number":"VBIJG-6934",
   "rounding":0.0,
   "service_end_date":"",
   "service_start_date":"",
   "shipping":0.0,
   "subtotal":0.0,
   "tags":[],
   "tax":97.66,
   "tax_lines":[
      {
         "base":0.0,
         "name":"",
         "order":0,
         "rate":7.5,
         "total":97.66
      }
   ],
   "tip":0.0,
   "total":97.66,
   "tracking_number":"",
   "updated":"2021-06-22 20:11:11",
   "vat_number":"",
   "vendor":{
      "address":"",
      "email":"",
      "fax_number":"",
      "name":"In-N-Out Burger",
      "phone_number":"",
      "raw_name":"In-N-Out Burger",
      "vendor_logo":"https://cdn.veryfi.com/logos/us/949103001.png",
      "vendor_reg_number":"",
      "vendor_type":"Restaurant",
      "web":""
   },
   "vendor_account_number":"",
   "vendor_bank_name":"",
   "vendor_bank_number":"",
   "vendor_bank_swift":"",
   "vendor_iban":""
}`

	expected := &scheme.Document{
		Created:       "2021-06-22 20:11:10",
		CurrencyCode:  "USD",
		Date:          "2021-06-22 16:11:10",
		ID:            36966934,
		ImgFileName:   "7a0371f1-f695-4f9b-9e2b-da54cdf189fc.jpg",
		InvoiceNumber: "98",
		LineItems: []scheme.LineItem{
			scheme.LineItem{
				Description: "98 Meat Pty Xchz",
				ID:          67185481,
				Price:       0.0,
				Quantity:    1.0,
				Total:       90.85,
				Type:        "food",
			},
		},
		OCRText:            "\n\\x0c2004-10-31\n\t8:21 PM\nYOUR GUEST NUMBER IS\n98\nIN-N-OUT BURGER LAS VEGAS EASTERN\n2004-10-31\t\t8:21 PM\n165 1 5 98\nCashier: SAM\nGUEST #: 98\nCounter-Eat in\n\t2.65\nDbDb\t\t88.20\n98 Meat Pty Xchz\n\t90.85\nCounter-Eat In\t\t6.81\nTAX 7.50%\t\t97.66\nAmount Due\n\t$97.66\nCASH TENDER\t\t$.00\nChange\n2004-10-31\t\t8:21 PM\nTHANK YOU!\n",
		PaymentDisplayName: "Cash",
		PaymentType:        "cash",
		ReferenceNumber:    "VBIJG-6934",
		Tax:                97.66,
		TaxLines: []scheme.TaxLine{
			scheme.TaxLine{
				Rate:  7.5,
				Total: 97.66,
			},
		},
		Total:   97.66,
		Updated: "2021-06-22 20:11:11",
		Vendor: scheme.Vendor{
			Name:       "In-N-Out Burger",
			RawName:    "In-N-Out Burger",
			VendorLogo: "https://cdn.veryfi.com/logos/us/949103001.png",
			VendorType: "Restaurant",
		},
	}

	server := test.NewHTTPServer()
	defer server.Close()

	server.Serve(t, "/api/v7/partner/documents/36966934", 200, in)

	client, err := NewClientV7(&Options{
		EnvironmentURL: server.URL,
		ClientID:       "testClientID",
		Username:       "testUsername",
		APIKey:         "testAPIKey",
	})
	client.SetTLSConfig(&tls.Config{InsecureSkipVerify: true})
	assert.NotNil(t, client)
	assert.NoError(t, err)

	resp, err := client.GetDocument("36966934", scheme.DocumentGetOptions{})
	assert.NotNil(t, resp)
	assert.NoError(t, err)
	assert.Equal(t, expected, resp)
}
