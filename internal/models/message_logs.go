package models

import "gorm.io/gorm"

// MessageLog is used to map your locations database table to your go code.
type MessageLog struct {
	gorm.Model
	OrganizationID uint64 `gorm:"colum:org_id" json:"org_id"`
	Recipient      string `gorm:"colum:recepient" json:"recepient"`
	Message        string `gorm:"colum:message" json:"message"`
	Status         string `gorm:"colum:status" json:"status"`
}
