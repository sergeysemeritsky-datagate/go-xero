package xero

import (
	"context"
	"net/http"
)

type TrackingCategoriesService service

func (s *TrackingCategoriesService) GetTrackingCategories(ctx context.Context) (*TrackingCategoriesResponse, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "api.xro/2.0/TrackingCategories", nil)
	if err != nil {
		return nil, nil, err
	}

	var c *TrackingCategoriesResponse
	resp, err := s.client.Do(ctx, req, &c)
	if err != nil {
		return nil, nil, err
	}

	return c, resp, nil
}

type TrackingCategoriesResponse struct {
	TrackingCategories []*TrackingCategory `json:"TrackingCategories"`
}

type Options struct {
	TrackingOptionID string `json:"TrackingOptionID"`
	Name             string `json:"Name"`
	Status           string `json:"Status"`
}

type TrackingCategory struct {
	Name               string     `json:"Name"`
	Status             string     `json:"Status"`
	TrackingCategoryID string     `json:"TrackingCategoryID"`
	Options            []*Options `json:"Options"`
}
