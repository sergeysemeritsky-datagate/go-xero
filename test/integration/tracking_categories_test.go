package integration

import (
	"context"
	"testing"
)

func Test_Tracking_Categories(t *testing.T) {
	trackingCategories, _, err := client.TrackingCategories.GetTrackingCategories(context.TODO())
	if err != nil {
		t.Fatalf("TrackingCategoryService.GetTrackingCategories returned error: %v", err)
	}
	if trackingCategories == nil {
		t.Fatalf("TrackingCategoryService.GetTrackingCategories returned nil")
	}
}
