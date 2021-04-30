package veryfi

import (
	"github.com/hoanhan101/veryfi-go/veryfi/scheme"
)

// Client API for Veryfi.
type Client interface {
	// Config returns the unified configuration info.
	Config() *Options

	// ProcessDocumentURL processes a file using a URL.
	ProcessDocumentURL(scheme.DocumentURLOptions) (*scheme.Document, error)
}
