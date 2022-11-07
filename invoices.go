package xero

import (
	"context"
	"io"
	"net/http"
	"time"
)

type InvoicesService service

// InvoiceListOptions specifies the optional parameters to the InvoicesService.GetInvoices.
type InvoiceListOptions struct {
	Id             []string `url:"IDs,omitempty,comma"`
	Status         []string `url:"Statuses,omitempty,comma"`
	ContactIDs     []string `url:"ContactIDs,omitempty,comma"`
	InvoiceNumbers []string `url:"InvoiceNumbers,omitempty,comma"`

	ModifiedAfter time.Time `url:"-"`

	SummaryOnly bool `url:"summaryOnly"`

	// For paginated result sets, page of results to retrieve.
	Page int `url:"page,omitempty"`

	ListOptions
}

func (s *InvoicesService) GetInvoices(ctx context.Context, opts *InvoiceListOptions) (*InvoicesResponse, *http.Response, error) {
	u := "api.xro/2.0/invoices"
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	addModifiedSinceHeader(req, opts.ModifiedAfter)

	var c *InvoicesResponse
	resp, err := s.client.Do(ctx, req, &c)
	if err != nil {
		return nil, nil, err
	}

	return c, resp, nil
}

func (s *InvoicesService) GetInvoiceHistory(ctx context.Context, id string) (*HistoryRecords, *http.Response, error) {
	return getHistory(ctx, s.client, "invoices", id)
}

func (s *InvoicesService) AddInvoiceHistoryRecord(ctx context.Context, id string, records *HistoryRecords) (*http.Response, error) {
	return addHistory(ctx, s.client, "invoices", id, records)
}

func (s *InvoicesService) CreateInvoice(ctx context.Context, inv *Invoice) (*InvoicesResponse, *http.Response, error) {
	req, err := s.client.NewRequest("PUT", "api.xro/2.0/invoices", inv)
	if err != nil {
		return nil, nil, err
	}

	var c *InvoicesResponse
	resp, err := s.client.Do(ctx, req, &c)
	if err != nil {
		return nil, nil, err
	}

	return c, resp, nil
}

func (s *InvoicesService) UpdateInvoice(ctx context.Context, inv *Invoice) (*InvoicesResponse, *http.Response, error) {
	req, err := s.client.NewRequest("POST", "api.xro/2.0/invoices", inv)
	if err != nil {
		return nil, nil, err
	}

	var c *InvoicesResponse
	resp, err := s.client.Do(ctx, req, &c)
	if err != nil {
		return nil, nil, err
	}

	return c, resp, nil
}

func (s *InvoicesService) CreateInvoices(ctx context.Context, inv []*Invoice) (*InvoicesResponse, *http.Response, error) {
	batch := &InvoicesBatch{Invoices: inv}

	req, err := s.client.NewRequest("PUT", "api.xro/2.0/invoices?summarizeErrors=false", batch)
	if err != nil {
		return nil, nil, err
	}

	var c *InvoicesResponse
	resp, err := s.client.Do(ctx, req, &c)
	if err != nil {
		return nil, nil, err
	}

	return c, resp, nil
}

func (s *InvoicesService) UpdateInvoices(ctx context.Context, inv []*Invoice) (*InvoicesResponse, *http.Response, error) {
	batch := &InvoicesBatch{Invoices: inv}

	req, err := s.client.NewRequest("POST", "api.xro/2.0/invoices?summarizeErrors=false", batch)
	if err != nil {
		return nil, nil, err
	}

	var c *InvoicesResponse
	resp, err := s.client.Do(ctx, req, &c)
	if err != nil {
		return nil, nil, err
	}

	return c, resp, nil
}

func (s *InvoicesService) GetAttachments(ctx context.Context, id string) (*AttachmentsResponse, *http.Response, error) {
	return getAttachments(ctx, s.client, "invoices", id)
}

func (s *InvoicesService) GetAttachment(ctx context.Context, id, filename string, w io.Writer) (*http.Response, error) {
	return getAttachment(ctx, s.client, "invoices", id, filename, w)
}

func (s *InvoicesService) UploadAttachment(ctx context.Context, id, filename string, includeOnline bool, r io.Reader) (*http.Response, error) {
	return uploadAttachment(ctx, s.client, "invoices", id, filename, includeOnline, r)
}

type InvoicesResponse struct {
	Invoices    []*Invoice `json:"Invoices"`
	Elements    []*Invoice `json:"Elements"`
	Status      string     `json:"Status"`
	ErrorNumber *int       `json:"ErrorNumber,omitempty"`
}

type Invoice struct {
	Type            *string  `json:"Type,omitempty"`
	InvoiceID       *string  `json:"InvoiceID,omitempty"`
	InvoiceNumber   *string  `json:"InvoiceNumber,omitempty"`
	Reference       *string  `json:"Reference,omitempty"`
	AmountDue       *float64 `json:"AmountDue,omitempty"`
	AmountPaid      *float64 `json:"AmountPaid,omitempty"`
	AmountCredited  *float64 `json:"AmountCredited,omitempty"`
	SubTotal        *float64 `json:"SubTotal,omitempty"`
	TotalTax        *float64 `json:"TotalTax,omitempty"`
	Total           *float64 `json:"Total,omitempty"`
	SentToContact   *bool    `json:"SentToContact,omitempty"`
	Status          *string  `json:"Status,omitempty"`
	LineAmountTypes *string  `json:"LineAmountTypes,omitempty"`
	CurrencyCode    *string  `json:"CurrencyCode,omitempty"`
	Date            *NetDate `json:"Date,omitempty"` // "/Date(1496361600000+0000)/"
	DateString      *string  `json:"DateString,omitempty"`
	DueDate         *NetDate `json:"DueDate,omitempty"` // "/Date(1496620800000+0000)/"
	DueDateString   *string  `json:"DueDateString,omitempty"`
	UpdatedDateUTC  *NetDate `json:"UpdatedDateUTC,omitempty"` // "/Date(1496620800000+0000)/"
	BrandingThemeID *string  `json:"BrandingThemeID,omitempty"`

	Contact   *Contact    `json:"Contact,omitempty"`
	LineItems []*LineItem `json:"LineItems,omitempty"`
	Payments  []*Payment  `json:"Payments,omitempty"`

	ValidationErrors      []*ValidationError `json:"ValidationErrors,omitempty"`
	StatusAttributeString string             `json:"StatusAttributeString,omitempty"`
}

type LineItem struct {
	LineItemID       *string            `json:"LineItemID,omitempty"`
	Description      *string            `json:"Description,omitempty"`
	Quantity         *float64           `json:"Quantity,omitempty"`
	UnitAmount       *float64           `json:"UnitAmount,omitempty"`
	AccountCode      *string            `json:"AccountCode,omitempty"`
	TaxType          *string            `json:"TaxType,omitempty"`
	TaxAmount        *float64           `json:"TaxAmount,omitempty"`
	LineAmount       *float64           `json:"LineAmount,omitempty"`
	DiscountRate     *float64           `json:"DiscountRate,omitempty"`
	DiscountAmount   *float64           `json:"DiscountAmount,omitempty"`
	Tracking         []*Tracking        `json:"Tracking,omitempty"`
	ValidationErrors []*ValidationError `json:"ValidationErrors,omitempty"`
}

type Tracking struct {
	TrackingCategoryId *string `json:"TrackingCategoryID,omitempty"`
	Name               *string `json:"Name,omitempty"`
	Option             *string `json:"Option,omitempty"`
}

type InvoicesBatch struct {
	Invoices []*Invoice `json:"Invoices"`
}
