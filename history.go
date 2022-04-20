package xero

import (
	"context"
	"fmt"
	"net/http"
)

type HistoryRecords struct {
	Records []*HistoryRecord `json:"HistoryRecords"`
}

type HistoryRecord struct {
	Changes       *string  `json:"Changes"`
	DateUTCString *string  `json:"DateUTCString"`
	DateUTC       *NetDate `json:"DateUTC,omitempty"`
	User          *string  `json:"User,omitempty"`
	Details       *string  `json:"Details,omitempty"`
}

func getHistory(ctx context.Context, client *Client, entity string, id string) (*HistoryRecords, *http.Response, error) {
	u := fmt.Sprintf("api.xro/2.0/%s/%s/history", entity, id)

	req, err := client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var c *HistoryRecords
	resp, err := client.Do(ctx, req, &c)
	if err != nil {
		return nil, nil, err
	}

	return c, resp, nil
}

func addHistory(ctx context.Context, client *Client, entity, id string, records *HistoryRecords) (*http.Response, error) {
	u := fmt.Sprintf("api.xro/2.0/%s/%s/history", entity, id)

	req, err := client.NewRequest("PUT", u, records)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(ctx, req, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
