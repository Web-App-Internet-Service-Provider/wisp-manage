package api

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

// Model definition
type Model struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type User struct {
	Model
	Email        string `gorm:"email" json:"email"`
	Name         string `gorm:"name" json:"name"`
	PhoneNumber  string `gorm:"phonenumber" json:"phonenumber"`
	PasswordHash string `gorm:"password_hash" json:"password_hash"`

	Role          *Role         `belongs_to:"role"`
	Organization  *Organization `belongs_to:"organization"`
	Extensionlogs ExtensionLogs `gorm:"extensionlog,omitempty" has_many:"extensionlog" `

	Password             string `gorm:"-" json:"-"`
	PasswordConfirmation string `gorm:"-" json:"-"`
}

type Station struct {
	Model
	Name           string        `gorm:"name" json:"name"`
	IP_Address     string        `gorm:"ipaddress" json:"ipaddress"`
	NAS_Password   string        `gorm:"naspassword" json:"naspassword"`
	NasApiID       int           `gorm:"nas_api_id" json:"nas_api_id"`
	LocationID     uuid.UUID     `gorm:"location_id" json:"location_id"`
	OrganizationID uuid.UUID     `gorm:"organization_id" json:"organization_id"`
	Plans          Plans         `gorm:"plans" many_to_many:"planstations" json:"-"`
	Customers      []Customer    `gorm:"customer,omitempty" has_many:"customer" `
	Location       *Location     `belongs_to:"location" `
	Organization   *Organization `belongs_to:"organization"`
}

type Service struct {
	Model
	CustomerID    uuid.UUID     `gorm:"customer_id" json:"customer_id"`
	PlanID        uuid.UUID     `gorm:"plan_id" json:"plan_id"`
	Expiry        time.Time     `gorm:"expirytime" json:"expirytime"`
	Status        string        `gorm:"status" json:"status"`
	Type          string        `gorm:"type" json:"type"`
	Customer      Customer      `belongs_to:"customer"`
	Plan          Plan          `belongs_to:"Plan"`
	Extensionlogs ExtensionLogs `gorm:"extensionlog,omitempty" has_many:"extensionlog" `
}

type Role struct {
	Model
	Name         string      `gorm:"name" json:"name"`
	PermissionID uuid.UUID   `gorm:"permission_id" json:"permission_id"`
	Users        []User      `gorm:"user,omitempty" has_many:"user" order_by:"id desc"`
	Permission   *Permission `belongs_to:"permission"`
}

type Pool struct {
	Model
	Name           string    `gorm:"name" json:"name"`
	OrganizationID uuid.UUID `gorm:"organization_id" json:"organization_id"`
	Plan           *Plan     `has_one:"plan"`
}

type Plantype struct {
	Model
	Name  string `gorm:"name" json:"name"`
	Plans []Plan `gorm:"plan,omitempty" has_many:"plan" order_by:"id desc"`
}

type Planstation struct {
	Model
	StationID uuid.UUID `gorm:"station_id" json:"station_id"`
	PlanID    uuid.UUID `gorm:"plan_id" json:"plan_id"`

	Plans    Plan    `belongs_to:"plan"`
	Stations Station `belongs_to:"station"`
}

type Plan struct {
	Model
	Name           string       `gorm:"name" json:"name"`
	PlantypeID     uuid.UUID    `gorm:"plantype_id" json:"plantype_id"`
	Price          int          `gorm:"price" json:"price"`
	ExpirationType string       `gorm:"expirationtype" json:"expirationtype"`
	Speed          string       `gorm:"speed" json:"speed"`
	PoolID         uuid.UUID    `gorm:"pool_id" json:"pool_id"`
	Device         string       `gorm:"device" json:"device"`
	OrganizationID uuid.UUID    `gorm:"organization_id" json:"organization_id"`
	Stations       Stations     `gorm:"stations" many_to_many:"planstations" json:"-"`
	Customers      []Customer   `gorm:"customer,omitempty" has_many:"customers" order_by:"id desc"`
	Plantype       Plantype     `belongs_to:"plantype"`
	Pool           Pool         `belongs_to:"pool"`
	Service        Services     `gorm:"service,omitempty" has_many:"services" order_by:"id desc"`
	Organization   Organization `belongs_to:"organization"`
}

type Permission struct {
	Model
	Name string `gorm:"name" json:"name"`

	Roles []Role `gorm:"role,omitempty" has_many:"role" order_by:"id desc"`
}

type Paymentmethod struct {
	Model
	Name string `gorm:"name" json:"name"`

	Paymentdetail Paymentdetails `gorm:"paymentdetail,omitempty" has_many:"paymentdetail"`
}

type Paymentdetail struct {
	Model
	Amount      int       `gorm:"amount" json:"amount"`
	PaymentDate time.Time `gorm:"Payment_date" json:"payment_date"`

	Paymentmethod *Paymentmethod `belongs_to:"paymentmethod"`
	Customer      *Customer      `belongs_to:"customer"`
}

type Organization struct {
	Model
	Name            string     `gorm:"name" json:"name"`
	ShortName       string     `gorm:"short_name" json:"short_name"`
	Website         string     `gorm:"website" json:"website"`
	Email           string     `gorm:"email" json:"email"`
	SecondEmail     string     `gorm:"secondemail" json:"secondemail"`
	PhoneNumber     string     `gorm:"phone_no" json:"phone_no"`
	MobileNumber    string     `gorm:"mobile_no" json:"mobile_no"`
	Address         string     `gorm:"address" json:"address"`
	Pin             string     `gorm:"pin" json:"pin"`
	BankAccount     string     `gorm:"bank_acc" json:"bank_acc"`
	Vat             string     `gorm:"vat" json:"vat"`
	Locations       []Location `gorm:"locations,omitempty" has_many:"locations"`
	Plans           []Plan     `gorm:"plans,omitempty" has_many:"plans"`
	Stations        []Station  `gorm:"stations,omitempty" has_many:"stations"`
	Users           []User     `gorm:"users,omitempty" has_many:"users"`
	MessageProvider string     `gorm:"message_provider" json:"message_provider"`
}

type OrgSetting struct {
	Model
	Key   string `gorm:"key" json:"key"`
	Value string `gorm:"value" json:"value"`
}

type Messagetemplate struct {
	ID      uuid.UUID `gorm:"id" json:"id"`
	Type    string    `gorm:"type" json:"type"`
	Message string    `gorm:"message" json:"message"`
}

type Messagelog struct {
	Model
	Recipient string `gorm:"recepient" json:"recepient"`
	Message   string `gorm:"message" json:"message"`
	Status    string `gorm:"status" json:"status"`
}

type Location struct {
	Model
	Name string `gorm:"name" json:"name"`

	Organization Organization `belongs_to:"organization"`
	Stations     []Station    `gorm:"stations,omitempty" has_many:"station"`
}

type Invoice struct {
	Model
	InvoiceNumber string    `gorm:"invoice_number" json:"invoice_number"`
	AmountDue     int       `gorm:"amount_due" json:"amount_due"`
	DueDate       time.Time `gorm:"due_date" json:"due_date"`
	ItemID        uuid.UUID `gorm:"item_id" json:"item_id"`
	Status        string    `gorm:"status" json:"status"`

	Customer Customer `belongs_to:"customer"`
}

type ExtensionLog struct {
	Model
	Duration int `gorm:"duration" json:"duration"`

	User    User    `belongs_to:"user"`
	Service Service `belongs_to:"service"`
}

type Customer struct {
	Model
	Username       string         `gorm:"username" json:"username"`
	Password       string         `gorm:"password" json:"password"`
	FirstName      string         `gorm:"firstname" json:"firstname"`
	LastName       string         `gorm:"lastname" json:"lastname"`
	PhoneNumber    string         `gorm:"phonenumber" json:"phonenumber"`
	Email          string         `gorm:"email" json:"email"`
	Gender         string         `gorm:"gender" json:"gender"`
	IDNumber       string         `gorm:"id_no" json:"id_no"`
	Status         string         `gorm:"status" json:"status"`
	Category       string         `gorm:"category" json:"category"`
	Device         string         `gorm:"device" json:"device"`
	Expiry         time.Time      `gorm:"expiry" json:"expiry"`
	Organization   Organization   `belongs_to:"organization"`
	Plan           Plan           `belongs_to:"plan"`
	Station        Station        `belongs_to:"station"`
	Paymentdetails Paymentdetails `gorm:"paymentdetail,omitempty" has_many:"paymentdetail"`
	Services       Services       `gorm:"has_many:services;"`
	Invoices       Invoices       `gorm:"has_many:invoices;"`
}

type BillingStatement struct {
	Model
	InvoiceNumber string `gorm:"invoice_number" gorm:"invoice_number" json:"invoice_number"`
	DebitAmount   int    `gorm:"debit_amount" gorm:"debit_amount" json:"debit_amount"`
	CreditAmount  int    `gorm:"credit_amount" gorm:"credit_amount" json:"credit_amount"`

	Customer Customer `belongs_to:"customer"`
}

// `gorm:"references:CompanyID"`
