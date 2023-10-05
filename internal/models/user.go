package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

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

// func HashPassword-> GenerateFromPassword returns the bcrypt hash of the password at the given cost
func HashPassword(pass *string) {
	bytePass := []byte(*pass)
	hPass, _ := bcrypt.GenerateFromPassword(bytePass, bcrypt.DefaultCost)
	*pass = string(hPass)
}

// func ComparePassword -> CompareHashAndPassword compares a bcrypt hashed password with its possible plaintext equivalent. Returns nil on success, or an error on failure.
func (user *User) ComparePassword(databasePassword string, providedPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(databasePassword), []byte(providedPassword)) == nil
}
