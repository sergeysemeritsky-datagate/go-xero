package integration

import (
	"context"
	"testing"
)

func TestAccounts(t *testing.T) {
	accounts, _, err := client.Accounts.GetAccounts(context.TODO())
	if err != nil {
		t.Fatalf("Accounts.GetAccounts returned error: %v", err)
	}
	if accounts == nil {
		t.Fatalf("Accounts.GetAccounts returned nil")
	}

	if len(accounts.Accounts) == 0 {
		t.Fatalf("Accounts.GetAccounts returned no accounts")
	}

	if accounts.Accounts[0].AccountID == nil {
		t.Errorf("Accounts.GetAccounts returned account with no ID")
	}
}
