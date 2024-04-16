package integration

import (
	"context"
	"testing"
)

func Test_Branding_Themes(t *testing.T) {
	brandingThemes, _, err := client.BrandingThemes.GetBrandingThemes(context.TODO())
	if err != nil {
		t.Fatalf("BrandingThemeService.GetBrandingThemes returned error: %v", err)
	}
	if brandingThemes == nil {
		t.Fatalf("BrandingThemeService.GetBrandingThemes returned nil")
	}
}
