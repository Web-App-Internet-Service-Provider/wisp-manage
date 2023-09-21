package models

import (
	"time"

	"gorm.io/gorm"
)

// Invoice is used to map your invoices database table to your go code.
type Invoice struct {
	gorm.Model
	InvoiceNumber string    `gorm:"column:invoice_number" json:"invoice_number"`
	CustomerID    uint64    `gorm:"column:customer_id" json:"customer_id"`
	AmountDue     int       `gorm:"column:amount_due" json:"amount_due"`
	DueDate       time.Time `gorm:"column:due_date" json:"due_date"`
	ItemID        uint64    `gorm:"column:item_id" json:"item_id"`
	Status        string    `gorm:"column:status" json:"status"`
}
