package xero

import (
	"context"
	"net/http"
	"time"
)

type ContactsService service

type ContactListOptions struct {
	Id     []string `url:"IDs,omitempty,comma"`
	Status []string `url:"Statuses,omitempty,comma"`

	ModifiedAfter time.Time `url:"-"`

	SummaryOnly     bool `url:"summaryOnly"`
	IncludeArchived bool `url:"includeArchived"`

	// For paginated result sets, page of results to retrieve.
	Page int `url:"page,omitempty"`

	ListOptions
}

func (s *ContactsService) GetContacts(ctx context.Context, opts *ContactListOptions) (*ContactsResponse, *http.Response, error) {
	u := "api.xro/2.0/contacts"
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	addModifiedSinceHeader(req, opts.ModifiedAfter)

	var c *ContactsResponse
	resp, err := s.client.Do(ctx, req, &c)
	if err != nil {
		return nil, nil, err
	}

	return c, resp, nil
}

func (s *ContactsService) UpsertContacts(ctx context.Context, contacts []*Contact) (*ContactsResponse, *http.Response, error) {
	req, err := s.client.NewRequest("POST", "api.xro/2.0/contacts", &ContactsRequest{Contacts: contacts})
	if err != nil {
		return nil, nil, err
	}

	var c *ContactsResponse
	resp, err := s.client.Do(ctx, req, &c)
	if err != nil {
		return nil, nil, err
	}

	return c, resp, nil
}

type ContactsResponse struct {
	Contacts    []*Contact `json:"Contacts"`
	Status      string     `json:"Status"`
	ErrorNumber *int       `json:"ErrorNumber,omitempty"`
}

type ContactsRequest struct {
	Contacts []*Contact `json:"Contacts"`
}

type Contact struct {
	ContactId                 *string            `json:"ContactId,omitempty"`
	ContactStatus             *string            `json:"ContactStatus,omitempty"`
	Name                      *string            `json:"Name,omitempty"`
	FirstName                 *string            `json:"FirstName,omitempty"`
	LastName                  *string            `json:"LastName,omitempty"`
	CompanyNumber             *string            `json:"CompanyNumber,omitempty"`
	EmailAddress              *string            `json:"EmailAddress,omitempty"`
	SkypeUserName             *string            `json:"SkypeUserName,omitempty"`
	BankAccountDetails        *string            `json:"BankAccountDetails,omitempty"`
	TaxNumber                 *string            `json:"TaxNumber,omitempty"`
	AccountsReceivableTaxType *string            `json:"AccountsReceivableTaxType,omitempty"`
	AccountsPayableTaxType    *string            `json:"AccountsPayableTaxType,omitempty"`
	Addresses                 []*Address         `json:"Addresses,omitempty"`
	Phones                    []*Phone           `json:"Phones,omitempty"`
	UpdatedDateUTC            *NetDate           `json:"UpdatedDateUTC,omitempty"`
	ContactGroups             []*ContactGroup    `json:"ContactGroups,omitempty"`
	IsSupplier                *bool              `json:"IsSupplier,omitempty"`
	IsCustomer                *bool              `json:"IsCustomer,omitempty"`
	BrandingTheme             *BrandingTheme     `json:"BrandingTheme,omitempty"`
	PaymentTerms              *PaymentTerms      `json:"PaymentTerms,omitempty"`
	ContactPersons            []*ContactPerson   `json:"ContactPersons,omitempty"`
	DefaultCurrency           *string            `json:"DefaultCurrency,omitempty"`
	Balances                  *ContactBalance    `json:"Balances,omitempty"`
	HasAttachments            *bool              `json:"HasAttachments,omitempty"`
	HasValidationErrors       *bool              `json:"HasValidationErrors,omitempty"`
	ValidationErrors          []*ValidationError `json:"ValidationErrors,omitempty"`
}

type Address struct {
	AddressType  *string `json:"AddressType,omitempty"`
	AddressLine1 *string `json:"AddressLine1,omitempty"`
	AddressLine2 *string `json:"AddressLine2,omitempty"`
	AddressLine3 *string `json:"AddressLine3,omitempty"`
	AddressLine4 *string `json:"AddressLine4,omitempty"`
	City         *string `json:"City,omitempty"`
	Region       *string `json:"Region,omitempty"`
	PostalCode   *string `json:"PostalCode,omitempty"`
	Country      *string `json:"Country,omitempty"`
	AttentionTo  *string `json:"AttentionTo,omitempty"`
}

type Phone struct {
	PhoneType        *string `json:"PhoneType"`
	PhoneNumber      *string `json:"PhoneNumber,omitempty"`
	PhoneAreaCode    *string `json:"PhoneAreaCode,omitempty"`
	PhoneCountryCode *string `json:"PhoneCountryCode,omitempty"`
}

type Balance struct {
	Outstanding *float64 `json:"Outstanding,omitempty"`
	Overdue     *float64 `json:"Overdue,omitempty"`
}

type ContactBalance struct {
	AccountsReceivable *Balance `json:"AccountsReceivable,omitempty"`
	AccountsPayable    *Balance `json:"AccountsPayable,omitempty"`
}

type ContactPerson struct {
	FirstName       *string `json:"FirstName,omitempty"`
	LastName        *string `json:"LastName,omitempty"`
	EmailAddress    *string `json:"EmailAddress,omitempty"`
	IncludeInEmails *bool   `json:"IncludeInEmails,omitempty"`
}

type PaymentTerms struct {
	Bills *PaymentTerm `json:"Bills,omitempty"`
	Sales *PaymentTerm `json:"Sales,omitempty"`
}

type PaymentTerm struct {
	Day  *int   `json:"Day"`
	Type string `json:"Type"`
}

type BrandingTheme struct {
	BrandingThemeID *string `json:"BrandingThemeID,omitempty"`
	Name            *string `json:"Name,omitempty"`
}
