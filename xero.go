package xero

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/google/go-querystring/query"
)

var errNonNilContext = errors.New("context must be non-nil")

var BaseUrl = "https://api.xero.com/"

type Client struct {
	client *http.Client

	// BaseURL should always be specified with a trailing slash.
	BaseURL  *url.URL
	TenantId string

	common service

	Tenants            *TenantsService
	Connections        *ConnectionsService
	Accounts           *AccountsService
	Invoices           *InvoicesService
	Contacts           *ContactsService
	ContactGroups      *ContactGroupsService
	CreditNotes        *CreditNotesService
	Payments           *PaymentsService
	TrackingCategories *TrackingCategoriesService
	BrandingThemes     *BrandingThemesService
}

type service struct {
	client *Client
}

// addOptions adds the parameters in opts as URL query parameters to s. opts
// must be a struct whose fields may contain "url" tags.
func addOptions(s string, opts interface{}) (string, error) {
	v := reflect.ValueOf(opts)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(opts)
	if err != nil {
		return s, err
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}

func addModifiedSinceHeader(r *http.Request, t time.Time) {
	if t.IsZero() {
		return
	}

	r.Header.Set("If-Modified-Since", t.Format("2006-01-02T15:04:05"))
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	baseURL, _ := url.Parse(BaseUrl)

	c := &Client{client: httpClient, BaseURL: baseURL}
	c.common.client = c

	c.Tenants = (*TenantsService)(&c.common)
	c.Connections = (*ConnectionsService)(&c.common)
	c.Accounts = (*AccountsService)(&c.common)
	c.Invoices = (*InvoicesService)(&c.common)
	c.Contacts = (*ContactsService)(&c.common)
	c.ContactGroups = (*ContactGroupsService)(&c.common)
	c.CreditNotes = (*CreditNotesService)(&c.common)
	c.Payments = (*PaymentsService)(&c.common)
	c.TrackingCategories = (*TrackingCategoriesService)(&c.common)
	c.BrandingThemes = (*BrandingThemesService)(&c.common)

	return c
}

func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.BaseURL)
	}
	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	contentType := ""

	var b io.Reader
	if body != nil {
		switch body.(type) {
		case io.Reader:
			b = body.(io.Reader)
		default:
			var buf io.ReadWriter
			buf = &bytes.Buffer{}
			enc := json.NewEncoder(buf)
			enc.SetEscapeHTML(false)
			err := enc.Encode(body)
			if err != nil {
				return nil, err
			}
			b = buf
			contentType = "application/json"
		}
	}

	req, err := http.NewRequest(method, u.String(), b)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	if c.TenantId != "" {
		req.Header.Set("Xero-tenant-id", c.TenantId)
	}

	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}
	return req, nil
}

func (c *Client) BareDo(ctx context.Context, req *http.Request) (*http.Response, error) {
	if ctx == nil {
		return nil, errNonNilContext
	}
	req = req.WithContext(ctx)

	resp, err := c.client.Do(req)
	if err != nil {
		// If we got an error, and the context has been canceled,
		// the context's error is probably more useful.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		return nil, err
	}

	err = CheckResponse(resp)
	return resp, err
}

func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.BareDo(ctx, req)
	if err != nil {
		return resp, err
	}
	defer resp.Body.Close()

	switch v := v.(type) {
	case nil:
	case io.Writer:
		_, err = io.Copy(v, resp.Body)
	default:
		decErr := json.NewDecoder(resp.Body).Decode(v)
		if decErr == io.EOF {
			decErr = nil // ignore EOF errors caused by empty response body
		}
		if decErr != nil {
			err = decErr
		}
	}
	return resp, err
}

type ErrorResponse struct {
	Response    *http.Response // HTTP response that caused this error
	Limits      *Limits
	ErrorNumber int              `json:"ErrorNumber"`
	Type        string           `json:"Type"`
	Message     string           `json:"Message"`
	Elements    []*ErrorElements `json:"Elements"`
}

type ErrorElements struct {
	ValidationErrors []*ValidationError `json:"ValidationErrors"`
}

type ValidationError struct {
	Message string `json:"Message"`
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v %d: %v %v - %v",
		r.Response.Request.Method, r.Response.Request.URL, r.Response.StatusCode,
		r.ErrorNumber, r.Type, r.Message)
}

func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}
	errorResponse := &ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		json.Unmarshal(data, errorResponse)
	}

	errorResponse.Limits = GetLimits(r)

	// Re-populate error response body
	r.Body = ioutil.NopCloser(bytes.NewBuffer(data))

	return errorResponse
}

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *bool { return &v }

// Int is a helper routine that allocates a new int value
// to store v and returns a pointer to it.
func Int(v int) *int { return &v }

// Int64 is a helper routine that allocates a new int64 value
// to store v and returns a pointer to it.
func Int64(v int64) *int64 { return &v }

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string { return &v }

func Float64(v float64) *float64 { return &v }

// ListOptions specifies the optional parameters to various List methods that
// support offset pagination.
type ListOptions struct {
	Where string `url:"where,omitempty"`
	Order string `url:"order,omitempty"`
}

func strToInt(val string) (int, bool) {
	if i, err := strconv.Atoi(val); err != nil {
		return i, false
	} else {
		return i, true
	}
}

type Limits struct {
	dayLimitRemaining       int
	minuteLimitRemaining    int
	appMinuteLimitRemaining int
	retryAfter              int
}

func (l *Limits) DayLimitRemaining() int {
	return l.dayLimitRemaining
}

func (l *Limits) MinuteLimitRemaining() int {
	return l.minuteLimitRemaining
}

func (l *Limits) AppMinuteLimitRemaining() int {
	return l.appMinuteLimitRemaining
}

func (l *Limits) RetryAfterSeconds() int {
	return l.retryAfter
}

func (l *Limits) RetryAfter() time.Duration {
	return time.Duration(l.RetryAfterSeconds()) * time.Second
}

func GetLimits(res *http.Response) *Limits {
	limits := &Limits{}

	if dayLimitRemaining, ok := strToInt(res.Header.Get("X-DayLimit-Remaining")); ok {
		limits.dayLimitRemaining = dayLimitRemaining
	}

	if minuteLimitRemaining, ok := strToInt(res.Header.Get("X-MinLimit-Remaining")); ok {
		limits.minuteLimitRemaining = minuteLimitRemaining
	}

	if appMinuteLimitRemaining, ok := strToInt(res.Header.Get("X-AppMinLimit-Remaining")); ok {
		limits.appMinuteLimitRemaining = appMinuteLimitRemaining
	}

	if res.StatusCode == 429 {
		if retryAfter, ok := strToInt(res.Header.Get("Retry-After")); ok {
			limits.retryAfter = retryAfter
		}
	} else {
		limits.retryAfter = 0
	}

	return limits
}
