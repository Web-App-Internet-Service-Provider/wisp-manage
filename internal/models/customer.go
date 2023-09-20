package models

import (
	"time"

	"gorm.io/gorm"
)

// Customer is used to map your customers database table to your go code.
type Customer struct {
	gorm.Model
	Username       string    `gorm:"colum:username" json:"username"`
	Password       string    `gorm:"colum:password" json:"password"`
	FirstName      string    `gorm:"colum:firstname" json:"firstname"`
	LastName       string    `gorm:"colum:lastname" json:"lastname"`
	PhoneNumber    string    `gorm:"colum:phone_number" json:"phone_number"`
	Email          string    `gorm:"colum:email" json:"email"`
	Gender         string    `gorm:"colum:gender" json:"gender"`
	IDNumber       string    `gorm:"colum:id_no" json:"id_no"`
	Status         string    `gorm:"colum:status" json:"status"`
	Category       string    `gorm:"colum:category" json:"category"`
	Device         string    `gorm:"colum:device" json:"device"`
	Expiry         time.Time `gorm:"colum:expiry" json:"expiry"`
	OrganizationID uint64    `gorm:"colum:organization_id" json:"organization_id"`
	PlanID         uint64    `gorm:"colum:plan_id" json:"plan_id"`
	StationID      uint64    `gorm:"colum:station_id" json:"station_id"`
}
