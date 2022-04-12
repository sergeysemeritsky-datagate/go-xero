package xero

import (
	"context"
	"net/http"
)

type ConnectionsService service

func (s *ConnectionsService) GetConnections(ctx context.Context) ([]*Connection, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "connections", nil)
	if err != nil {
		return nil, nil, err
	}

	var c []*Connection
	resp, err := s.client.Do(ctx, req, &c)
	if err != nil {
		return nil, nil, err
	}

	return c, resp, nil
}

type Connection struct {
	Id             string `json:"id"`
	AuthEventId    string `json:"authEventId"`
	TenantId       string `json:"tenantId"`
	TenantType     string `json:"tenantType"`
	TenantName     string `json:"tenantName"`
	CreatedDateUtc string `json:"createdDateUtc"`
	UpdatedDateUtc string `json:"updatedDateUtc"`
}
