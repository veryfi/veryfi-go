package veryfi

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/go-resty/resty/v2"
	"github.com/veryfi/veryfi-go/veryfi/scheme"
)

// Client implements a Veryfi API Client.
type Client struct {
	// options is the global config options of the client.
	options *Options

	// client holds the resty.Client.
	client *resty.Client

	// apiVersion is the current API version of Veryfi that we are
	// communicating with.
	apiVersion string
}

// NewClientV7 returns a new instance of a client for v7 API.
func NewClientV7(opts *Options) (*Client, error) {
	c, err := createClient(opts)
	if err != nil {
		return nil, errors.Wrap(err, "fail to create a client")
	}

	return &Client{
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
func (c *Client) Config() *Options {
	return c.options
}

// ProcessDocumentUpload returns the processed document.
func (c *Client) ProcessDocumentUpload(opts scheme.DocumentUploadOptions) (*scheme.Document, error) {
	out := new(*scheme.Document)
	encodedFile, err := Base64EncodeFile(opts.FilePath)
	if err != nil {
		return nil, err
	}

	payload := scheme.DocumentUploadBase64Options{
		FileData:              encodedFile,
		DocumentSharedOptions: opts.DocumentSharedOptions,
	}
	if err := c.post(documentURI, payload, out); err != nil {
		return nil, err
	}

	return *out, nil
}

// ProcessDocumentUploadBase64 returns the processed base64 encoded document.
func (c *Client) ProcessDocumentUploadBase64(opts scheme.DocumentUploadBase64Options) (*scheme.Document, error) {
	out := new(*scheme.Document)
	if err := c.post(documentURI, opts, out); err != nil {
		return nil, err
	}

	return *out, nil
}

// ProcessDocumentURL returns the processed document using URL.
func (c *Client) ProcessDocumentURL(opts scheme.DocumentURLOptions) (*scheme.Document, error) {
	out := new(*scheme.Document)
	if err := c.post(documentURI, opts, out); err != nil {
		return nil, err
	}

	return *out, nil
}

// UpdateDocument updates and returns the processed document.
func (c *Client) UpdateDocument(documentID string, opts scheme.DocumentUpdateOptions) (*scheme.Document, error) {
	out := new(*scheme.Document)
	if err := c.put(fmt.Sprintf("%s%s", documentURI, documentID), opts, out); err != nil {
		return nil, err
	}

	return *out, nil
}

// SearchDocuments returns a list of processed documents with matching queries.
func (c *Client) SearchDocuments(opts scheme.DocumentSearchOptions) (*[]scheme.Document, error) {
	out := new(*[]scheme.Document)
	if err := c.get(documentURI, opts, out); err != nil {
		return nil, err
	}

	return *out, nil
}

// GetDocument returns a processed document with matching queries.
func (c *Client) GetDocument(documentID string, opts scheme.DocumentGetOptions) (*scheme.Document, error) {
	out := new(*scheme.Document)
	if err := c.get(fmt.Sprintf("%s%s", documentURI, documentID), opts, out); err != nil {
		return nil, err
	}

	return *out, nil
}

// DeleteDocument deletes a processed document.
func (c *Client) DeleteDocument(documentID string) error {
	err := c.rdelete(fmt.Sprintf("%s%s", documentURI, documentID))
	if err != nil {
		return err
	}

	return nil
}

// GetLineItems returns all line items for a processed document.
func (c *Client) GetLineItems(documentID string) (*scheme.LineItems, error) {
	out := new(*scheme.LineItems)
	if err := c.get(fmt.Sprintf("%s%s%s", documentURI, documentID, lineItemURI), nil, out); err != nil {
		return nil, err
	}

	return *out, nil
}

// AddLineItem returns a added line item for a processed document.
func (c *Client) AddLineItem(documentID string, opts scheme.LineItemOptions) (*scheme.LineItem, error) {
	out := new(*scheme.LineItem)
	if err := c.post(fmt.Sprintf("%s%s%s", documentURI, documentID, lineItemURI), opts, out); err != nil {
		return nil, err
	}

	return *out, nil
}

// GetLineItem returns a line item for a processed document.
func (c *Client) GetLineItem(documentID string, lineItemID string) (*scheme.LineItem, error) {
	out := new(*scheme.LineItem)
	if err := c.get(fmt.Sprintf("%s%s%s%s", documentURI, documentID, lineItemURI, lineItemID), nil, out); err != nil {
		return nil, err
	}

	return *out, nil
}

// UpdateLineItem returns an updated line item for a processed document.
func (c *Client) UpdateLineItem(documentID string, lineItemID string, opts scheme.LineItemOptions) (*scheme.LineItem, error) {
	out := new(*scheme.LineItem)
	if err := c.put(fmt.Sprintf("%s%s%s%s", documentURI, documentID, lineItemURI, lineItemID), opts, out); err != nil {
		return nil, err
	}

	return *out, nil
}

// DeleteLineItem deletes a line item in a document.
func (c *Client) DeleteLineItem(documentID string, lineItemID string) error {
	err := c.rdelete(fmt.Sprintf("%s%s%s%s", documentURI, documentID, lineItemURI, lineItemID))
	if err != nil {
		return err
	}

	return nil
}

// GetTags returns all tags for a processed document.
func (c *Client) GetTags(documentID string) (*scheme.Tags, error) {
	out := new(*scheme.Tags)
	if err := c.get(fmt.Sprintf("%s%s%s", documentURI, documentID, tagURI), nil, out); err != nil {
		return nil, err
	}

	return *out, nil
}

// GetGlobalTags returns all globally existing tags.
func (c *Client) GetGlobalTags() (*scheme.Tags, error) {
	out := new(*scheme.Tags)
	if err := c.get(globalTagURI, nil, out); err != nil {
		return nil, err
	}

	return *out, nil
}

// AddTag returns an added tag for a processed document.
func (c *Client) AddTag(documentID string, opts scheme.TagOptions) (*scheme.Tag, error) {
	out := new(*scheme.Tag)
	if err := c.put(fmt.Sprintf("%s%s%s", documentURI, documentID, tagURI), opts, out); err != nil {
		return nil, err
	}

	return *out, nil
}

// DeleteTag deletes a tag from a document.
func (c *Client) DeleteTag(documentID string, tagID string) error {
	err := c.rdelete(fmt.Sprintf("%s%s%s%s", documentURI, documentID, tagURI, tagID))
	if err != nil {
		return err
	}

	return nil
}

// DeleteGlobalTag deletes a tag from all documents.
func (c *Client) DeleteGlobalTag(tagID string) error {
	err := c.rdelete(fmt.Sprintf("%s%s", globalTagURI, tagID))
	if err != nil {
		return err
	}

	return nil
}

// request returns an authorized request to Veryfi API.
func (c *Client) request(okScheme interface{}, errScheme interface{}) *resty.Request {
	return c.setBaseURL().R().
		SetHeaders(map[string]string{
			"Content-Type":  "application/json",
			"Accept":        "application/json",
			"CLIENT-ID":     c.options.ClientID,
			"AUTHORIZATION": fmt.Sprintf("apikey %s:%s", c.options.Username, c.options.APIKey),
		}).
		SetResult(okScheme).
		SetError(errScheme)
}

// setBaseURL returns a client that uses Veryfi's base URL.
func (c *Client) setBaseURL() *resty.Client {
	return c.client.SetHostURL(buildURL(c.options.EnvironmentURL, "api", c.apiVersion))
}

// post performs a POST request against Veryfi API.
func (c *Client) post(uri string, body interface{}, okScheme interface{}) error {
	errScheme := new(scheme.Error)
	request := c.request(okScheme, errScheme).SetBody(body)

	_, err := request.Post(uri)

	return check(err, errScheme)
}

// put performs a PUT request against Veryfi API.
func (c *Client) put(uri string, body interface{}, okScheme interface{}) error {
	errScheme := new(scheme.Error)
	request := c.request(okScheme, errScheme).SetBody(body)
	_, err := request.Put(uri)

	return check(err, errScheme)
}

// get performs a GET request against Veryfi API.
func (c *Client) get(uri string, queryParams interface{}, okScheme interface{}) error {
	errScheme := new(scheme.Error)
	request := c.request(okScheme, errScheme)
	if queryParams != nil {
		request.SetQueryParams(structToMap(queryParams))
	}

	_, err := request.Get(uri)

	return check(err, errScheme)
}

// rdelete performs a DELETE request against Veryfi API.
func (c *Client) rdelete(uri string) error {
	errScheme := new(scheme.Error)
	request := c.request(map[string]string{}, errScheme)
	_, err := request.Delete(uri)

	return check(err, errScheme)
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
