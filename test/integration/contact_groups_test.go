package integration

import (
	"context"
	"github.com/glebteterin/go-xero"
	"testing"
)

func Test_ContactGroups(t *testing.T) {
	ctx := context.TODO()

	// Create group

	group := &xero.ContactGroup{
		Name: xero.String("Test Group"),
	}

	res, _, err := client.ContactGroups.CreateContactGroup(ctx, group)
	if err != nil {
		t.Fatalf("ContactGroups.CreateContactGroup returned error: %v", err)
	}

	if len(res.ContactGroups) == 0 {
		t.Fatalf("ContactGroups.CreateContactGroup returned no groups")
	}

	newGroup := res.ContactGroups[0]

	if *newGroup.Name != *group.Name {
		t.Fatalf("New group is %v, want %v", *newGroup.Name, *group.Name)
	}

	// Add contact to the group

	contacts, _, err := client.Contacts.GetContacts(ctx, &xero.ContactListOptions{})
	if err != nil {
		t.Fatalf("Contacts.GetContacts returned error: %v", err)
	}

	groupContacts, _, err := client.ContactGroups.AddContacts(ctx, *newGroup.ContactGroupId, []*xero.Contact{{ContactId: contacts.Contacts[0].ContactId}})
	if err != nil {
		t.Fatalf("ContactGroups.AddContacts returned error: %v", err)
	}

	if len(groupContacts.Contacts) == 0 {
		t.Fatalf("ContactGroups.AddContacts returned no contacts")
	}

	// Remove all contacts from the group

	_, err = client.ContactGroups.RemoveContacts(ctx, *newGroup.ContactGroupId)
	if err != nil {
		t.Fatalf("ContactGroups.RemoveContacts returned error: %v", err)
	}

	// Verify deletion

	groups, _, err := client.ContactGroups.GetContactGroup(ctx, *newGroup.ContactGroupId)
	if err != nil {
		t.Fatalf("Contacts.GetContactGroup returned error: %v", err)
	}

	if len(groups.ContactGroups) == 0 {
		t.Fatalf("Contacts.GetContactGroup returned no groups")
	}

	if len(groups.ContactGroups[0].Contacts) > 0 {
		t.Fatalf("Group still has contacts")
	}

	newGroup.Status = xero.String("DELETED")
	_, _, err = client.ContactGroups.UpdateContactGroup(ctx, newGroup)
	if err != nil {
		t.Fatalf("ContactGroups.UpdateContactGroup returned error: %v", err)
	}
}
