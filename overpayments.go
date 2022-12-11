package xero

type Overpayment struct {
	OverpaymentId   *string       `json:"OverpaymentID,omitempty"`
	Type            *string       `json:"Type,omitempty"`
	Contact         *Contact      `json:"Contact,omitempty"`
	Date            *NetDate      `json:"Date,omitempty"`
	Status          *string       `json:"Status,omitempty"`
	LineAmountTypes *string       `json:"LineAmountTypes,omitempty"`
	LineItems       []*LineItem   `json:"LineItems,omitempty"`
	SubTotal        *float64      `json:"SubTotal,omitempty"`
	TotalTax        *float64      `json:"TotalTax,omitempty"`
	Total           *float64      `json:"Total,omitempty"`
	UpdatedDateUTC  *NetDate      `json:"UpdatedDateUTC,omitempty"`
	CurrencyCode    *string       `json:"CurrencyCode,omitempty"`
	CurrencyRate    *float64      `json:"CurrencyRate,omitempty"`
	RemainingCredit *float64      `json:"RemainingCredit,omitempty"`
	Allocations     []*Allocation `json:"Allocations,omitempty"`
	Payments        []*Payment    `json:"Payments,omitempty"`
	HasAttachments  *bool         `json:"HasAttachments,omitempty"`
}

type Allocation struct {
	Amount     *string  `json:"Amount,omitempty"`
	DateString *string  `json:"DateString,omitempty"`
	Date       *NetDate `json:"Date,omitempty"`
	Invoice    *Invoice `json:"Invoice,omitempty"`
}
