package scheme

// DocumentUploadOptions describes the query parameters to process a multipart/form-data file upload.
type DocumentUploadOptions struct {
	FilePath string

	DocumentSharedOptions
}

// DocumentUploadBase64Options describes the query parameters to process a Base64 encoded document.
type DocumentUploadBase64Options struct {
	FileData string `json:"file_data,omitempty"`

	DocumentSharedOptions
}

// DocumentURLOptions describes the query parameters to process a file using a URL.
type DocumentURLOptions struct {
	FileURL  string   `json:"file_url,omitempty"`
	FileURLS []string `json:"file_urls,omitempty"`

	DocumentSharedOptions
}

// DocumentSharedOptions describes the shared query parameters among the processing API.
type DocumentSharedOptions struct {
	FileName          string   `json:"file_name,omitempty"`
	Categories        []string `json:"categories,omitempty"`
	Tags              []string `json:"tags,omitempty"`
	MaxPagesToProcess *int     `json:"max_pages_to_process,omitempty"`
	BoostMode         bool     `json:"boost_mode,omitempty"`
	AutoDelete        bool     `json:"auto_delete,omitempty"`
	DetectBlur        bool     `json:"detect_blur,omitempty"`
	ParseAddress      bool     `json:"parse_address,omitempty"`
	ExternalID        string   `json:"external_id,omitempty"`
	Async             bool     `json:"async,omitempty"`
	ConfidenceDetails bool     `json:"confidence_details,omitempty"`
	BoundingBoxes     bool     `json:"bounding_boxes,omitempty"`
}

// DocumentUpdateOptions describes the query parameters to update a document.
type DocumentUpdateOptions struct {
	BillToName    string              `json:"bill_to_name,omitempty"`
	BillToAddress string              `json:"bill_to_address,omitempty"`
	Category      string              `json:"category,omitempty"`
	Date          string              `json:"date,omitempty"`
	DueDate       string              `json:"due_date,omitempty"`
	InvoiceNumber string              `json:"invoice_number,omitempty"`
	Subtotal      float64             `json:"subtotal,omitempty"`
	Tax           float64             `json:"tax,omitempty"`
	Tip           float64             `json:"tip,omitempty"`
	Total         float64             `json:"total,omitempty"`
	Vendor        VendorUpdateOptions `json:"vendor,omitempty"`
	ExternalID    string              `json:"external_id,omitempty"`
	Status        DocumentStatus      `json:"status,omitempty"` // Possible values: [processed, reviewed, archived]
}

// VendorUpdateOptions describes the update options for vendor.
type VendorUpdateOptions struct {
	Name    string `json:"name,omitempty"`
	Address string `json:"address,omitempty"`
}

// DocumentSearchOptions describes the query parameters to search document.
type DocumentSearchOptions struct {
	Q                 string         `json:"q"`
	ExternalID        string         `json:"external_id"`
	Tag               string         `json:"tag"`
	CreatedGT         string         `json:"created__gt"`
	CreatedGTE        string         `json:"created__gte"`
	CreatedLT         string         `json:"created__lt"`
	CreatedLTE        string         `json:"created__lte"`
	Status            DocumentStatus `json:"status"`
	DeviceID          string         `json:"device_id"`
	Owner             string         `json:"owner"`
	UpdatedGT         string         `json:"updated__gt"`
	UpdatedGTE        string         `json:"updated__gte"`
	UpdatedLT         string         `json:"updated__lt"`
	UpdatedLTE        string         `json:"updated__lte"`
	DateGT            string         `json:"date__gt"`
	DateGTE           string         `json:"date__gte"`
	DateLT            string         `json:"date__lt"`
	DateLTE           string         `json:"date__lte"`
	Page              string         `json:"page"`
	PageSize          string         `json:"page_size"`
	TrackTotalResults string         `json:"track_total_results"`
}

type DetailedDocumentSearchOptions struct {
	Q                 string         `json:"q"`
	ExternalID        string         `json:"external_id"`
	Tag               string         `json:"tag"`
	CreatedGT         string         `json:"created__gt"`
	CreatedGTE        string         `json:"created__gte"`
	CreatedLT         string         `json:"created__lt"`
	CreatedLTE        string         `json:"created__lte"`
	Status            DocumentStatus `json:"status"`
	DeviceID          string         `json:"device_id"`
	Owner             string         `json:"owner"`
	UpdatedGT         string         `json:"updated__gt"`
	UpdatedGTE        string         `json:"updated__gte"`
	UpdatedLT         string         `json:"updated__lt"`
	UpdatedLTE        string         `json:"updated__lte"`
	DateGT            string         `json:"date__gt"`
	DateGTE           string         `json:"date__gte"`
	DateLT            string         `json:"date__lt"`
	DateLTE           string         `json:"date__lte"`
	Page              string         `json:"page"`
	PageSize          string         `json:"page_size"`
	TrackTotalResults string         `json:"track_total_results"`
	BoundingBoxes     bool           `json:"bounding_boxes"`
	ConfidenceDetails bool           `json:"confidence_details"`
}

// DocumentGetOptions describes the query parameters to get a document.
type DocumentGetOptions struct {
	ReturnAuditTrail string `json:"return_audit_trail"`
}

type DocumentGetDetailedOptions struct {
	ReturnAuditTrail  string `json:"return_audit_trail"`
	BoundingBoxes     bool   `json:"bounding_boxes"`
	ConfidenceDetails bool   `json:"confidence_details"`
}

// LineItemOptions describes the query parameters to add a line to a document.
type LineItemOptions struct {
	Order         int     `json:"order"`
	SKU           string  `json:"sku"`
	Description   string  `json:"description"`
	Category      string  `json:"category"`
	Total         float64 `json:"total"`
	Tax           float64 `json:"tax"`
	Price         float64 `json:"price"`
	UnitOfMeasure string  `json:"unit_of_measure"`
	Quantity      float64 `json:"quantity"`
}

type DocumentStatus string

const (
	Processed DocumentStatus = "processed"
	Reviewed  DocumentStatus = "reviewed"
	Archived  DocumentStatus = "archived"
)

type Documents struct {
	Documents []Document    `json:"documents"`
	Meta      DocumentsMeta `json:"meta"`
}

type DocumentsMeta struct {
	DocumentsPerPage int `json:"documents_per_page"`
	PageNumber       int `json:"page_number"`
	TotalPages       int `json:"total_pages"`
	TotalResults     int `json:"total_results"`
}

// Document describes the response.
type Document struct {
	AccountingEntryType     string         `json:"accounting_entry_type"`
	AccountNumber           string         `json:"account_number"`
	Balance                 string         `json:"balance"`
	Barcodes                []string       `json:"barcodes"`
	BillTo                  ToField        `json:"bill_to"`
	Cashback                string         `json:"cashback"`
	Category                string         `json:"category"`
	Created                 string         `json:"created_date"`
	CountryCode             string         `json:"country_code"`
	CurrencyCode            string         `json:"currency_code"`
	Date                    string         `json:"date"`
	DefaultCategory         string         `json:"default_category"`
	DeliveryDate            string         `json:"delivery_date"`
	DeliveryNoteNumber      string         `json:"delivery_note_number"`
	Discount                float64        `json:"discount"`
	DocumentReferenceNumber string         `json:"document_reference_number"`
	DocumentTitle           string         `json:"document_title"`
	DuplicateOf             int            `json:"duplicate_of"`
	DueDate                 string         `json:"due_date"`
	ExchangeRate            float64        `json:"exch_rate"`
	ExternalID              string         `json:"external_id"`
	FinalBalance            float64        `json:"final_balance"`
	GuestCount              string         `json:"guest_count"`
	ID                      int            `json:"id"`
	ImgFileName             string         `json:"img_file_name"`
	ImgThumbnailURL         string         `json:"img_thumbnail_url"`
	ImgURL                  string         `json:"img_url"`
	Incoterms               string         `json:"incoterms"`
	Insurance               float64        `json:"insurance"`
	InvoiceNumber           string         `json:"invoice_number"`
	IsApproved              bool           `json:"is_approved"`
	IsDocument              bool           `json:"is_document"`
	IsDuplicate             bool           `json:"is_duplicate"`
	IsMoneyIn               bool           `json:"is_money_in"`
	IsTransaction           bool           `json:"is_transaction"`
	LicensePlateNumber      string         `json:"license_plate_number"`
	LineItems               []LineItem     `json:"line_items_with_scores"`
	Model                   string         `json:"model"`
	Notes                   string         `json:"notes"`
	OCRText                 string         `json:"ocr_text"`
	OrderDate               string         `json:"order_date"`
	Payment                 PaymentsInfo   `json:"payment"`
	PaymentLinks            []string       `json:"payment_links"`
	PDFURL                  string         `json:"pdf_url"`
	PreviousBalance         float64        `json:"previous_balance"`
	PurchaseOrderNumber     string         `json:"purchase_order_number"`
	ReferenceNumber         string         `json:"reference_number"`
	Rounding                float64        `json:"rounding"`
	ServerName              string         `json:"server_name"`
	ServiceEndDate          string         `json:"service_end_date"`
	ServiceStartDate        string         `json:"service_start_date"`
	ShipDate                string         `json:"ship_date"`
	Shipping                float64        `json:"shipping"`
	ShipTo                  ToField        `json:"ship_to"`
	Status                  DocumentStatus `json:"status"`
	StoreNumber             string         `json:"store_number"`
	Subtotal                float64        `json:"subtotal"`
	Tags                    []string       `json:"tags"`
	Tax                     float64        `json:"tax"`
	TaxLines                []TaxLine      `json:"tax_lines_with_scores"`
	Tip                     float64        `json:"tip"`
	Total                   float64        `json:"total"`
	TotalQuantity           float64        `json:"total_quantity"`
	TotalWeight             string         `json:"total_weight"`
	TrackingNumber          string         `json:"tracking_number"`
	TrackingNumbers         []string       `json:"tracking_numbers"`
	Updated                 string         `json:"updated_date"`
	VendingPerson           string         `json:"vending_person"`
	VendingPersonNumber     string         `json:"vending_person_number"`
	Vendor                  Vendor         `json:"vendor"`
	Vendors                 []string       `json:"vendors"`
	VINNumber               string         `json:"vin_number"`
	Warnings                []string       `json:"warnings"`
	Weights                 []string       `json:"weights"`
}

// ToField describes the to field response.
type ToField struct {
	Name          string        `json:"name"`
	Address       string        `json:"address"`
	ParsedAddress ParsedAddress `json:"parsed_address"`
	Email         string        `json:"email"`
	VATNumber     string        `json:"vat_number"`
	PhoneNumber   string        `json:"phone_number"`
	RegNumber     string        `json:"reg_number"`
}

// ParsedAddress describes the parsed address response.
type ParsedAddress struct {
	City          string `json:"city"`
	Country       string `json:"country"`
	Postcode      string `json:"postcode"`
	State         string `json:"state"`
	StreetAddress string `json:"street_address"`
	House         string `json:"house"`
	HouseNumber   string `json:"house_number"`
	Road          string `json:"road"`
	Unit          string `json:"unit"`
	Level         string `json:"level"`
	Staircase     string `json:"staircase"`
	POBox         string `json:"po_box"`
	Suburb        string `json:"suburb"`
	CityDistrict  string `json:"city_district"`
	Island        string `json:"island"`
	StateDistrict string `json:"state_district"`
	CountryRegion string `json:"country_region"`
	WorldRegion   string `json:"world_region"`
}

// PaymentsInfo describes the payment response.
type PaymentsInfo struct {
	CardNumber  string `json:"card_number"`
	DisplayName string `json:"display_name"`
	Terms       string `json:"terms"`
	Type        string `json:"type"`
}

// LineItems describes the line items in a document response.
type LineItems struct {
	LineItems []LineItem `json:"line_items"`
}

// LineItem describes the line item in a document response.
type LineItem struct {
	Category              string      `json:"category"`
	CountryOfOrigin       string      `json:"country_of_origin"`
	Date                  string      `json:"date"`
	Description           string      `json:"description"`
	Discount              float64     `json:"discount"`
	DiscountPrice         float64     `json:"discount_price"`
	DiscountRate          float64     `json:"discount_rate"`
	EndDate               string      `json:"end_date"`
	FullDescription       string      `json:"full_description"`
	GrossTotal            float64     `json:"gross_total"`
	HSN                   string      `json:"hsn"`
	ID                    int         `json:"id"`
	Lot                   string      `json:"lot"`
	Manufacturer          string      `json:"manufacturer"`
	NetTotal              float64     `json:"net_total"`
	NormalizedDescription string      `json:"normalized_description"`
	Order                 int         `json:"order"`
	Price                 float64     `json:"price"`
	ProductInfo           ProductInfo `json:"product_info"`
	Quantity              float64     `json:"quantity"`
	Reference             string      `json:"reference"`
	Section               string      `json:"section"`
	SKU                   string      `json:"sku"`
	StartDate             string      `json:"start_date"`
	Subtotal              float64     `json:"subtotal"`
	Tags                  []string    `json:"tags"`
	Tax                   float64     `json:"tax"`
	TaxCode               string      `json:"tax_code"`
	TaxRate               float64     `json:"tax_rate"`
	Text                  string      `json:"text"`
	Total                 float64     `json:"total"`
	Type                  string      `json:"type"`
	UnitOfMeasure         string      `json:"unit_of_measure"`
	UPC                   string      `json:"upc"`
	Weight                string      `json:"weight"`
}

// ProductInfo describes the product info in a document response.
type ProductInfo struct {
	Brand               string   `json:"brand"`
	Category            []string `json:"category"`
	ExpandedDescription string   `json:"expanded_description"`
}

// TaxLine describes the tax line response.
type TaxLine struct {
	Base           float64 `json:"base"`
	Code           string  `json:"code"`
	Name           string  `json:"name"`
	Order          int     `json:"order"`
	Rate           float64 `json:"rate"`
	Total          float64 `json:"total"`
	TotalInclusive float64 `json:"total_inclusive"`
}

// Vendor describes the vendor response.
type Vendor struct {
	ABNNumber       string        `json:"abn_number"`
	AccountCurrency string        `json:"account_currency"`
	AccountNumber   string        `json:"account_number"`
	BankName        string        `json:"bank_name"`
	BankNumber      string        `json:"bank_number"`
	BankSwift       string        `json:"bank_swift"`
	ExternalID      string        `json:"external_id"`
	FaxNumber       string        `json:"fax_number"`
	FullAddress     string        `json:"full_address"`
	IBAN            string        `json:"iban"`
	RawName         string        `json:"raw_name"`
	Types           string        `json:"types"`
	Web             string        `json:"web"`
	Name            string        `json:"name"`
	Address         string        `json:"address"`
	ParsedAddress   ParsedAddress `json:"parsed_address"`
	Email           string        `json:"email"`
	VATNumber       string        `json:"vat_number"`
	PhoneNumber     string        `json:"phone_number"`
	RegNumber       string        `json:"reg_number"`
	Logo            string        `json:"logo"`
	Lat             float64       `json:"lat"`
	Lng             float64       `json:"lng"`
	Type            string        `json:"type"`
}

// Tags describes the tags response.
type Tags struct {
	Tags []Tag `json:"tags"`
}

// Tag describes the tag response.
type Tag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// TagOptions describes query parameters to update a tag in a document.
type TagOptions struct {
	Name string `json:"name"`
}

// DetailedField represents a field with confidence scores and metadata
type DetailedField struct {
	Value          string    `json:"value,omitempty"`
	Score          *float64  `json:"score,omitempty"`
	OCRScore       *float64  `json:"ocr_score,omitempty"`
	BoundingBox    []float64 `json:"bounding_box,omitempty"`
	BoundingRegion []float64 `json:"bounding_region,omitempty"`
	Rotation       int       `json:"rotation,omitempty"`
}

// DetailedFloatField represents a numeric field with confidence scores
type DetailedFloatField struct {
	Value          float64   `json:"value,omitempty"`
	Score          *float64  `json:"score,omitempty"`
	OCRScore       *float64  `json:"ocr_score,omitempty"`
	BoundingBox    []float64 `json:"bounding_box,omitempty"`
	BoundingRegion []float64 `json:"bounding_region,omitempty"`
	Rotation       int       `json:"rotation,omitempty"`
}

// DetailedDateField represents a date field with confidence scores
type DetailedDateField struct {
	Value          string    `json:"value,omitempty"` // ISO 8601 date format
	Score          *float64  `json:"score,omitempty"`
	OCRScore       *float64  `json:"ocr_score,omitempty"`
	BoundingBox    []float64 `json:"bounding_box,omitempty"`
	BoundingRegion []float64 `json:"bounding_region,omitempty"`
	Rotation       int       `json:"rotation,omitempty"`
}

// DetailedBoolField represents a boolean field with confidence scores
type DetailedBoolField struct {
	Value          bool      `json:"value,omitempty"`
	Score          *float64  `json:"score,omitempty"`
	OCRScore       *float64  `json:"ocr_score,omitempty"`
	BoundingBox    []float64 `json:"bounding_box,omitempty"`
	BoundingRegion []float64 `json:"bounding_region,omitempty"`
	Rotation       int       `json:"rotation,omitempty"`
}

// DetailedVendor extends Vendor with detailed fields
type DetailedVendor struct {
	ABNNumber       *DetailedField `json:"abn_number,omitempty"`
	AccountCurrency *DetailedField `json:"account_currency,omitempty"`
	AccountNumber   *DetailedField `json:"account_number,omitempty"`
	BankName        *DetailedField `json:"bank_name,omitempty"`
	BankNumber      *DetailedField `json:"bank_number,omitempty"`
	BankSwift       *DetailedField `json:"bank_swift,omitempty"`
	ExternalID      string         `json:"external_id,omitempty"`
	FaxNumber       *DetailedField `json:"fax_number,omitempty"`
	FullAddress     *DetailedField `json:"full_address,omitempty"`
	IBAN            *DetailedField `json:"iban,omitempty"`
	RawName         *DetailedField `json:"raw_name,omitempty"`
	Types           *DetailedField `json:"types,omitempty"`
	Web             *DetailedField `json:"web,omitempty"`
	Name            *DetailedField `json:"name,omitempty"`
	Address         *DetailedField `json:"address,omitempty"`
	ParsedAddress   *ParsedAddress `json:"parsed_address,omitempty"`
	Email           *DetailedField `json:"email,omitempty"`
	VATNumber       *DetailedField `json:"vat_number,omitempty"`
	PhoneNumber     *DetailedField `json:"phone_number,omitempty"`
	RegNumber       *DetailedField `json:"reg_number,omitempty"`
	Logo            string         `json:"logo,omitempty"`
	Lat             *float64       `json:"lat,omitempty"`
	Lng             *float64       `json:"lng,omitempty"`
	Type            *DetailedField `json:"type"`
}

// DetailedBillTo represents bill_to information with confidence scores
type DetailedToField struct {
	Name          *DetailedField `json:"name"`
	Address       *DetailedField `json:"address"`
	ParsedAddress *ParsedAddress `json:"parsed_address"`
	Email         *DetailedField `json:"email"`
	VATNumber     *DetailedField `json:"vat_number"`
	PhoneNumber   *DetailedField `json:"phone_number"`
	RegNumber     *DetailedField `json:"reg_number"`
}

// DetailedPayment represents payment information with confidence scores
type DetailedPayment struct {
	CardNumber  *DetailedField `json:"card_number"`
	DisplayName string         `json:"display_name"`
	Terms       *DetailedField `json:"terms"`
	Type        *DetailedField `json:"type"`
}

// DetailedLineItem extends LineItem with confidence scores
type DetailedLineItem struct {
	Category    *DetailedField      `json:"category"`
	Date        *DetailedDateField  `json:"date"`
	Description *DetailedField      `json:"description"`
	Discount    *DetailedFloatField `json:"discount"`
	ID          int                 `json:"id"`
	Order       int                 `json:"order"`
	Price       *DetailedFloatField `json:"price"`
	Quantity    *DetailedFloatField `json:"quantity"`
	Reference   *DetailedField      `json:"reference"`
	Section     *DetailedField      `json:"section"`
	SKU         *DetailedField      `json:"sku"`
	UPC         *DetailedField      `json:"upc"`
	Tags        []string            `json:"tags"`

	Tax           *DetailedFloatField `json:"tax"`
	TaxRate       *DetailedFloatField `json:"tax_rate"`
	Total         *DetailedFloatField `json:"total"`
	Type          string              `json:"type"`
	UnitOfMeasure *DetailedField      `json:"unit_of_measure"`
}

// DetailedTaxLine extends TaxLine with confidence scores
type DetailedTaxLine struct {
	Order int                 `json:"order"`
	Name  *DetailedField      `json:"name"`
	Rate  *DetailedFloatField `json:"rate"`
	Total *DetailedFloatField `json:"total"`
}

type DetailedDocuments struct {
	Documents []DetailedDocument `json:"documents"`
	Meta      DocumentsMeta      `json:"meta"`
}

// DetailedDocument extends Document with detailed field information
type DetailedDocument struct {
	ABNNumber           *DetailedField      `json:"abn_number"`
	AccountNumber       *DetailedField      `json:"account_number"`
	BillTo              DetailedToField     `json:"bill_to"`
	CardNumber          *DetailedField      `json:"card_number"`
	Category            *DetailedField      `json:"category"`
	Created             *DetailedField      `json:"created"`
	CurrencyCode        *DetailedField      `json:"currency_code"`
	Date                *DetailedDateField  `json:"date"`
	DeliveryDate        *DetailedDateField  `json:"delivery_date"`
	Discount            *DetailedFloatField `json:"discount"`
	ReferenceNumber     string              `json:"reference_number"`
	DueDate             *DetailedDateField  `json:"due_date"`
	ExternalID          string              `json:"external_id"`
	ID                  int                 `json:"id"`
	ImgFileName         string              `json:"img_file_name"`
	ImgThumbnailURL     string              `json:"img_thumbnail_url"`
	ImgURL              string              `json:"img_url"`
	Insurance           *DetailedFloatField `json:"insurance"`
	InvoiceNumber       *DetailedField      `json:"invoice_number"`
	IsDuplicate         bool                `json:"is_duplicate"`
	LineItems           []LineItem          `json:"line_items"`
	LineItemsWithScores []DetailedLineItem  `json:"line_items_with_scores"`
	OCRText             string              `json:"ocr_text"`
	OrderDate           *DetailedDateField  `json:"order_date"`
	Payment             *DetailedPayment    `json:"payment"`
	PhoneNumber         *DetailedField      `json:"phone_number"`
	PurchaseOrderNumber *DetailedField      `json:"purchase_order_number"`
	Rounding            *DetailedFloatField `json:"rounding"`
	ServiceEndDate      *DetailedDateField  `json:"service_end_date"`
	ServiceStartDate    *DetailedDateField  `json:"service_start_date"`
	ShipDate            *DetailedDateField  `json:"ship_date"`
	ShipTo              DetailedToField     `json:"ship_to"`
	Status              DocumentStatus      `json:"status"`
	StoreNumber         *DetailedField      `json:"store_number"`
	Subtotal            *DetailedFloatField `json:"subtotal"`
	Tags                []string            `json:"tags"`
	Tax                 *DetailedFloatField `json:"tax"`
	TaxLines            []TaxLine           `json:"tax_lines"`
	TaxLinesWithScores  []DetailedTaxLine   `json:"tax_lines_with_scores"`
	Tip                 *DetailedFloatField `json:"tip"`
	Total               *DetailedFloatField `json:"total"`
	TotalWeight         *DetailedField      `json:"total_weight"`
	TrackingNumber      *DetailedField      `json:"tracking_number"`
	Updated             string              `json:"updated"`
	VATNumber           *DetailedField      `json:"vat_number"`
	Vendor              *DetailedVendor     `json:"vendor"`
	VendorIban          *DetailedField      `json:"vendor_iban"`
}
