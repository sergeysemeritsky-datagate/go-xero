package xero

import (
	"context"
	"net/http"
)

type AccountsService service

func (s *AccountsService) GetAccounts(ctx context.Context) (*AccountsResponse, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "api.xro/2.0/accounts", nil)
	if err != nil {
		return nil, nil, err
	}

	var c *AccountsResponse
	resp, err := s.client.Do(ctx, req, &c)
	if err != nil {
		return nil, nil, err
	}

	return c, resp, nil
}

type AccountsResponse struct {
	Accounts []*Account `json:"Accounts"`
	Status   string     `json:"Status"`
}

type Account struct {
	AccountID               *string  `json:"AccountID,omitempty"`
	Code                    *string  `json:"Code,omitempty"`
	Name                    *string  `json:"Name,omitempty"`
	Status                  *string  `json:"Status,omitempty"`
	Type                    *string  `json:"Type,omitempty"`
	TaxType                 *string  `json:"TaxType,omitempty"`
	Description             *string  `json:"Description,omitempty"`
	Class                   *string  `json:"Class,omitempty"`
	SystemAccount           *string  `json:"SystemAccount,omitempty"`
	EnablePaymentsToAccount *bool    `json:"EnablePaymentsToAccount,omitempty"`
	ShowInExpenseClaims     *bool    `json:"ShowInExpenseClaims,omitempty"`
	BankAccountType         *string  `json:"BankAccountType,omitempty"`
	ReportingCode           *string  `json:"ReportingCode,omitempty"`
	ReportingCodeName       *string  `json:"ReportingCodeName,omitempty"`
	HasAttachments          *bool    `json:"HasAttachments,omitempty"`
	UpdatedDateUTC          *NetDate `json:"UpdatedDateUTC,omitempty"`
	AddToWatchlist          *bool    `json:"AddToWatchlist,omitempty"`
}

