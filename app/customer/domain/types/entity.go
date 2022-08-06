package types

import (
	"time"

	"github.com/google/uuid"
)

type (
	Customer struct {
		Id         int            `gorm:"column:Id;primary_key"`
		CustomerId uuid.UUID      `gorm:"column:CustomerId"`
		FirstName  string         `gorm:"column:FirstName"`
		LastName   string         `gorm:"column:LastName"`
		Email      string         `gorm:"column:Email"`
		BirthDate  time.Time      `gorm:"column:BirthDate"`
		Gender     GenderTypeEnum `gorm:"column:Gender"`
		Adress     *string        `gorm:"column:Gender"`
		IsActive   bool           `gorm:"column:IsActive"`
	}

	AccountInformation struct {
		Id           int       `gorm:"column:Id;primary_key"`
		CustomerId   uuid.UUID `gorm:"column:CustomerId"`
		PasswordHash string    `gorm:"column:PasswordHash"`
		RegisterDate time.Time `gorm:"column:RegisterDate"`
	}
)