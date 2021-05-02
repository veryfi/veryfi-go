package veryfi

import (
	"fmt"
	"github.com/pkg/errors"

	"github.com/go-resty/resty/v2"
	"github.com/hoanhan101/veryfi-go/veryfi/scheme"
)

// httpClient implements a Veryfi API Client.
type httpClient struct {
	// options is the global config options of the client.
	options *Options

	// client holds the resty.Client.
	client *resty.Client

	// apiVersion is the current API version of Veryfi that we are
	// communicating with.
	apiVersion string
}

// NewClientV7 returns a new instance of a client for v7 API.
func NewClientV7(opts *Options) (Client, error) {
	c, err := createClient(opts)
	if err != nil {
		return nil, errors.Wrap(err, "fail to create a client")
	}

	return &httpClient{
		options:    opts,
		client:     c,
		apiVersion: "v7",
	}, nil
}

// createClient setups a resty client with configured options.
func createClient(opts *Options) (*resty.Client, error) {
	err := setDefaults(opts)
	if err != nil {
		return nil, err
	}

	// Create a resty client with configured options.
	client := resty.New()
	client = client.
		SetTimeout(opts.HTTP.Timeout).
		SetRetryCount(int(opts.HTTP.Retry.Count)).
		SetRetryWaitTime(opts.HTTP.Retry.WaitTime).
		SetRetryMaxWaitTime(opts.HTTP.Retry.MaxWaitTime)

	return client, nil
}

// Config returns the client configuration options.
func (c *httpClient) Config() *Options {
	return c.options
}

// ProcessDocumentUpload returns the processed document using file upload.
func (c *httpClient) ProcessDocumentUpload(opts scheme.DocumentUploadOptions) (*scheme.Document, error) {
	out := new(*scheme.Document)
	if err := c.post(documentURI, opts.FilePath, opts.DocumentSharedOptions, out); err != nil {
		return nil, err
	}

	return *out, nil
}

// ProcessDocumentURL returns the processed document using URL.
func (c *httpClient) ProcessDocumentURL(opts scheme.DocumentURLOptions) (*scheme.Document, error) {
	out := new(*scheme.Document)
	if err := c.post(documentURI, "", opts, out); err != nil {
		return nil, err
	}

	return *out, nil
}

// post performs a POST request against Veryfi API.
func (c *httpClient) post(uri string, filePath string, body interface{}, okScheme interface{}) error {
	errScheme := new(scheme.Error)
	request := c.setBaseURL().R().
		SetBody(body).
		SetHeaders(map[string]string{
			"Content-Type":  "application/json",
			"Accept":        "application/json",
			"CLIENT-ID":     c.options.ClientID,
			"AUTHORIZATION": fmt.Sprintf("apikey %s:%s", c.options.Username, c.options.APIKey),
		}).
		SetResult(okScheme).
		SetError(errScheme)

	if len(filePath) != 0 {
		request.SetFile("file", filePath)
	}

	_, err := request.Post(uri)

	return check(err, errScheme)
}

// setBaseURL returns a client that uses Veryfi's base URL.
func (c *httpClient) setBaseURL() *resty.Client {
	return c.client.SetHostURL(buildURL(c.options.EnvironmentURL, "api", c.apiVersion))
}

// check validates returned response from Veryfi.
func check(err error, errResp *scheme.Error) error {
	if err != nil {
		return errors.Wrap(err, "fail to make a request to Veryfi")
	}

	// Parse down to a more meaningful error.
	if *errResp != (scheme.Error{}) {
		ctx := ""
		if len(errResp.Error) > 0 {
			ctx = errResp.Error
		}
		if len(errResp.Message) > 0 {
			ctx = errResp.Message
		}

		return errors.Errorf(
			"get a response from Veryfi with status=%s and context=%s",
			errResp.Status, ctx,
		)
	}

	return nil
}
