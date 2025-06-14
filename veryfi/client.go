package veryfi

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/go-resty/resty/v2"
	"github.com/veryfi/veryfi-go/v3/veryfi/scheme"
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

	// pkgVersion is the current SDK version.
	pkgVersion string
}

// NewClientV8 returns a new instance of a client for v8 API.
func NewClientV8(opts *Options) (*Client, error) {
	c, err := createClient(opts)
	if err != nil {
		return nil, errors.Wrap(err, "fail to create a client")
	}

	return &Client{
		options:    opts,
		client:     c,
		apiVersion: "v8",
		pkgVersion: "2.1.1",
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
		SetRetryMaxWaitTime(opts.HTTP.Retry.MaxWaitTime).
		OnAfterResponse(func(c *resty.Client, resp *resty.Response) error {
			if resp.IsError() && resp.Error() != nil {
				errorStruct := resp.Error().(*scheme.Error)
				errorStruct.Status = resp.RawResponse.Status
				body := resp.Body()
				if len(body) > 0 {
					c.JSONUnmarshal(body, errorStruct)
				}
			}
			return nil
		})

	return client, nil
}

// Config returns the client configuration options.
func (c *Client) Config() *Options {
	return c.options
}

// SetTLSConfig sets the TLS configurations for underling transportation layer.
func (c *Client) SetTLSConfig(config *tls.Config) {
	c.client.SetTLSClientConfig(config)
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

// ProcessDetailedDocumentUpload returns the processed document with confidence scores and bounding boxes
func (c *Client) ProcessDetailedDocumentUpload(opts scheme.DocumentUploadOptions) (*scheme.DetailedDocument, error) {
	out := new(*scheme.DetailedDocument)
	encodedFile, err := Base64EncodeFile(opts.FilePath)
	if err != nil {
		return nil, err
	}

	payload := scheme.DocumentUploadBase64Options{
		FileData:              encodedFile,
		DocumentSharedOptions: opts.DocumentSharedOptions,
	}
	// Always enable confidence details and bounding boxes
	payload.DocumentSharedOptions.ConfidenceDetails = true
	payload.DocumentSharedOptions.BoundingBoxes = true
	if err := c.post(documentURI, payload, out); err != nil {
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

// ProcessDetailedDocumentURL returns the processed document using URL with confidence scores and bounding boxes.
func (c *Client) ProcessDetailedDocumentURL(opts scheme.DocumentURLOptions) (*scheme.DetailedDocument, error) {
	out := new(*scheme.DetailedDocument)
	opts.DocumentSharedOptions.ConfidenceDetails = true
	opts.DocumentSharedOptions.BoundingBoxes = true
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
func (c *Client) SearchDocuments(opts scheme.DocumentSearchOptions) (*scheme.Documents, error) {
	out := new(*scheme.Documents)
	if err := c.get(documentURI, opts, out); err != nil {
		return nil, err
	}

	return *out, nil
}

// SearchDetailedDocuments returns a list of processed documents with matching queries.
func (c *Client) SearchDetailedDocuments(opts scheme.DocumentSearchOptions) (*scheme.DetailedDocuments, error) {
	out := new(*scheme.DetailedDocuments)
	detailedOpts := scheme.DetailedDocumentSearchOptions{
		Q:                 opts.Q,
		ExternalID:        opts.ExternalID,
		Tag:               opts.Tag,
		CreatedGT:         opts.CreatedGT,
		CreatedGTE:        opts.CreatedGTE,
		CreatedLT:         opts.CreatedLT,
		CreatedLTE:        opts.CreatedLTE,
		Status:            opts.Status,
		DeviceID:          opts.DeviceID,
		Owner:             opts.Owner,
		UpdatedGT:         opts.UpdatedGT,
		UpdatedGTE:        opts.UpdatedGTE,
		UpdatedLT:         opts.UpdatedLT,
		UpdatedLTE:        opts.UpdatedLTE,
		DateGT:            opts.DateGT,
		DateGTE:           opts.DateGTE,
		DateLT:            opts.DateLT,
		DateLTE:           opts.DateLTE,
		Page:              opts.Page,
		PageSize:          opts.PageSize,
		TrackTotalResults: opts.TrackTotalResults,
		BoundingBoxes:     true,
		ConfidenceDetails: true,
	}
	if err := c.get(documentURI, detailedOpts, out); err != nil {
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

// GetDetailedDocument returns a processed document with detailed field information
func (c *Client) GetDetailedDocument(documentID string, opts scheme.DocumentGetOptions) (*scheme.DetailedDocument, error) {
	out := new(*scheme.DetailedDocument)
	detailedOpts := scheme.DocumentGetDetailedOptions{
		ReturnAuditTrail:  opts.ReturnAuditTrail,
		ConfidenceDetails: true,
		BoundingBoxes:     true,
	}
	err := c.get(fmt.Sprintf("%s%s", documentURI, documentID), detailedOpts, out)
	if err != nil {
		return nil, err
	}
	return *out, nil
}

// request returns an authorized request to Veryfi API.
func (c *Client) request(payload interface{}, okScheme interface{}, errScheme interface{}) *resty.Request {
	timestamp := int(time.Now().Unix())
	return c.setBaseURL().R().
		SetHeaders(map[string]string{
			"User-Agent":                 fmt.Sprintf("Go Veryfi-Go/%s", c.pkgVersion),
			"Content-Type":               "application/json",
			"Accept":                     "application/json",
			"Client-Id":                  c.options.ClientID,
			"Authorization":              fmt.Sprintf("apikey %s:%s", c.options.Username, c.options.APIKey),
			"X-Veryfi-Request-Timestamp": strconv.Itoa(timestamp),
			"X-Veryfi-Request-Signature": c.generateSignature(payload, timestamp),
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
	request := c.request(body, okScheme, errScheme).SetBody(body)

	_, err := request.Post(uri)

	return check(err, errScheme)
}

// put performs a PUT request against Veryfi API.
func (c *Client) put(uri string, body interface{}, okScheme interface{}) error {
	errScheme := new(scheme.Error)
	request := c.request(body, okScheme, errScheme).SetBody(body)
	_, err := request.Put(uri)

	return check(err, errScheme)
}

// get performs a GET request against Veryfi API.
func (c *Client) get(uri string, queryParams interface{}, okScheme interface{}) error {
	errScheme := new(scheme.Error)
	request := c.request(queryParams, okScheme, errScheme)
	if queryParams != nil {
		request.SetQueryParams(structToMap(queryParams))
	}

	_, err := request.Get(uri)

	return check(err, errScheme)
}

// rdelete performs a DELETE request against Veryfi API.
func (c *Client) rdelete(uri string) error {
	errScheme := new(scheme.Error)
	request := c.request(struct{}{}, map[string]string{}, errScheme)
	_, err := request.Delete(uri)

	return check(err, errScheme)
}

// generateSignature for a given request.
func (c *Client) generateSignature(s interface{}, timestamp int) string {
	p := []string{fmt.Sprintf("timestamp:%v", timestamp)}
	for k, v := range structToMap(s) {
		p = append(p, fmt.Sprintf("%v:%v", k, v))
	}

	h := hmac.New(sha256.New, []byte(c.options.ClientSecret))
	h.Write([]byte(strings.Join(p, ",")))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// check validates returned response from Veryfi.
func check(err error, errResp *scheme.Error) error {
	if err != nil {
		return errors.Wrap(err, "fail to make a request to Veryfi")
	}

	// HTTP errors are handled by OnAfterResponse, but we still check for
	// any structured errors that might not have triggered HTTP error status codes
	if errResp != nil && *errResp != (scheme.Error{}) {
		ctx := ""
		if len(errResp.Error) > 0 {
			ctx = errResp.Error
		}
		if errResp.Details != nil {
			ctx = fmt.Sprintf("%v", errResp.Details)
		}

		return errors.Errorf(
			"get a response from Veryfi with status=%s and context=%s",
			errResp.Status, ctx,
		)
	}

	return nil
}
