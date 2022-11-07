package xero

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

type AttachmentsResponse struct {
	Attachments []*Attachment `json:"Attachments"`
}

type Attachment struct {
	AttachmentID  *string  `json:"AttachmentID,omitempty"`
	FileName      *string  `json:"FileName,omitempty"`
	Url           *string  `json:"Url,omitempty"`
	MimeType      *string  `json:"MimeType,omitempty"`
	ContentLength *float64 `json:"ContentLength,omitempty"`
	IncludeOnline *bool    `json:"IncludeOnline,omitempty"`
}

func getAttachments(ctx context.Context, client *Client, entity, id string) (*AttachmentsResponse, *http.Response, error) {
	u := fmt.Sprintf("api.xro/2.0/%s/%s/attachments", entity, id)

	req, err := client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var c *AttachmentsResponse
	resp, err := client.Do(ctx, req, &c)
	if err != nil {
		return nil, nil, err
	}

	return c, resp, nil
}

func getAttachment(ctx context.Context, client *Client, entity, id, filename string, w io.Writer) (*http.Response, error) {
	u := fmt.Sprintf("api.xro/2.0/%s/%s/attachments/%s", entity, id, filename)

	req, err := client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(ctx, req, w)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func uploadAttachment(ctx context.Context, client *Client, entity, id, filename string, includeOnline bool, r io.Reader) (*http.Response, error) {
	u := fmt.Sprintf("api.xro/2.0/%s/%s/attachments/%s", entity, id, filename)
	if includeOnline {
		u += "?IncludeOnline=true"
	}

	req, err := client.NewRequest("POST", u, r)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(ctx, req, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
