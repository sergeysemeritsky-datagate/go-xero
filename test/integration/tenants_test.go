package integration

import (
	"context"
	"testing"
)

func Test_Tenants(t *testing.T) {
	tenants, _, err := client.Tenants.GetTenants(context.TODO())
	if err != nil {
		t.Fatalf("Tenants.GetTenants returned error: %v", err)
	}
	if tenants == nil {
		t.Fatalf("Tenants.GetTenants returned nil")
	}

	if len(tenants) == 0 {
		t.Fatalf("Tenants.GetTenants returned no accounts")
	}
}