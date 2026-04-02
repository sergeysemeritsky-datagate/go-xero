package integration

import (
	"context"
	"testing"
)

func TestItems(t *testing.T) {
	items, _, err := client.Items.GetItems(context.TODO())
	if err != nil {
		t.Fatalf("Items.GetItems returned error: %v", err)
	}
	if items == nil {
		t.Fatalf("Items.GetItems returned nil")
	}

	if len(items.Items) == 0 {
		t.Fatalf("Items.GetItems returned no accounts")
	}

	if items.Items[0].ItemID == nil {
		t.Errorf("Items.GetItems returned account with no ID")
	}
}
