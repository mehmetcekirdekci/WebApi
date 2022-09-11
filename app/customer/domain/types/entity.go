package types

import (
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		Id           primitive.ObjectID     `json:"id bson:"_id,omitempty"`
		CustomerId   uuid.UUID 				`bson:"customerId"`
		PasswordHash string 				`bson:"passwordHash"`   
		RegisterDate time.Time 				`bson:"registerDate"`
		Jwt 		 string 				`bson:"jwt"`
	}
)