package veryfi

import (
	"crypto/tls"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/veryfi/veryfi-go/veryfi/scheme"
	"github.com/veryfi/veryfi-go/veryfi/test"
)

func setUp(t *testing.T, useDetailedReceipt bool) (test.HTTPServer, *Client, string, interface{}) {
	server := test.NewHTTPServer()
	assert.NotNil(t, server)

	pwd, err := os.Getwd()
	assert.NoError(t, err)

	mockReceiptPath := fmt.Sprintf("%v/testdata/%v", pwd, "receipt_public.jpg")
	
	var mockReceiptData string
	var expected interface{}
	
	if useDetailedReceipt {
		mockReceiptData = fmt.Sprintf("%v/testdata/%v", pwd, "detailed_receipt_public.json")
		mockResp, err := os.ReadFile(mockReceiptData)
		assert.NoError(t, err)

		mockRespStr := string(mockResp)
		server.Serve(t, "/api/v8/partner/documents/36966934/", 200, mockRespStr)
		server.Serve(t, "/api/v8/partner/documents/", 200, mockRespStr)
		expected = &scheme.DetailedDocument{
			ID:              36966934,
			ReferenceNumber: "VFFFF-10000",
			ImgFileName:     "36966934.jpg",
			ImgThumbnailURL: "https://scdn.veryfi.com/receipts/thumbnail.jpg",
			ImgURL:          "https://scdn.veryfi.com/receipts/img.jpg",
			IsDuplicate:     true,
			OCRText:         "Walgreens\n#03296 191 E 3RD AVE\nSAN MATEO, CA 94401\n\t650-342-2723\n117\t4782 0022 05/24/2022 1:10 PM\nRED BULL ENRGY DRNK CNS 8.4OZ 6PK\n61126943157\tA\t8.79 SALE\nREGULAR PRICE 9.99\nMYWALGREENS SAVINGS 1.20\nRETURN VALUE 8.79\nCA REDMP VAL\n00000007211\t\t0.30\nCOCA COLA MINICAN 7.5Z 6PK\n04900006101\tA\t4.99 SALE\nRETURN VALUE 4.99\nCA REDMP VAL\n00000007211\t\t0.30\nNAB OREO CKIES C/PK 5.25OZ WHSE\n04400000749\t\t2.69\nRETURN VALUE 2.69\nDORITOS NACHO\n02840032505\t\t2.00\n1 @ 2.19 or 2/4.00\nRETURN VALUE 2.00\nF/LAYS REGULAR 2.63OZ\n02840032382\t\t2.00\n1 @ 2.19 or 2/4.00\nRETURN VALUE 2.00\nSCOTCH BRITE H/D KITCHN SPONGE 3S\n02120057235\tA\t4.79\nRETURN VALUE 4.79\nPALMOLIVE DISH OXI POWER\t10OZ\n03500000168\tA\t1.49\nRETURN VALUE 1.49\nSHOPPING BAG FEE\t0.25\nSUBTOTAL\t\t27.60\nSALES TAX A=9.625%\t1.93\nTOTAL\t\t29.53\nVISA ACCT 1850\t29.53\nAUTH CODE\t\t798553\nCHANGE\t\t.00\n\nMYWALGREENS SAVINGS\tof\t1.20\nTHANK YOU FOR SHOPPING AT WALGREENS\nREDEEM $1 WALGREENS CASH REWARDS ON YOUR\nNEXT PURCHASE! WALGREENS CASH REWARDS\nCANNOT BE REDEEMED ON SOME ITEMS. FOR\nFULL DETAILS SEE MYWALGREENS.COM\nRFN# 0329-6224-7823-2205-2403\n\n*****\nmyW\nTOTAL SAVINGS\t\t$1.20\nSAVINGS VALUE\t\t4%\n$1.40 W CASH REWARDS AVAILABLE\nmyWalgreens ACCT # *********0053\n\t008\nOPENING BALANCE\t\t$1.14\nEARNED THIS VISIT\t\t$0.26\nCLOSING BALANCE\t\t$1.40",
			Status:          "processed",
			
			// Detailed fields with scores
			AccountNumber: &scheme.DetailedField{
				Value:    "0053",
				Score:    1.0,
				OCRScore: 0.93,
				BoundingBox: []float64{0, 0.4543, 0.9355, 0.7466, 0.9512},
				BoundingRegion: []float64{0.4543, 0.9355, 0.7466, 0.9355, 0.7466, 0.9512, 0.4543, 0.9512},
				Rotation: 0,
			},
			Category: &scheme.DetailedField{
				Value: "Job Supplies",
				Score: 0.94,
			},
			CurrencyCode: &scheme.DetailedField{
				Value: "USD",
				Score: 0.96,
			},
			Date: &scheme.DetailedDateField{
				Value:    "2022-05-24 13:10:00",
				Score:    1.0,
				BoundingBox: []float64{0, 0.5332, 0.1168, 0.7573, 0.131},
				BoundingRegion: []float64{0.5332, 0.1168, 0.7573, 0.1168, 0.7573, 0.131, 0.5332, 0.131},
				Rotation: 0,
			},
			DeliveryDate: &scheme.DetailedDateField{
				Value:    "",
				Score:    1.0,
				OCRScore: 0.0,
				BoundingBox: nil,
				BoundingRegion: nil,
				Rotation: 0,
			},
			Discount: &scheme.DetailedFloatField{
				Value:    1.2,
				Score:    0.99,
				OCRScore: 0.98,
				BoundingBox: []float64{0, 0.7646, 0.8789, 0.8823, 0.8906},
				BoundingRegion: []float64{0.7646, 0.8789, 0.8823, 0.8789, 0.8823, 0.8906, 0.7646, 0.8906},
				Rotation: 0,
			},
			DueDate: &scheme.DetailedDateField{
				Value:    "",
				Score:    1.0,
				OCRScore: 0.0,
				BoundingBox: nil,
				BoundingRegion: nil,
				Rotation: 0,
			},
			Insurance: &scheme.DetailedFloatField{
				Value:    0,
				Score:    1.0,
				OCRScore: 0.0,
				BoundingBox: nil,
				BoundingRegion: nil,
				Rotation: 0,
			},
			InvoiceNumber: &scheme.DetailedField{
				Value:    "4782",
				Score:    1.0,
				OCRScore: 0.99,
				BoundingBox: []float64{0, 0.2482, 0.1201, 0.3438, 0.1334},
				BoundingRegion: []float64{0.2482, 0.1201, 0.3437, 0.1201, 0.3437, 0.1334, 0.2482, 0.1334},
				Rotation: 0,
			},
			OrderDate: &scheme.DetailedDateField{
				Value:          "",
				Score:          1,
				OCRScore:       0,
				BoundingBox:    nil,
				BoundingRegion: nil,
				Rotation:       0,
			},
			PurchaseOrderNumber: &scheme.DetailedField{
				Value:    "",
				Score:    0.99,
				OCRScore: 0,
				BoundingBox: nil,
				BoundingRegion: nil,
				Rotation: 0,
			},
			Rounding: &scheme.DetailedFloatField{
				Value:    0,
				Score:    0.99,
				OCRScore: 0,
				BoundingBox: nil,
				BoundingRegion: nil,
				Rotation: 0,
			},
			ServiceEndDate: &scheme.DetailedDateField{
				Value:    "",
				Score:    1,
				OCRScore: 0,
				BoundingBox: nil,
				BoundingRegion: nil,
				Rotation: 0,
			},
			ServiceStartDate: &scheme.DetailedDateField{
				Value:    "",
				Score:    1,
				OCRScore: 0,
				BoundingBox: nil,
				BoundingRegion: nil,
				Rotation: 0,
			},
			ShipDate: &scheme.DetailedDateField{
				Value:    "",
				Score:    1,
				OCRScore: 0,
				BoundingBox: nil,
				BoundingRegion: nil,
				Rotation: 0,
			},
			StoreNumber: &scheme.DetailedField{
				Value:    "03296",
				Score:    1.0,
				OCRScore: 0.96,
				BoundingBox: []float64{
					0, 0.2744, 0.0699, 0.408, 0.0821,
				},
				BoundingRegion: []float64{
					0.2744, 0.0699, 0.408, 0.0699, 0.408, 0.0821, 0.2744, 0.0821,
				},
				Rotation: 0,
			},
			Subtotal: &scheme.DetailedFloatField{
				Value:    27.6,
				Score:    1.0,
				OCRScore: 0.99,
				BoundingBox: []float64{
					0, 0.71, 0.5151, 0.8232, 0.5273,
				},
				BoundingRegion: []float64{
					0.71, 0.5151, 0.8232, 0.5151, 0.8232, 0.5273, 0.71, 0.5273,
				},
				Rotation: 0,
			},
			Tax: &scheme.DetailedFloatField{
				Value:    1.93,
				Score:    1.0,
				OCRScore: 0.99,
				BoundingBox: []float64{
					0, 0.7349, 0.5278, 0.8223, 0.5386,
				},
				BoundingRegion: []float64{
					0.7349, 0.5278, 0.8223, 0.5278, 0.8223, 0.5386, 0.7349, 0.5386,
				},
				Rotation: 0,
			},
			Tip: &scheme.DetailedFloatField{
				Value:    0,
				Score:    1,
				OCRScore: 0,
				BoundingBox: nil,
				BoundingRegion: nil,
				Rotation: 0,
			},
			Total: &scheme.DetailedFloatField{
				Value:    29.53,
				Score:    1.0,
				OCRScore: 0.99,
				BoundingBox: []float64{
					0, 0.71, 0.5518, 0.8203, 0.562,
				},
				BoundingRegion: []float64{
					0.71, 0.5518, 0.8203, 0.5518, 0.8203, 0.562, 0.71, 0.562,
				},
				Rotation: 0,
			},
			TotalWeight: &scheme.DetailedField{
				Value:    "",
				Score:    0.99,
				OCRScore: 0,
				BoundingBox: nil,
				BoundingRegion: nil,
				Rotation: 0,
			},
			TrackingNumber: &scheme.DetailedField{
				Value:    "",
				Score:    1,
				OCRScore: 0,
				BoundingBox: nil,
				BoundingRegion: nil,
				Rotation: 0,
			},

			// Payment information
			Payment: &scheme.DetailedPayment{
				CardNumber: &scheme.DetailedField{
					Value:    "1850",
					Score:    1.0,
					OCRScore: 0.99,
					BoundingBox: []float64{0, 0.345, 0.5664, 0.4333, 0.5786},
					BoundingRegion: []float64{0.345, 0.5664, 0.4333, 0.5664, 0.4333, 0.5786, 0.345, 0.5786},
					Rotation: 0,
				},
				Type: &scheme.DetailedField{
					Value: "visa",
					Score: 0.97,
				},
				DisplayName: "Visa ***1850",
				Terms: &scheme.DetailedField{
					Value: "",
					Score: 1.0,
					OCRScore: 0.0,
					Rotation: 0,
				},
			},

			// Vendor information
			Vendor: &scheme.DetailedVendor{
				ABNNumber: &scheme.DetailedField{
					Value:    "",
					Score:    1,
					OCRScore: 0,
					BoundingBox: nil,
					BoundingRegion: nil,
					Rotation: 0,
				},
				AccountCurrency: &scheme.DetailedField{
					Value:    "",
					Score:    1,
					OCRScore: 0,
					BoundingBox: nil,
					BoundingRegion: nil,
					Rotation: 0,
				},
				AccountNumber: &scheme.DetailedField{
					Value:    "",
					Score:    1,
					OCRScore: 0,
					BoundingBox: nil,
					BoundingRegion: nil,
					Rotation: 0,
				},
				BankName: &scheme.DetailedField{
					Value:    "",
					Score:    1,
					OCRScore: 0,
					BoundingBox: nil,
					BoundingRegion: nil,
					Rotation: 0,
				},
				BankNumber: &scheme.DetailedField{
					Value:    "",
					Score:    1,
					OCRScore: 0,
					BoundingBox: nil,
					BoundingRegion: nil,
					Rotation: 0,
				},
				BankSwift: &scheme.DetailedField{
					Value:    "",
					Score:    1,
					OCRScore: 0,
					BoundingBox: nil,
					BoundingRegion: nil,
					Rotation: 0,
				},
				Name: &scheme.DetailedField{
					Value: "Walgreens",
					Score: 0.98,
				},
				IBAN: &scheme.DetailedField{
					Value:    "",
					Score:    1,
					OCRScore: 0,
					BoundingBox: nil,
					BoundingRegion: nil,
					Rotation: 0,
				},
				PhoneNumber: &scheme.DetailedField{
					Value:    "650-342-2723",
					Score:    1.0,
					OCRScore: 0.99,
					BoundingBox: []float64{0, 0.3628, 0.0937, 0.6265, 0.1068},
					BoundingRegion: []float64{0.3628, 0.0937, 0.6265, 0.0937, 0.6265, 0.1068, 0.3628, 0.1068},
					Rotation: 0,
				},
				Address: &scheme.DetailedField{
					Value: "191 E 3Rd Ave, San Mateo, CA 94401",
				},
				RawName: &scheme.DetailedField{
					Value:    "Walgreens",
					Score:    1.0,
					OCRScore: 0.98,
					BoundingBox: []float64{0, 0.0337, 0.0187, 0.8896, 0.0629},
					BoundingRegion: []float64{0.0337, 0.0187, 0.8896, 0.0187, 0.8896, 0.0629, 0.0337, 0.0629},
					Rotation: 0,
				},
				Type: &scheme.DetailedField{
					Value: "drugstores, convenience stores, cosmetics & beauty supply",
				},
				Logo: "https://cdn.veryfi.com/logos/us/126568182.jpeg",
				Lat:  37.564947,
				Lng:  -122.3234454,
				ParsedAddress: &scheme.ParsedAddress{
					City: "",
					Country: "",
					Postcode: "",
					State: "",
					StreetAddress: "",
					House: "",
					HouseNumber: "",
					Road: "",
					Unit: "",
					Level: "",
					Staircase: "",
					POBox: "",
					Suburb: "",
					CityDistrict: "",
					Island: "",
					StateDistrict: "",
					CountryRegion: "",
					WorldRegion: "",
				},
				Email: &scheme.DetailedField{
					Value:    "",
					Score:    1.0,
					OCRScore: 0.0,
					BoundingBox: nil,
					BoundingRegion: nil,
					Rotation: 0,
				},
				VATNumber: &scheme.DetailedField{
					Value:    "",
					Score:    1.0,
					OCRScore: 0.0,
					BoundingBox: nil,
					BoundingRegion: nil,
					Rotation: 0,
				},
				RegNumber: &scheme.DetailedField{
					Value:    "",
					Score:    1.0,
					OCRScore: 0.0,
					BoundingBox: nil,
					BoundingRegion: nil,
					Rotation: 0,
				},
				Web: &scheme.DetailedField{
					Value:    "",
					Score:    0.4,
					OCRScore: 0.0,
					BoundingBox: nil,
					BoundingRegion: nil,
					Rotation: 0,
				},
				FaxNumber: &scheme.DetailedField{
					Value:    "",
					Score:    1.0,
					OCRScore: 0.0,
					BoundingBox: nil,
					BoundingRegion: nil,
					Rotation: 0,
				},
			},

			// Line items
			LineItems: []scheme.LineItem{
				{
					Description: "RED BULL ENRGY DRNK CNS 8.4OZ 6PK",
					ID:          1346628550,
					Order:       0,
					Price:       9.99,
					Quantity:    1.0,
					Discount:    1.2,
					Total:       8.79,
					SKU:         "61126943157",
					Type:        "food",
				},
				{
					Description: "CA REDMP VAL",
					ID:          1346628551,
					Order:       1,
					Quantity:    1.0,
					Total:       0.3,
					SKU:         "00000007211",
					Type:        "fee",
				},
			},

			// Line items with scores
			LineItemsWithScores: []scheme.DetailedLineItem{
				{
					Date: &scheme.DetailedDateField{
						Value:          "",
						Score:          0,
						OCRScore:       0,
						BoundingBox:    nil,
						BoundingRegion: nil,
						Rotation:       0,
					},
					Description: &scheme.DetailedField{
						Value:    "RED BULL ENRGY DRNK CNS 8.4OZ 6PK",
						Score:    0.78,
						OCRScore: 0.98,
						Rotation: 0,
						BoundingBox: []float64{
							0, 0.0326, 0.1411, 0.7573, 0.1592,
						},
						BoundingRegion: []float64{
							0.0325, 0.146, 0.7573, 0.1401, 0.7574, 0.1541, 0.0326, 0.16,
						},
					},
					Discount: &scheme.DetailedFloatField{
						Value:    1.2,
						Score:    1.0,
						OCRScore: 0.99,
						Rotation: 0,
						BoundingBox: []float64{
							0, 0.5615, 0.1785, 0.6489, 0.1909,
						},
						BoundingRegion: []float64{
							0.5615, 0.1785, 0.6489, 0.1785, 0.6489, 0.1909, 0.5615, 0.1909,
						},
					},
					Price: &scheme.DetailedFloatField{
						Value:    9.99,
						Score:    0.97,
						OCRScore: 0.99,
						Rotation: 0,
						BoundingBox: []float64{
							0, 0.4258, 0.1681, 0.5195, 0.1803,
						},
						BoundingRegion: []float64{
							0.4258, 0.1681, 0.5195, 0.1681, 0.5195, 0.1803, 0.4258, 0.1803,
						},
					},
					Quantity: &scheme.DetailedFloatField{
						Value:          1,
						Score:          0,
						OCRScore:       0,
						BoundingBox:    nil,
						BoundingRegion: nil,
						Rotation:       0,
					},
					Reference: &scheme.DetailedField{
						Value:          "",
						Score:          0,
						OCRScore:       0,
						BoundingBox:    nil,
						BoundingRegion: nil,
						Rotation:       0,
					},
					Section: &scheme.DetailedField{
						Value:          "",
						Score:          0,
						OCRScore:       0,
						BoundingBox:    nil,
						BoundingRegion: nil,
						Rotation:       0,
					},
					Total: &scheme.DetailedFloatField{
						Value:    8.79,
						Score:    0.95,
						OCRScore: 0.99,
						Rotation: 0,
						BoundingBox: []float64{
							0, 0.7319, 0.1523, 0.8223, 0.1646,
						},
						BoundingRegion: []float64{
							0.7319, 0.1523, 0.8223, 0.1523, 0.8223, 0.1646, 0.7319, 0.1646,
						},
					},
					SKU: &scheme.DetailedField{
						Value:    "61126943157",
						Score:    0.94,
						OCRScore: 0.99,
						Rotation: 0,
						BoundingBox: []float64{
							0, 0.1209, 0.1571, 0.365, 0.1699,
						},
						BoundingRegion: []float64{
							0.1209, 0.1571, 0.365, 0.1571, 0.365, 0.1699, 0.1209, 0.1699,
						},
					},
					Tax: &scheme.DetailedFloatField{
						Value:          0,
						Score:          0,
						OCRScore:       0,
						BoundingBox:    nil,
						BoundingRegion: nil,
						Rotation:       0,
					},
					TaxRate: &scheme.DetailedFloatField{
						Value:          0,
						Score:          0,
						OCRScore:       0,
						BoundingBox:    nil,
						BoundingRegion: nil,
						Rotation:       0,
					},
					UnitOfMeasure: &scheme.DetailedField{
						Value:    "",
						Score:    0,
						OCRScore: 0,
						BoundingBox: nil,
						BoundingRegion: nil,
						Rotation: 0,
					},
					UPC: &scheme.DetailedField{
						Value:          "",
						Score:          0,
						OCRScore:       0,
						BoundingBox:    nil,
						BoundingRegion: nil,
						Rotation:       0,
					},
					ID:    1346628550,
					Order: 0,
					Type:  "food",
				},
				{
					Date: &scheme.DetailedDateField{
						Value:          "",
						Score:          0,
						OCRScore:       0,
						BoundingBox:    nil,
						BoundingRegion: nil,
						Rotation:       0,
					},
					Description: &scheme.DetailedField{
						Value:    "CA REDMP VAL",
						Score:    0.79,
						OCRScore: 0.98,
						Rotation: 0,
						BoundingBox: []float64{
							0, 0.0337, 0.2073, 0.3008, 0.2206,
						},
						BoundingRegion: []float64{
							0.0336, 0.2082, 0.3008, 0.2069, 0.3008, 0.2196, 0.0337, 0.2208,
						},
					},
					Discount: &scheme.DetailedFloatField{
						Value:    0,
						Score:    0,
						OCRScore: 0,
						BoundingBox: nil,
						BoundingRegion: nil,
						Rotation: 0,
					},
					Total: &scheme.DetailedFloatField{
						Value:    0.3,
						Score:    0.8,
						OCRScore: 0.99,
						BoundingBox: []float64{
							0, 0.7329, 0.2141, 0.8232, 0.2253,
						},
						BoundingRegion: []float64{
							0.7329, 0.2141, 0.8232, 0.2141, 0.8232, 0.2253, 0.7329, 0.2253,
						},
						Rotation: 0,
					},
					UnitOfMeasure: &scheme.DetailedField{
						Value:          "",
						Score:          0,
						OCRScore:       0,
						BoundingBox:    nil,
						BoundingRegion: nil,
						Rotation: 0,
					},
					UPC: &scheme.DetailedField{
						Value:          "",
						Score:          0,
						OCRScore:       0,
						BoundingBox:    nil,
						BoundingRegion: nil,
						Rotation: 0,
					},
					Tax: &scheme.DetailedFloatField{
						Value:          0,
						Score:          0,
						OCRScore:       0,
						BoundingBox:    nil,
						BoundingRegion: nil,
						Rotation: 0,
					},
					TaxRate: &scheme.DetailedFloatField{
						Value:          0,
						Score:          0,
						OCRScore:       0,
						BoundingBox:    nil,
						BoundingRegion: nil,
						Rotation: 0,
					},
					Price: &scheme.DetailedFloatField{
						Value:    0,
						Score:    0,
						OCRScore: 0,
						BoundingBox: nil,
						BoundingRegion: nil,
						Rotation: 0,
					},
					Quantity: &scheme.DetailedFloatField{
						Value:          1,
						Score:          0,
						OCRScore:       0,
						BoundingBox:    nil,
						BoundingRegion: nil,
						Rotation: 0,
					},
					Reference: &scheme.DetailedField{
						Value:          "",
						Score:          0,
						OCRScore:       0,
						BoundingBox:    nil,
						BoundingRegion: nil,
						Rotation: 0,
					},
					Section: &scheme.DetailedField{
						Value:          "",
						Score:          0,
						OCRScore:       0,
						BoundingBox:    nil,
						BoundingRegion: nil,
						Rotation: 0,
					},
					SKU: &scheme.DetailedField{
						Value:    "00000007211",
						Score:    0.93,
						OCRScore: 0.97,
						BoundingBox: []float64{
							0, 0.1209, 0.2185, 0.365, 0.2312,
						},
						BoundingRegion: []float64{
							0.1209, 0.2185, 0.365, 0.2185, 0.365, 0.2312, 0.1209, 0.2312,
						},
						Rotation: 0,
					},
					ID:    1346628551,
					Order: 1,
					Type:  "fee",
				},
			},

			// Tax lines
			TaxLines: []scheme.TaxLine{
				{
					Order: 0,
					Rate:  9.625,
					Total: 1.93,
				},
			},

			// Tax lines with scores
			TaxLinesWithScores: []scheme.DetailedTaxLine{
				{
					Order: 0,
					Name: &scheme.DetailedField{
						Value:    "",
						Score:    0,
						OCRScore: 0,
						BoundingBox: nil,
						BoundingRegion: nil,
						Rotation: 0,
					},
					Rate: &scheme.DetailedFloatField{
						Value:    9.625,
						Score:    1.0,
						OCRScore: 0.8,
						BoundingBox: []float64{
							0, 0.3396, 0.5293, 0.5205, 0.543,
						},
						BoundingRegion: []float64{
							0.3396, 0.543, 0.3396, 0.5293, 0.5205, 0.5293, 0.5205, 0.543,
						},
						Rotation: 0,
					},
					Total: &scheme.DetailedFloatField{
						Value:    1.93,
						Score:    1.0,
						OCRScore: 0.99,
						BoundingBox: []float64{
							0, 0.7349, 0.5278, 0.8223, 0.5386,
						},
						BoundingRegion: []float64{
							0.7349, 0.5278, 0.8223, 0.5278, 0.8223, 0.5386, 0.7349, 0.5386,
						},
						Rotation: 0,
					},
				},
			},

			BillTo: scheme.DetailedToField{
				Name: &scheme.DetailedField{
					Value:    "",
					Score:    0.95,
					OCRScore: 0.0,
					BoundingBox: nil,
					BoundingRegion: nil,
					Rotation: 0,
				},
				Address: &scheme.DetailedField{
					Value:    "",
					Score:    1.0,
					OCRScore: 0.0,
					BoundingBox: nil,
					BoundingRegion: nil,
					Rotation: 0,
				},
				ParsedAddress: nil,
				Email: &scheme.DetailedField{
					Value:    "",
					Score:    1.0,
					OCRScore: 0.0,
					BoundingBox: nil,
					BoundingRegion: nil,
					Rotation: 0,
				},
				VATNumber: &scheme.DetailedField{
					Value:    "",
					Score:    1.0,
					OCRScore: 0.0,
					BoundingBox: nil,
					BoundingRegion: nil,
					Rotation: 0,
				},
				PhoneNumber: &scheme.DetailedField{
					Value:    "",
					Score:    1.0,
					OCRScore: 0.0,
					BoundingBox: nil,
					BoundingRegion: nil,
					Rotation: 0,
				},
				RegNumber: &scheme.DetailedField{
					Value:    "",
					Score:    1.0,
					OCRScore: 0.0,
					BoundingBox: nil,
					BoundingRegion: nil,
					Rotation: 0,
				},
			},

			ShipTo: scheme.DetailedToField{
				Name: &scheme.DetailedField{
					Value:    "",
					Score:    1,
					OCRScore: 0,
					BoundingBox: nil,
					BoundingRegion: nil,
					Rotation: 0,
				},
				Address: &scheme.DetailedField{
					Value:    "",
					Score:    1,
					OCRScore: 0,
					BoundingBox: nil,
					BoundingRegion: nil,
					Rotation: 0,
				},
				ParsedAddress: nil,
			},
		}
	} else {
		mockReceiptData = fmt.Sprintf("%v/testdata/%v", pwd, "receipt_public.json")
		mockResp, err := os.ReadFile(mockReceiptData)
		assert.NoError(t, err)

		mockRespStr := string(mockResp)
		server.Serve(t, "/api/v8/partner/documents/36966934/", 200, mockRespStr)
		server.Serve(t, "/api/v8/partner/documents/", 200, mockRespStr)
		
		expected = &scheme.Document{
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
	server, client, _, expected := setUp(t, false)
	defer server.Close()

	resp, err := client.GetDocument("36966934", scheme.DocumentGetOptions{})
	assert.NotNil(t, resp)
	assert.NoError(t, err)
	assert.Equal(t, expected, resp)
}

func TestUnitClientV8_GetDetailedDocument(t *testing.T) {
	server, client, _, expected := setUp(t, true)
	defer server.Close()

	resp, err := client.GetDetailedDocument("36966934", scheme.DocumentGetOptions{})
	
	assert.NotNil(t, resp)
	assert.NoError(t, err)
	assert.Equal(t, expected, resp)
}


func TestUnitClientV8_ProcessDocumentUpload(t *testing.T) {
	server, client, mockReceiptPath, expected := setUp(t, false)
	defer server.Close()

	resp, err := client.ProcessDocumentUpload(scheme.DocumentUploadOptions{
		FilePath: mockReceiptPath,
	})
	assert.NotNil(t, resp)
	assert.NoError(t, err)
	assert.Equal(t, expected, resp)
}

func TestUnitClientV8_ProcessDocumentURL(t *testing.T) {
	server, client, _, expected := setUp(t, false)
	defer server.Close()

	resp, err := client.ProcessDocumentURL(scheme.DocumentURLOptions{
		FileURL: "http://cdn-dev.veryfi.com/testing/veryfi-python/receipt_public.jpg",
	})
	assert.NotNil(t, resp)
	assert.NoError(t, err)
	assert.Equal(t, expected, resp)
}
