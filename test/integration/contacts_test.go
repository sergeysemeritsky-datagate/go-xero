package integration

import (
	"context"
	"github.com/glebteterin/go-xero"
	"testing"
)

func Test_GetContacts(t *testing.T) {
	contacts, _, err := client.Contacts.GetContacts(context.TODO(),
		&xero.ContactListOptions{ListOptions: xero.ListOptions{
			Where: "ContactStatus==\"ARCHIVED\" && (IsSupplier==false && IsCustomer == false)"},
			IncludeArchived: true,
			Page:            1,
		})

	if err != nil {
		t.Fatalf("Contacts.GetContacts returned error: %v", err)
	}
	if contacts == nil {
		t.Fatalf("Contacts.GetContacts returned nil")
	}

	if len(contacts.Contacts) == 0 {
		t.Fatalf("Accounts.GetAccounts returned no contacts")
	}

	for _, c := range contacts.Contacts {
		if *c.ContactStatus != "ARCHIVED" {
			t.Errorf("Accounts.GetAccounts returned not archived contact")
			break
		}
		if *c.IsSupplier != false {
			t.Errorf("Accounts.GetAccounts returned supplied")
			break
		}
		if *c.IsCustomer != false {
			t.Errorf("Accounts.GetAccounts returned customer")
			break
		}
	}
}
