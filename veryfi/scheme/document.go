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
	MaxPagesToProcess int      `json:"max_pages_to_process,omitempty"`
	BoostMode         int      `json:"boost_mode,omitempty"`
	ExternalID        string   `json:"external_id,omitempty"`
	Async             int      `json:"async,omitempty"`
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
}

// VendorUpdateOptions describes the update options for vendor.
type VendorUpdateOptions struct {
	Name    string `json:"name,omitempty"`
	Address string `json:"address,omitempty"`
}

// DocumentSearchOptions describes the query parameters to search document.
type DocumentSearchOptions struct {
	Q          string `json:"q"`
	ExternalID string `json:"external_id"`
	Tag        string `json:"tag"`
	CreatedGT  string `json:"created__gt"`
	CreatedGTE string `json:"created__gte"`
	CreatedLT  string `json:"created__lt"`
	CreatedLTE string `json:"created__lte"`
}

// DocumentGetOptions describes the query parameters to get a document.
type DocumentGetOptions struct {
	ReturnAuditTrail string `json:"return_audit_trail"`
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

// Document describes the response.
type Document struct {
	ABNNumber           string     `json:"abn_number"`
	AccountNumber       string     `json:"account_number"`
	BillToAddress       string     `json:"bill_to_address"`
	BillToName          string     `json:"bill_to_name"`
	BillToVATNumber     string     `json:"bill_to_vat_number"`
	CardNumber          string     `json:"card_number"`
	Category            string     `json:"category"`
	Created             string     `json:"created"`
	CurrencyCode        string     `json:"currency_code"`
	Date                string     `json:"date"`
	DeliveryDate        string     `json:"delivery_date"`
	Discount            float64    `json:"discount"`
	ReferenceNumber     string     `json:"reference_number"`
	DueDate             string     `json:"due_date"`
	ExternalID          string     `json:"external_id"`
	ID                  int        `json:"id"`
	ImgFileName         string     `json:"img_file_name"`
	ImgThumbnailURL     string     `json:"img_thumbnail_url"`
	ImgURL              string     `json:"img_url"`
	Insurance           string     `json:"insurance"`
	InvoiceNumber       string     `json:"invoice_number"`
	IsDuplicate         int        `json:"is_duplicate"`
	LineItems           []LineItem `json:"line_items"`
	OCRText             string     `json:"ocr_text"`
	OrderDate           string     `json:"order_date"`
	PaymentDisplayName  string     `json:"payment_display_name"`
	PaymentTerms        string     `json:"payment_terms"`
	PaymentType         string     `json:"payment_type"`
	PhoneNumber         string     `json:"phone_number"`
	PurchaseOrderNumber string     `json:"purchase_order_number"`
	Rounding            float64    `json:"rounding"`
	ServiceEndDate      string     `json:"service_end_date"`
	ServiceStartDate    string     `json:"service_start_date"`
	ShipDate            string     `json:"ship_date"`
	ShipToAddress       string     `json:"ship_to_address"`
	ShipToName          string     `json:"ship_to_name"`
	Shipping            float64    `json:"shipping"`
	StoreNumber         string     `json:"store_number"`
	Subtotal            float64    `json:"subtotal"`
	Tax                 float64    `json:"tax"`
	TaxLines            []TaxLine  `json:"tax_lines"`
	Tip                 float64    `json:"tip"`
	Total               float64    `json:"total"`
	TotalWeight         string     `json:"total_weight"`
	TrackingNumber      string     `json:"tracking_number"`
	Updated             string     `json:"updated"`
	VATNumber           string     `json:"vat_number"`
	Vendor              Vendor     `json:"vendor"`
	VendorAccountNumber string     `json:"vendor_account_number"`
	VendorBankName      string     `json:"vendor_bank_name"`
	VendorBankNumber    string     `json:"vendor_bank_number"`
	VendorBankSwift     string     `json:"vendor_bank_swift"`
	VendorIban          string     `json:"vendor_iban"`
}

// LineItems describes the line items in a document response.
type LineItems struct {
	LineItems []LineItem `json:"line_items"`
}

// LineItem describes the line item in a document response.
type LineItem struct {
	Date          string  `json:"date"`
	Description   string  `json:"description"`
	Discount      float64 `json:"discount"`
	ID            int     `json:"id"`
	Order         int     `json:"order"`
	Price         float64 `json:"price"`
	Quantity      float64 `json:"quantity"`
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
