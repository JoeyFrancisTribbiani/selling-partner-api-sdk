// Package listings provides primitives to interact the openapi HTTP API.
//
// Code generated by go-sdk-codegen DO NOT EDIT.
package listings

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	runt "runtime"
	"strings"

	"github.com/JoeyFrancisTribbiani/selling-partner-api-sdk/pkg/runtime"
)

// RequestBeforeFn  is the function signature for the RequestBefore callback function
type RequestBeforeFn func(ctx context.Context, req *http.Request) error

// ResponseAfterFn  is the function signature for the ResponseAfter callback function
type ResponseAfterFn func(ctx context.Context, rsp *http.Response) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Endpoint string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A callback for modifying requests which are generated before sending over
	// the network.
	RequestBefore RequestBeforeFn

	// A callback for modifying response which are generated before sending over
	// the network.
	ResponseAfter ResponseAfterFn

	// The user agent header identifies your application, its version number, and the platform and programming language you are using.
	// You must include a user agent header in each request submitted to the sales partner API.
	UserAgent string
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(endpoint string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Endpoint: endpoint,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the endpoint URL always has a trailing slash
	if !strings.HasSuffix(client.Endpoint, "/") {
		client.Endpoint += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = http.DefaultClient
	}
	// setting the default useragent
	if client.UserAgent == "" {
		client.UserAgent = fmt.Sprintf("selling-partner-api-sdk/v1.0 (Language=%s; Platform=%s-%s)", strings.Replace(runt.Version(), "go", "go/", -1), runt.GOOS, runt.GOARCH)
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithUserAgent set up useragent
// add user agent to every request automatically
func WithUserAgent(userAgent string) ClientOption {
	return func(c *Client) error {
		c.UserAgent = userAgent
		return nil
	}
}

// WithRequestBefore allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestBefore(fn RequestBeforeFn) ClientOption {
	return func(c *Client) error {
		c.RequestBefore = fn
		return nil
	}
}

// WithResponseAfter allows setting up a callback function, which will be
// called right after get response the request. This can be used to log.
func WithResponseAfter(fn ResponseAfterFn) ClientOption {
	return func(c *Client) error {
		c.ResponseAfter = fn
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// DeleteListingsItem request
	DeleteListingsItem(ctx context.Context, sellerId string, sku string, params *DeleteListingsItemParams) (*http.Response, error)

	// GetListingsItem request
	GetListingsItem(ctx context.Context, sellerId string, sku string, params *GetListingsItemParams) (*http.Response, error)

	// PatchListingsItem request  with any body
	PatchListingsItemWithBody(ctx context.Context, sellerId string, sku string, params *PatchListingsItemParams, contentType string, body io.Reader) (*http.Response, error)

	PatchListingsItem(ctx context.Context, sellerId string, sku string, params *PatchListingsItemParams, body PatchListingsItemJSONRequestBody) (*http.Response, error)

	// PutListingsItem request  with any body
	PutListingsItemWithBody(ctx context.Context, sellerId string, sku string, params *PutListingsItemParams, contentType string, body io.Reader) (*http.Response, error)

	PutListingsItem(ctx context.Context, sellerId string, sku string, params *PutListingsItemParams, body PutListingsItemJSONRequestBody) (*http.Response, error)
}

func (c *Client) DeleteListingsItem(ctx context.Context, sellerId string, sku string, params *DeleteListingsItemParams) (*http.Response, error) {
	req, err := NewDeleteListingsItemRequest(c.Endpoint, sellerId, sku, params)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.ResponseAfter != nil {
		err = c.ResponseAfter(ctx, rsp)
		if err != nil {
			return nil, err
		}
	}
	return rsp, nil
}

func (c *Client) GetListingsItem(ctx context.Context, sellerId string, sku string, params *GetListingsItemParams) (*http.Response, error) {
	req, err := NewGetListingsItemRequest(c.Endpoint, sellerId, sku, params)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.ResponseAfter != nil {
		err = c.ResponseAfter(ctx, rsp)
		if err != nil {
			return nil, err
		}
	}
	return rsp, nil
}

func (c *Client) PatchListingsItemWithBody(ctx context.Context, sellerId string, sku string, params *PatchListingsItemParams, contentType string, body io.Reader) (*http.Response, error) {
	req, err := NewPatchListingsItemRequestWithBody(c.Endpoint, sellerId, sku, params, contentType, body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.ResponseAfter != nil {
		err = c.ResponseAfter(ctx, rsp)
		if err != nil {
			return nil, err
		}
	}
	return rsp, nil
}

func (c *Client) PatchListingsItem(ctx context.Context, sellerId string, sku string, params *PatchListingsItemParams, body PatchListingsItemJSONRequestBody) (*http.Response, error) {
	req, err := NewPatchListingsItemRequest(c.Endpoint, sellerId, sku, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) PutListingsItemWithBody(ctx context.Context, sellerId string, sku string, params *PutListingsItemParams, contentType string, body io.Reader) (*http.Response, error) {
	req, err := NewPutListingsItemRequestWithBody(c.Endpoint, sellerId, sku, params, contentType, body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.ResponseAfter != nil {
		err = c.ResponseAfter(ctx, rsp)
		if err != nil {
			return nil, err
		}
	}
	return rsp, nil
}

func (c *Client) PutListingsItem(ctx context.Context, sellerId string, sku string, params *PutListingsItemParams, body PutListingsItemJSONRequestBody) (*http.Response, error) {
	req, err := NewPutListingsItemRequest(c.Endpoint, sellerId, sku, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

// NewDeleteListingsItemRequest generates requests for DeleteListingsItem
func NewDeleteListingsItemRequest(endpoint string, sellerId string, sku string, params *DeleteListingsItemParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParam("simple", false, "sellerId", sellerId)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParam("simple", false, "sku", sku)
	if err != nil {
		return nil, err
	}

	queryUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/listings/2021-08-01/items/%s/%s", pathParam0, pathParam1)
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	queryValues := queryUrl.Query()

	if queryFrag, err := runtime.StyleParam("form", false, "marketplaceIds", params.MarketplaceIds); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if params.IssueLocale != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "issueLocale", *params.IssueLocale); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryUrl.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("DELETE", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetListingsItemRequest generates requests for GetListingsItem
func NewGetListingsItemRequest(endpoint string, sellerId string, sku string, params *GetListingsItemParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParam("simple", false, "sellerId", sellerId)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParam("simple", false, "sku", sku)
	if err != nil {
		return nil, err
	}

	queryUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/listings/2021-08-01/items/%s/%s", pathParam0, pathParam1)
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	queryValues := queryUrl.Query()

	if queryFrag, err := runtime.StyleParam("form", false, "marketplaceIds", params.MarketplaceIds); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if params.IssueLocale != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "issueLocale", *params.IssueLocale); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.IncludedData != nil {

		if queryFrag, err := runtime.StyleParam("form", false, "includedData", *params.IncludedData); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryUrl.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewPatchListingsItemRequest calls the generic PatchListingsItem builder with application/json body
func NewPatchListingsItemRequest(endpoint string, sellerId string, sku string, params *PatchListingsItemParams, body PatchListingsItemJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPatchListingsItemRequestWithBody(endpoint, sellerId, sku, params, "application/json", bodyReader)
}

// NewPatchListingsItemRequestWithBody generates requests for PatchListingsItem with any type of body
func NewPatchListingsItemRequestWithBody(endpoint string, sellerId string, sku string, params *PatchListingsItemParams, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParam("simple", false, "sellerId", sellerId)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParam("simple", false, "sku", sku)
	if err != nil {
		return nil, err
	}

	queryUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/listings/2021-08-01/items/%s/%s", pathParam0, pathParam1)
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	queryValues := queryUrl.Query()

	if queryFrag, err := runtime.StyleParam("form", false, "marketplaceIds", params.MarketplaceIds); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if params.IssueLocale != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "issueLocale", *params.IssueLocale); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryUrl.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("PATCH", queryUrl.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)
	return req, nil
}

// NewPutListingsItemRequest calls the generic PutListingsItem builder with application/json body
func NewPutListingsItemRequest(endpoint string, sellerId string, sku string, params *PutListingsItemParams, body PutListingsItemJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPutListingsItemRequestWithBody(endpoint, sellerId, sku, params, "application/json", bodyReader)
}

// NewPutListingsItemRequestWithBody generates requests for PutListingsItem with any type of body
func NewPutListingsItemRequestWithBody(endpoint string, sellerId string, sku string, params *PutListingsItemParams, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParam("simple", false, "sellerId", sellerId)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParam("simple", false, "sku", sku)
	if err != nil {
		return nil, err
	}

	queryUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/listings/2021-08-01/items/%s/%s", pathParam0, pathParam1)
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	queryValues := queryUrl.Query()

	if queryFrag, err := runtime.StyleParam("form", false, "marketplaceIds", params.MarketplaceIds); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if params.IssueLocale != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "issueLocale", *params.IssueLocale); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryUrl.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("PUT", queryUrl.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)
	return req, nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(endpoint string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(endpoint, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Endpoint = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// DeleteListingsItem request
	DeleteListingsItemWithResponse(ctx context.Context, sellerId string, sku string, params *DeleteListingsItemParams) (*DeleteListingsItemResp, error)

	// GetListingsItem request
	GetListingsItemWithResponse(ctx context.Context, sellerId string, sku string, params *GetListingsItemParams) (*GetListingsItemResp, error)

	// PatchListingsItem request  with any body
	PatchListingsItemWithBodyWithResponse(ctx context.Context, sellerId string, sku string, params *PatchListingsItemParams, contentType string, body io.Reader) (*PatchListingsItemResp, error)

	PatchListingsItemWithResponse(ctx context.Context, sellerId string, sku string, params *PatchListingsItemParams, body PatchListingsItemJSONRequestBody) (*PatchListingsItemResp, error)

	// PutListingsItem request  with any body
	PutListingsItemWithBodyWithResponse(ctx context.Context, sellerId string, sku string, params *PutListingsItemParams, contentType string, body io.Reader) (*PutListingsItemResp, error)

	PutListingsItemWithResponse(ctx context.Context, sellerId string, sku string, params *PutListingsItemParams, body PutListingsItemJSONRequestBody) (*PutListingsItemResp, error)
}

type DeleteListingsItemResp struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ListingsItemSubmissionResponse
	JSON400      *ErrorList
	JSON403      *ErrorList
	JSON413      *ErrorList
	JSON415      *ErrorList
	JSON429      *ErrorList
	JSON500      *ErrorList
	JSON503      *ErrorList
}

// Status returns HTTPResponse.Status
func (r DeleteListingsItemResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DeleteListingsItemResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetListingsItemResp struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Item
	JSON400      *ErrorList
	JSON403      *ErrorList
	JSON404      *ErrorList
	JSON413      *ErrorList
	JSON415      *ErrorList
	JSON429      *ErrorList
	JSON500      *ErrorList
	JSON503      *ErrorList
}

// Status returns HTTPResponse.Status
func (r GetListingsItemResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetListingsItemResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PatchListingsItemResp struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ListingsItemSubmissionResponse
	JSON400      *ErrorList
	JSON403      *ErrorList
	JSON413      *ErrorList
	JSON415      *ErrorList
	JSON429      *ErrorList
	JSON500      *ErrorList
	JSON503      *ErrorList
}

// Status returns HTTPResponse.Status
func (r PatchListingsItemResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PatchListingsItemResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PutListingsItemResp struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ListingsItemSubmissionResponse
	JSON400      *ErrorList
	JSON403      *ErrorList
	JSON413      *ErrorList
	JSON415      *ErrorList
	JSON429      *ErrorList
	JSON500      *ErrorList
	JSON503      *ErrorList
}

// Status returns HTTPResponse.Status
func (r PutListingsItemResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PutListingsItemResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// DeleteListingsItemWithResponse request returning *DeleteListingsItemResponse
func (c *ClientWithResponses) DeleteListingsItemWithResponse(ctx context.Context, sellerId string, sku string, params *DeleteListingsItemParams) (*DeleteListingsItemResp, error) {
	rsp, err := c.DeleteListingsItem(ctx, sellerId, sku, params)
	if err != nil {
		return nil, err
	}
	return ParseDeleteListingsItemResp(rsp)
}

// GetListingsItemWithResponse request returning *GetListingsItemResponse
func (c *ClientWithResponses) GetListingsItemWithResponse(ctx context.Context, sellerId string, sku string, params *GetListingsItemParams) (*GetListingsItemResp, error) {
	rsp, err := c.GetListingsItem(ctx, sellerId, sku, params)
	if err != nil {
		return nil, err
	}
	return ParseGetListingsItemResp(rsp)
}

// PatchListingsItemWithBodyWithResponse request with arbitrary body returning *PatchListingsItemResponse
func (c *ClientWithResponses) PatchListingsItemWithBodyWithResponse(ctx context.Context, sellerId string, sku string, params *PatchListingsItemParams, contentType string, body io.Reader) (*PatchListingsItemResp, error) {
	rsp, err := c.PatchListingsItemWithBody(ctx, sellerId, sku, params, contentType, body)
	if err != nil {
		return nil, err
	}
	return ParsePatchListingsItemResp(rsp)
}

func (c *ClientWithResponses) PatchListingsItemWithResponse(ctx context.Context, sellerId string, sku string, params *PatchListingsItemParams, body PatchListingsItemJSONRequestBody) (*PatchListingsItemResp, error) {
	rsp, err := c.PatchListingsItem(ctx, sellerId, sku, params, body)
	if err != nil {
		return nil, err
	}
	return ParsePatchListingsItemResp(rsp)
}

// PutListingsItemWithBodyWithResponse request with arbitrary body returning *PutListingsItemResponse
func (c *ClientWithResponses) PutListingsItemWithBodyWithResponse(ctx context.Context, sellerId string, sku string, params *PutListingsItemParams, contentType string, body io.Reader) (*PutListingsItemResp, error) {
	rsp, err := c.PutListingsItemWithBody(ctx, sellerId, sku, params, contentType, body)
	if err != nil {
		return nil, err
	}
	return ParsePutListingsItemResp(rsp)
}

func (c *ClientWithResponses) PutListingsItemWithResponse(ctx context.Context, sellerId string, sku string, params *PutListingsItemParams, body PutListingsItemJSONRequestBody) (*PutListingsItemResp, error) {
	rsp, err := c.PutListingsItem(ctx, sellerId, sku, params, body)
	if err != nil {
		return nil, err
	}
	return ParsePutListingsItemResp(rsp)
}

// ParseDeleteListingsItemResp parses an HTTP response from a DeleteListingsItemWithResponse call
func ParseDeleteListingsItemResp(rsp *http.Response) (*DeleteListingsItemResp, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &DeleteListingsItemResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ListingsItemSubmissionResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 413:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON413 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 415:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON415 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 503:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON503 = &dest

	}

	return response, nil
}

// ParseGetListingsItemResp parses an HTTP response from a GetListingsItemWithResponse call
func ParseGetListingsItemResp(rsp *http.Response) (*GetListingsItemResp, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &GetListingsItemResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Item
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 413:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON413 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 415:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON415 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 503:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON503 = &dest

	}

	return response, nil
}

// ParsePatchListingsItemResp parses an HTTP response from a PatchListingsItemWithResponse call
func ParsePatchListingsItemResp(rsp *http.Response) (*PatchListingsItemResp, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &PatchListingsItemResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ListingsItemSubmissionResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 413:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON413 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 415:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON415 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 503:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON503 = &dest

	}

	return response, nil
}

// ParsePutListingsItemResp parses an HTTP response from a PutListingsItemWithResponse call
func ParsePutListingsItemResp(rsp *http.Response) (*PutListingsItemResp, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &PutListingsItemResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ListingsItemSubmissionResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 413:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON413 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 415:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON415 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 503:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON503 = &dest

	}

	return response, nil
}
