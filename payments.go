package xero

import (
	"context"
	"net/http"
	"time"
)

type PaymentsService service

type PaymentListOptions struct {
	ModifiedAfter time.Time `url:"-"`

	SummaryOnly bool `url:"summaryOnly"`

	// For paginated result sets, page of results to retrieve.
	Page int `url:"page,omitempty"`

	ListOptions
}

func (s *PaymentsService) GetPayments(ctx context.Context, opts *PaymentListOptions) (*PaymentsResponse, *http.Response, error) {
	u := "api.xro/2.0/payments"
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	addModifiedSinceHeader(req, opts.ModifiedAfter)

	var c *PaymentsResponse
	resp, err := s.client.Do(ctx, req, &c)
	if err != nil {
		return nil, nil, err
	}

	return c, resp, nil
}

func (s *PaymentsService) GetPayment(ctx context.Context, id string) (*PaymentsResponse, *http.Response, error) {
	u := "api.xro/2.0/payments/" + id

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var c *PaymentsResponse
	resp, err := s.client.Do(ctx, req, &c)
	if err != nil {
		return nil, nil, err
	}

	return c, resp, nil
}

func (s *PaymentsService) CreatePayments(ctx context.Context, payments []*Payment) (*PaymentsResponse, *http.Response, error) {
	batch := &PaymentsBatch{Payments: payments}

	req, err := s.client.NewRequest("PUT", "api.xro/2.0/payments?summarizeErrors=false", batch)
	if err != nil {
		return nil, nil, err
	}

	var c *PaymentsResponse
	resp, err := s.client.Do(ctx, req, &c)
	if err != nil {
		return nil, nil, err
	}

	return c, resp, nil
}

func (s *PaymentsService) DeletePayment(ctx context.Context, id string) (*http.Response, error) {
	u := "api.xro/2.0/payments/" + id

	payload := &Payment{
		Status: String("DELETED"),
	}

	req, err := s.client.NewRequest("POST", u, payload)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *PaymentsService) UpdatePayment(ctx context.Context, payment *Payment) (*http.Response, error) {
	u := "api.xro/2.0/payments/" + *payment.PaymentId

	req, err := s.client.NewRequest("PUT", u, payment)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type PaymentsResponse struct {
	Payments []*Payment `json:"Payments,omitempty"`
	Status   string     `json:"Status,omitempty"`
}
type Payment struct {
	PaymentId             *string            `json:"PaymentID,omitempty"`
	BankAccountNumber     *string            `json:"BankAccountNumber,omitempty"`
	Particulars           *string            `json:"Particulars,omitempty"`
	Code                  *string            `json:"Code,omitempty"`
	Reference             *string            `json:"Reference,omitempty"`
	Details               *string            `json:"Details,omitempty"`
	Amount                *float64           `json:"Amount,omitempty"`
	Date                  *NetDate           `json:"Date,omitempty"`
	IsReconciled          *bool              `json:"IsReconciled,omitempty"`
	Status                *string            `json:"Status,omitempty"`
	PaymentType           *string            `json:"PaymentType,omitempty"`
	UpdatedDateUTC        *NetDate           `json:"UpdatedDateUTC,omitempty"`
	BatchPaymentId        *string            `json:"BatchPaymentID,omitempty"`
	BatchPayment          *BatchPayment      `json:"BatchPayment,omitempty"`
	Account               *Account           `json:"Account,omitempty"`
	Invoice               *Invoice           `json:"Invoice,omitempty"`
	CurrencyRate          *float64           `json:"CurrencyRate,omitempty"`
	HasAccount            *bool              `json:"HasAccount,omitempty"`
	HasValidationErrors   *bool              `json:"HasValidationErrors,omitempty"`
	ValidationErrors      []*ValidationError `json:"ValidationErrors,omitempty"`
	StatusAttributeString string             `json:"StatusAttributeString,omitempty"`
}

type PaymentsBatch struct {
	Payments []*Payment `json:"Payments"`
}
