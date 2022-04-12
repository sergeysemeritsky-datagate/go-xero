package integration

import (
	"context"
	"fmt"
	"github.com/glebteterin/go-xero"
	"testing"
)

func Test_Invoices(t *testing.T) {
	ctx := context.TODO()

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
