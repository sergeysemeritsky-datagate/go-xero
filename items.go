package xero

import (
	"context"
	"net/http"
)

type ItemsService service

func (s *ItemsService) GetItems(ctx context.Context) (*ItemsResponse, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "api.xro/2.0/items", nil)
	if err != nil {
		return nil, nil, err
	}

	var c *ItemsResponse
	resp, err := s.client.Do(ctx, req, &c)
	if err != nil {
		return nil, nil, err
	}

	return c, resp, nil
}

type ItemsResponse struct {
	Items  []*Item `json:"Items"`
	Status string  `json:"Status"`
}

type Item struct {
	ItemID                    *string         `json:"ItemID,omitempty"`
	Code                      *string         `json:"Code,omitempty"`
	Name                      *string         `json:"Name,omitempty"`
	IsSold                    *bool           `json:"IsSold,omitempty"`
	IsPurchased               *bool           `json:"IsPurchased,omitempty"`
	Description               *string         `json:"Description,omitempty"`
	PurchaseDescription       *string         `json:"PurchaseDescription,omitempty"`
	PurchaseDetails           *PurchaseDetail `json:"PurchaseDetails,omitempty"`
	SalesDetails              *SaleDetail     `json:"SalesDetails,omitempty"`
	IsTrackedAsInventory      *bool           `json:"IsTrackedAsInventory,omitempty"`
	InventoryAssetAccountCode *string         `json:"InventoryAssetAccountCode,omitempty"`
	TotalCostPool             *float64        `json:"TotalCostPool,omitempty"`
	QuantityOnHand            *float64        `json:"QuantityOnHand,omitempty"`
	UpdatedDateUTC            *NetDate        `json:"UpdatedDateUTC,omitempty"`
}

type PurchaseDetail struct {
	UnitPrice       *float64 `json:"UnitPrice,omitempty"`
	AccountCode     *string  `json:"AccountCode,omitempty"`
	COGSAccountCode *string  `json:"COGSAccountCode,omitempty"`
	TaxType         *string  `json:"TaxType,omitempty"`
}

type SaleDetail struct {
	UnitPrice   *float64 `json:"UnitPrice,omitempty"`
	AccountCode *string  `json:"AccountCode,omitempty"`
	TaxType     *string  `json:"TaxType,omitempty"`
}
