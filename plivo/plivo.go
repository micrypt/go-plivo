// Public Domain (-) 2013-2014 The GoPlivo Authors.
// See the GoPlivo UNLICENSE file for details.

package plivo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

const (
	libraryVersion = "0.1"
	defaultBaseURL = "https://api.plivo.com/%s/Account/"
	userAgent      = "go-plivo/" + libraryVersion
	apiVersion     = "v1"
)

// A client manages communication with the API.
type Client struct {
	// HTTP client used to communicate with the API.
	client *http.Client

	// Base URL for API requests. This should always be specified with the trailing slash.
	BaseURL *url.URL

	// User agent used when communicating the API.
	UserAgent string

	// Services used for talking to different parts of the API.
	Account     *AccountService
	Application *ApplicationService
	Call        *CallService
	Message     *MessageService
	Number      *NumberService
	Endpoint    *EndpointService
	Conference  *ConferenceService

	authID    string
	authToken string
}

// NewClient returns a new Plivo API client. If client is nil http.DefaultClient will be used
func NewClient(client *http.Client, authID, authToken string) *Client {
	baseURL, _ := url.Parse(fmt.Sprintf(defaultBaseURL, apiVersion))

	if client == nil {
		client = http.DefaultClient
	}

	c := &Client{client: client, BaseURL: baseURL, UserAgent: userAgent, authID: authID, authToken: authToken}
	c.Account = &AccountService{client: c}
	c.Application = &ApplicationService{client: c}
	c.Call = &CallService{client: c}
	c.Message = &MessageService{client: c}
	c.Number = &NumberService{client: c}
	c.Endpoint = &EndpointService{client: c}
	c.Conference = &ConferenceService{client: c}
	return c
}

// NewRequest creates an API request
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	buf := new(bytes.Buffer)
	params := ""
	if body != nil {
		if method == "GET" {
			v, err := query.Values(body)
			if err != nil {
				return nil, err
			}
			params = "?" + v.Encode()
		} else {
			err := json.NewEncoder(buf).Encode(body)
			if err != nil {
				return nil, err
			}
		}
	}

	req, err := http.NewRequest(method, u.String()+params, buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", c.UserAgent)
	req.SetBasicAuth(c.authID, c.authToken)

	return req, nil
}

// Meta contains response metadata. This is usually pagination information.
type Meta struct {
	Previous string
	Next     string

	TotalCount int64
	Offset     int64
	Limit      int64
}

// Response is a Plivo API response. This wraps the standard http.Response
// returned from Plivo while providing convenient access to pagination.
type Response struct {
	*http.Response

	*Meta
}

// newResponse intialise a Response
func newResponse(r *http.Response) *Response {
	response := &Response{Response: r}
	return response
}

// Do sends an API request and returns the API response. The response is returned as an error if one occurs
// or an attempt is made to decode it into v and the result of this operation returned if it fails.
func (c *Client) Do(req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	response := newResponse(resp)

	err = checkResponse(resp)
	if err != nil {
		// Even though there was an error, return the response so that the caller can inspect it.
		return response, err
	}

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
	}

	return response, err
}

// Errors returned by the Plivo API.
type ErrorResponse struct {
	Response *http.Response
	Message  string  `json:"message"`
	Errors   []Error `json:"errors"`
}

// Fetches the string representation of an ErrorResponse.
func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v %+v",
		r.Response.Request.Method, r.Response.Request.URL,
		r.Response.StatusCode, r.Message, r.Errors)
}

// Error type contains more details about the error.
type Error struct {
	Resource string `json:"resource"` // Resource on which the error was generated.
	Field    string `json:"field"`    // Field on which the error occurred.
	Code     string `json:"code"`     // Validation error code.
}

// Error returns the string representation of an Error.
func (e *Error) Error() string {
	return fmt.Sprintf("%v error caused by %v field on %v resource",
		e.Code, e.Field, e.Resource)
}

// checkResponse checks the API response for errors and returns them if present.
// A response if considered an error if it has a status code outside the 200 range.
func checkResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}
	errorResponse := &ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err != nil && data != nil {
		json.Unmarshal(data, errorResponse)
	}
	return errorResponse
}

// limitOffset is a utility type for handling offsets in the API calls.
type limitOffset struct {
	limit  int64
	offset int64
}
