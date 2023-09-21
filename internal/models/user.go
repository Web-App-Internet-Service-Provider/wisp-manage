package models

import "gorm.io/gorm"

// User is used to map your locations database table to your go code.
type User struct {
	gorm.Model
	Email          string `gorm:"column:user_email" json:"email"`
	Name           string `gorm:"column:name" json:"name"`
	PhoneNumber    string `gorm:"column:phone_number" json:"phone_number"`
	OrganizationID uint64 `gorm:"column:organization_id" json:"organization_id"`
	PasswordHash   string `gorm:"column:password_hash" json:"password_hash"`
	Password       string `gorm:"column:password" json:"password"`
}
