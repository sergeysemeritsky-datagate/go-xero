package integration

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/glebteterin/go-xero"
)

func Test_Invoices(t *testing.T) {
	ctx := context.TODO()
	f, err := os.Open("invoice.pdf")
	if err != nil {
		t.Fatalf("Error opening invoice.pdf: %v", err)
	}
	defer func() {
		_ = f.Close()
	}()

	// Create invoice
	inv := &xero.Invoice{
		Type: xero.String("ACCREC"),
		Contact: &xero.Contact{
			Name: xero.String("Test Contact"),
		},
		DateString:      xero.String("2020-08-01T00:00:00"),
		DueDateString:   xero.String("2020-08-01T00:00:00"),
		LineAmountTypes: xero.String("Exclusive"),
		LineItems: []*xero.LineItem{
			&xero.LineItem{
				Description: xero.String("Line 1"),
				Quantity:    xero.Float64(1),
				UnitAmount:  xero.Float64(116),
				AccountCode: xero.String("200"),
			},
		},
	}

	invoices, _, err := client.Invoices.CreateInvoice(ctx, inv)
	if err != nil {
		t.Fatalf("Invoices.CreateInvoice returned error: %v", err)
	}

	if len(invoices.Invoices) == 0 {
		t.Fatalf("Invoices.CreateInvoice returned no invoices")
	}

	newInvoice := invoices.Invoices[0]
	log.Println(*newInvoice.InvoiceNumber)

	// List invoices

	invoices, _, err = client.Invoices.GetInvoices(ctx, &xero.InvoiceListOptions{Status: []string{"DRAFT"}, ListOptions: xero.ListOptions{Where: fmt.Sprintf("TotalTax==%v", *newInvoice.TotalTax)}})
	if err != nil {
		t.Fatalf("Invoices.GetInvoices returned error: %v", err)
	}
	if invoices == nil {
		t.Fatalf("Invoices.GetInvoices returned nil")
	}

	if len(invoices.Invoices) == 0 {
		t.Fatalf("Invoices.GetInvoices returned no invoices")
	}

	newInvoiceFound := false
	for _, inv := range invoices.Invoices {
		if *inv.InvoiceID == *newInvoice.InvoiceID {
			newInvoiceFound = true
			break
		}
	}

	if !newInvoiceFound {
		t.Fatalf("New invoice not found")
	}

	// Update Invoice

	newUnitAmount := 120.0
	newInvoice.LineItems[0].UnitAmount = xero.Float64(newUnitAmount)
	newInvoice.LineItems[0].LineAmount = xero.Float64(newUnitAmount)
	newInvoice.LineItems[0].TaxAmount = xero.Float64(newUnitAmount * 0.15)
	invoices, _, err = client.Invoices.UpdateInvoice(ctx, newInvoice)

	if err != nil {
		t.Fatalf("Invoices.UpdateInvoice returned error: %v", err)
	}
	if invoices == nil {
		t.Fatalf("Invoices.UpdateInvoice returned nil")
	}

	if len(invoices.Invoices) == 0 {
		t.Fatalf("Invoices.UpdateInvoice returned no invoices")
	}

	// Get invoice
	invoices, _, err = client.Invoices.GetInvoices(ctx, &xero.InvoiceListOptions{Id: []string{*newInvoice.InvoiceID}})

	if err != nil {
		t.Fatalf("Invoices.GetInvoices returned error: %v", err)
	}
	if invoices == nil {
		t.Fatalf("Invoices.GetInvoices returned nil")
	}

	if len(invoices.Invoices) == 0 {
		t.Fatalf("Invoices.GetInvoices returned no invoices")
	}

	if *invoices.Invoices[0].TotalTax != (newUnitAmount * 0.15) {
		t.Fatalf("Updated invoice TotalTax is %v, want %v", *invoices.Invoices[0].TotalTax, (newUnitAmount * 0.15))
	}

	// Upload Attachment

	_, err = client.Invoices.UploadAttachment(ctx, *newInvoice.InvoiceID, "invoice.pdf", true, f)
	if err != nil {
		t.Fatalf("Invoices.UploadAttachment returned error: %v", err)
	}

	// Get Attachments

	attachments, _, err := client.Invoices.GetAttachments(ctx, *newInvoice.InvoiceID)
	if err != nil {
		t.Fatalf("Invoices.GetAttachments returned error: %v", err)
	}
	if len(attachments.Attachments) == 0 {
		t.Fatalf("Invoices.GetAttachments returned no attachments")
	}

	var invoiceAttachment *xero.Attachment
	for _, a := range attachments.Attachments {
		if *a.FileName == "invoice.pdf" {
			invoiceAttachment = a
		}
	}
	if invoiceAttachment == nil {
		t.Fatalf("Attachment invoice.pdf not found")
	}
	if invoiceAttachment.IncludeOnline == nil || *invoiceAttachment.IncludeOnline == false {
		t.Fatalf("Attachment invoice.pdf is not marked as include online")
	}

	tmpFile, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatalf("ioutil.TempFile returned error: %v", err)
	}
	defer func() {
		_ = tmpFile.Close()
	}()
	_, err = client.Invoices.GetAttachment(ctx, *newInvoice.InvoiceID, "invoice.pdf", tmpFile)
	if err != nil {
		t.Fatalf("Invoices.GetAttachment returned error: %v", err)
	}
	if err := tmpFile.Close(); err != nil {
		t.Fatalf("error closing tmp file: %v", err)
	}
	same, err := FileCmp(f.Name(), tmpFile.Name(), 0)
	if err != nil {
		t.Fatalf("FileCmp returned error: %v", err)
	}
	if !same {
		t.Fatalf("downloaded attachment is not the same as original file")
	}

	// Delete invoice

	newInvoice.Status = xero.String("DELETED")
	invoices, _, err = client.Invoices.UpdateInvoice(ctx, newInvoice)

	if err != nil {
		t.Fatalf("Invoices.UpdateInvoice returned error: %v", err)
	}
	if invoices == nil {
		t.Fatalf("Invoices.UpdateInvoice returned nil")
	}

	if len(invoices.Invoices) == 0 {
		t.Fatalf("Invoices.UpdateInvoice returned no invoices")
	}
}
