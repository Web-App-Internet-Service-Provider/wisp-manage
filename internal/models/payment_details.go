package models

import (
	"time"

	"gorm.io/gorm"
)

// Paymentdetail is used by pop to map your paymentdetails database table to your go code.
type PaymentDetail struct {
	gorm.Model
	CustomerID      uint64    `gorm:"colum:customer_id" json:"customer_id"`
	PaymentmethodID uint64    `gorm:"colum:payment_mode_id" json:"payment_mode_id"`
	Amount          int       `gorm:"colum:amount" json:"amount"`
	PaymentDate     time.Time `gorm:"colum:payment_date" json:"payment_date"`
}
