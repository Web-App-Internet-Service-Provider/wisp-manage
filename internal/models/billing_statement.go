package models

import (
	"gorm.io/gorm"
)

// BillingStatement is used to map your billing_statements database table to your go code.
type BillingStatement struct {
	gorm.Model
	InvoiceNumber string `gorm:"column:invoice_number" json:"invoice_number"`
	DebitAmount   int    `gorm:"column:debit_amount" json:"debit_amount"`
	CreditAmount  int    `gorm:"column:credit_amount" json:"credit_amount"`
	CustomerID    uint64 `gorm:"column:customer_id" json:"customer_id"`
}
