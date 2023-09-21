package models

import "gorm.io/gorm"

// Organization is used to map your locations database table to your go code.
type Organization struct {
	gorm.Model
	Name            string `gorm:"column:organization_name" json:"organization_name"`
	ShortName       string `gorm:"column:organization_short_name" json:"organization_short_name"`
	Website         string `gorm:"column:organization_website" json:"organization_website"`
	Email           string `gorm:"column:organization_email" json:"organization_email"`
	SecondEmail     string `gorm:"column:organization_secondemail" json:"organization_secondemail"`
	PhoneNumber     string `gorm:"column:organization_phone_no" json:"organization_phone_no"`
	MobileNumber    string `gorm:"column:organization_mobile_no" json:"organization_mobile_no"`
	Address         string `gorm:"column:organization_address" json:"organization_address"`
	Pin             string `gorm:"column:organization_pin" json:"organization_pin"`
	BankAccount     string `gorm:"column:organization_bank_acc" json:"organization_bank_acc"`
	Vat             string `gorm:"column:organization_vat" json:"organization_vat"`
	MessageProvider string `gorm:"column:organization_message_provider" json:"organization_message_provider"`
}
