package scheme

// DocumentUploadOptions describes the query parameters to process a multipart/form-data file upload.
type DocumentUploadOptions struct {
	FileName string `json:"file_name,omitempty"`
	File     []byte `json:"file,omitempty"`

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
	Categories        []string `json:"categories,omitempty"`
	Tags              []string `json:"tags,omitempty"`
	MaxPagesToProcess int      `json:"max_pages_to_process,omitempty"`
	BoostMode         int      `json:"boost_mode,omitempty"`
	ExternalID        string   `json:"external_id,omitempty"`
	Async             int      `json:"async,omitempty"`
}

// Document describes the response.
type Document struct {
	ABNNumber               string  `json:"abn_number"`
	AccountNumber           string  `json:"account_number"`
	BillToAddress           string  `json:"bill_to_address"`
	BillToName              string  `json:"bill_to_name"`
	BillToVATNumber         string  `json:"bill_to_vat_number"`
	CardNumber              string  `json:"card_number"`
	Category                string  `json:"category"`
	Created                 string  `json:"created"`
	CurrencyCode            string  `json:"currency_code"`
	Date                    string  `json:"date"`
	DeliveryDate            string  `json:"delivery_date"`
	Discount                float64 `json:"discount"`
	DocumentReferenceNumber string  `json:"document_reference_number"`
	DueDate                 string  `json:"due_date"`
	ExternalID              string  `json:"external_id"`
	ID                      int     `json:"id"`
	ImgFileName             string  `json:"img_file_name"`
	ImgThumbnailURL         string  `json:"img_thumbnail_url"`
	ImgURL                  string  `json:"img_url"`
	Insurance               string  `json:"insurance"`
	InvoiceNumber           string  `json:"invoice_number"`
	IsDuplicate             int     `json:"is_duplicate"`
	LineItems               []LineDocument
	OCRText                 string  `json:"ocr_text"`
	OrderDate               string  `json:"order_date"`
	PaymentDisplayName      string  `json:"payment_display_name"`
	PaymentTerms            string  `json:"payment_terms"`
	PaymentType             string  `json:"payment_type"`
	PhoneNumber             string  `json:"phone_number"`
	PurchaseOrderNumber     string  `json:"purchase_order_number"`
	Rounding                float64 `json:"rounding"`
	ServiceEndDate          string  `json:"service_end_date"`
	ServiceStartDate        string  `json:"service_start_date"`
	ShipDate                string  `json:"ship_date"`
	ShipToAddress           string  `json:"ship_to_address"`
	ShipToName              string  `json:"ship_to_name"`
	Shipping                float64 `json:"shipping"`
	StoreNumber             string  `json:"store_number"`
	Subtotal                float64 `json:"subtotal"`
	Tax                     float64 `json:"tax"`
	TaxLines                []TaxLine
	Tip                     float64 `json:"tip"`
	Total                   float64 `json:"total"`
	TotalWeight             string  `json:"total_weight"`
	TrackingNumber          string  `json:"tracking_number"`
	Updated                 string  `json:"updated"`
	VATNumber               string  `json:"vat_number"`
	Vendor                  Vendor
	VendorAccountNumber     string `json:"vendor_account_number"`
	VendorBankName          string `json:"vendor_bank_name"`
	VendorBankNumber        string `json:"vendor_bank_number"`
	VendorBankSwift         string `json:"vendor_bank_swift"`
	VendorIban              string `json:"vendor_iban"`
}

// LineDocument describes the line document response.
type LineDocument struct {
	Date          string  `json:"date"`
	Description   string  `json:"description"`
	Discount      float64 `json:"discount"`
	ID            int     `json:"id"`
	Order         int     `json:"order"`
	Price         float64 `json:"price"`
	Quantity      int     `json:"quantity"`
	Reference     string  `json:"reference"`
	Section       string  `json:"section"`
	SKU           string  `json:"sku"`
	Tax           float64 `json:"tax"`
	TaxRate       float64 `json:"tax_rate"`
	Total         float64 `json:"total"`
	Type          string  `json:"type"`
	UnitOfMeasure string  `json:"unit_of_measure"`
}

// TaxLine describes the tax line response.
type TaxLine struct {
	Name  string  `json:"name"`
	Order int     `json:"order"`
	Rate  float64 `json:"rate"`
	Total float64 `json:"total"`
}

// Vendor describes the vendor response.
type Vendor struct {
	Address         string `json:"address"`
	Category        string `json:"category"`
	Email           string `json:"email"`
	FaxNumber       string `json:"fax_number"`
	Name            string `json:"name"`
	PhoneNumber     string `json:"phone_number"`
	RawName         string `json:"raw_name"`
	VendorLogo      string `json:"vendor_logo"`
	VendorRegNumber string `json:"vendor_reg_number"`
	VendorType      string `json:"vendor_type"`
	Web             string `json:"web"`
}
