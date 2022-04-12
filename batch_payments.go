package xero

type BatchPayment struct {
	Account        *Account   `json:"Account,omitempty"`
	Particulars    *string    `json:"Particulars,omitempty"`
	Code           *string    `json:"Code,omitempty"`
	Reference      *string    `json:"Reference,omitempty"`
	Narrative      *string    `json:"Narrative,omitempty"`
	BatchPaymentId *string    `json:"BatchPaymentID,omitempty"`
	DateString     *string    `json:"DateString,omitempty"`
	Date           *NetDate   `json:"Date,omitempty"`
	Payments       []*Payment `json:"Payments,omitempty"`
	Type           *string    `json:"Type,omitempty"`
	Status         *string    `json:"Status,omitempty"`
	TotalAmount    *float64   `json:"TotalAmount,omitempty"`
	UpdatedDateUTC *NetDate   `json:"UpdatedDateUTC,omitempty"`
	IsReconciled   *bool      `json:"IsReconciled,omitempty"`
}
