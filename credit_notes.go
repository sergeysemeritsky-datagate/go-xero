package xero

import (
	"context"
	"io"
	"net/http"
)

//TODO: add more methods
//TODO: test all mathods

type CreditNotesService service

func (s *CreditNotesService) CreateCreditNotes(ctx context.Context, notes []*CreditNote) (*CreditNotesResponse, *http.Response, error) {
	batch := &CreditNotesBatch{CreditNotes: notes}

	req, err := s.client.NewRequest("PUT", "api.xro/2.0/creditnotes?summarizeErrors=false", batch)
	if err != nil {
		return nil, nil, err
	}

	var c *CreditNotesResponse
	resp, err := s.client.Do(ctx, req, &c)
	if err != nil {
		return nil, nil, err
	}

	return c, resp, nil
}

func (s *CreditNotesService) GetAttachments(ctx context.Context, id string) (*AttachmentsResponse, *http.Response, error) {
	return getAttachments(ctx, s.client, "creditnotes", id)
}

func (s *CreditNotesService) GetAttachment(ctx context.Context, id, filename string, w io.Writer) (*http.Response, error) {
	return getAttachment(ctx, s.client, "creditnotes", id, filename, w)
}

func (s *CreditNotesService) UploadAttachment(ctx context.Context, id, filename string, includeOnline bool, r io.Reader) (*http.Response, error) {
	return uploadAttachment(ctx, s.client, "creditnotes", id, filename, includeOnline, r)
}

type CreditNotesResponse struct {
	CreditNotes []*CreditNote `json:"CreditNotes"`
	Elements    []*CreditNote `json:"Elements"`
	Status      string        `json:"Status"`
	ErrorNumber *int          `json:"ErrorNumber,omitempty"`
}

type CreditNote struct {
	Type             *string       `json:"Type,omitempty"`
	Contact          *Contact      `json:"Contact,omitempty"`
	Date             *NetDate      `json:"Date,omitempty"` // "/Date(1496361600000+0000)/"
	Status           *string       `json:"Status,omitempty"`
	LineAmountTypes  *string       `json:"LineAmountTypes,omitempty"`
	LineItems        []*LineItem   `json:"LineItems,omitempty"`
	SubTotal         *float64      `json:"SubTotal,omitempty"`
	TotalTax         *float64      `json:"TotalTax,omitempty"`
	Total            *float64      `json:"Total,omitempty"`
	UpdatedDateUTC   *NetDate      `json:"UpdatedDateUTC,omitempty"` // "/Date(1496620800000+0000)/"
	CurrencyCode     *string       `json:"CurrencyCode,omitempty"`
	FullyPaidOnDate  *NetDate      `json:"FullyPaidOnDate,omitempty"` // "/Date(1496620800000+0000)/"
	CreditNoteID     *string       `json:"CreditNoteID,omitempty"`
	CreditNoteNumber *string       `json:"CreditNoteNumber,omitempty"`
	Reference        *string       `json:"Reference,omitempty"`
	SentToContact    *bool         `json:"SentToContact,omitempty"`
	CurrencyRate     *float64      `json:"CurrencyRate,omitempty"`
	RemainingCredit  *float64      `json:"RemainingCredit,omitempty"`
	BrandingThemeID  *string       `json:"BrandingThemeID,omitempty"`
	Allocations      []*Allocation `json:"Allocations,omitempty"`

	ValidationErrors      []*ValidationError `json:"ValidationErrors,omitempty"`
	StatusAttributeString string             `json:"StatusAttributeString,omitempty"`
}

type CreditNotesBatch struct {
	CreditNotes []*CreditNote `json:"CreditNotes"`
}
