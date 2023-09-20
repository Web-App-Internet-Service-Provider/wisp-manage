package models

import "gorm.io/gorm"

// MessageTemplate is used to map your locations database table to your go code.
type MessageTemplate struct {
	gorm.Model
	Type    string `gorm:"column:type" json:"type"`
	Message string `gorm:"column:message" json:"message"`
}
