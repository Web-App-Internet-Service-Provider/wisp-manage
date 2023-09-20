package models

import "gorm.io/gorm"

// Paymentmethod is used to map your paymentmethods database table to your go code.
type PaymentMethod struct {
	gorm.Model
	Name            string `gorm:"column:payment_method_name" json:"payment_method_name"`
	PaymentDetailID uint64 `gorm:"column:paymentdetail" json:"paymentdetail"`
}
