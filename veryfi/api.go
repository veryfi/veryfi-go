package veryfi

import (
	"github.com/hoanhan101/veryfi-go/veryfi/scheme"
)

// Client API for Veryfi.
type Client interface {
	// Config returns the unified configuration info.
	Config() *Options

	// ProcessDocumentUpload processes a file using a multipart/form-data file upload.
	ProcessDocumentUpload(scheme.DocumentUploadOptions) (*scheme.Document, error)

	// ProcessDocumentUploadBase64 processes a Base64 encoded document.
	ProcessDocumentUploadBase64(scheme.DocumentUploadBase64Options) (*scheme.Document, error)

	// ProcessDocumentURL processes a file using a URL.
	ProcessDocumentURL(scheme.DocumentURLOptions) (*scheme.Document, error)

	// UpdateDocument updates a previously processed document.
	UpdateDocument(string, scheme.DocumentUpdateOptions) (*scheme.Document, error)
}
