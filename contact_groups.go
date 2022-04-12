package xero

import (
	"context"
	"fmt"
	"net/http"
)

type ContactGroupsService service

func (s *ContactGroupsService) GetContactGroup(ctx context.Context, id string) (*ContactGroupsResponse, *http.Response, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("api.xro/2.0/contactgroups/%s", id), nil)
	if err != nil {
		return nil, nil, err
	}

	var c *ContactGroupsResponse
	resp, err := s.client.Do(ctx, req, &c)
	if err != nil {
		return nil, nil, err
	}

	return c, resp, nil
}

func (s *ContactGroupsService) GetContactGroups(ctx context.Context, opts *ListOptions) (*ContactGroupsResponse, *http.Response, error) {
	u := "api.xro/2.0/contactgroups"
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var c *ContactGroupsResponse
	resp, err := s.client.Do(ctx, req, &c)
	if err != nil {
		return nil, nil, err
	}

	return c, resp, nil
}

func (s *ContactGroupsService) UpdateContactGroup(ctx context.Context, group *ContactGroup) (*ContactGroup, *http.Response, error) {
	req, err := s.client.NewRequest("POST", "api.xro/2.0/contactgroups/"+*group.ContactGroupId, group)
	if err != nil {
		return nil, nil, err
	}

	var c *ContactGroup
	resp, err := s.client.Do(ctx, req, &c)
	if err != nil {
		return nil, nil, err
	}

	return c, resp, nil
}

func (s *ContactGroupsService) CreateContactGroup(ctx context.Context, group *ContactGroup) (*ContactGroupsResponse, *http.Response, error) {
	req, err := s.client.NewRequest("PUT", "api.xro/2.0/contactgroups", group)
	if err != nil {
		return nil, nil, err
	}

	var c *ContactGroupsResponse
	resp, err := s.client.Do(ctx, req, &c)
	if err != nil {
		return nil, nil, err
	}

	return c, resp, nil
}

func (s *ContactGroupsService) AddContacts(ctx context.Context, groupId string, contacts []*Contact) (*ContactsResponse, *http.Response, error) {
	type Contacts struct {
		Contacts []*Contact `json:"Contacts"`
	}

	payload := &Contacts{
		Contacts: contacts,
	}

	req, err := s.client.NewRequest("PUT", fmt.Sprintf("api.xro/2.0/contactgroups/%s/contacts", groupId), payload)
	if err != nil {
		return nil, nil, err
	}

	var c *ContactsResponse
	resp, err := s.client.Do(ctx, req, &c)
	if err != nil {
		return nil, nil, err
	}

	return c, resp, nil
}

func (s *ContactGroupsService) RemoveContact(ctx context.Context, groupId string, contactId string) (*ContactsResponse, *http.Response, error) {
	req, err := s.client.NewRequest("DELETE",
		fmt.Sprintf("api.xro/2.0/contactgroups/%s/contacts/%s", groupId, contactId), nil)
	if err != nil {
		return nil, nil, err
	}

	var c *ContactsResponse
	resp, err := s.client.Do(ctx, req, &c)
	if err != nil {
		return nil, nil, err
	}

	return c, resp, nil
}

func (s *ContactGroupsService) RemoveContacts(ctx context.Context, groupId string) (*http.Response, error) {
	req, err := s.client.NewRequest("DELETE",
		fmt.Sprintf("api.xro/2.0/contactgroups/%s/contacts", groupId), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type ContactGroupsResponse struct {
	ContactGroups []*ContactGroup `json:"ContactGroups"`
	Status        string          `json:"Status"`
	ErrorNumber   *int            `json:"ErrorNumber,omitempty"`
}

type ContactGroup struct {
	ContactGroupId      *string    `json:"ContactGroupID,omitempty"`
	Name                *string    `json:"Name,omitempty"`
	Status              *string    `json:"Status,omitempty"`
	Contacts            []*Contact `json:"Contacts,omitempty"`
	HasValidationErrors *bool      `json:"HasValidationErrors,omitempty"`
}
