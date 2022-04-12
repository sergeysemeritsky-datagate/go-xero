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
	AccountID               string  `json:"AccountID"`
	Code                    string  `json:"Code"`
	Name                    string  `json:"Name"`
	Status                  string  `json:"Status"`
	Type                    string  `json:"Type"`
	TaxType                 string  `json:"TaxType"`
	Description             string  `json:"Description"`
	Class                   string  `json:"Class"`
	SystemAccount           string  `json:"SystemAccount"`
	EnablePaymentsToAccount bool    `json:"EnablePaymentsToAccount"`
	ShowInExpenseClaims     bool    `json:"ShowInExpenseClaims"`
	BankAccountType         string  `json:"BankAccountType"`
	ReportingCode           string  `json:"ReportingCode"`
	ReportingCodeName       string  `json:"ReportingCodeName"`
	HasAttachments          bool    `json:"HasAttachments"`
	UpdatedDateUTC          NetDate `json:"UpdatedDateUTC"`
	AddToWatchlist          bool    `json:"AddToWatchlist"`
}

