package xero

import (
	"context"
	"net/http"
)

type BrandingThemesService service

func (s *BrandingThemesService) GetBrandingThemes(ctx context.Context) (*BrandingThemesResponse, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "api.xro/2.0/BrandingThemes", nil)
	if err != nil {
		return nil, nil, err
	}

	var c *BrandingThemesResponse
	resp, err := s.client.Do(ctx, req, &c)
	if err != nil {
		return nil, nil, err
	}

	return c, resp, nil
}

type BrandingThemesResponse struct {
	BrandingThemes []*BrandingTheme `json:"BrandingThemes"`
}
