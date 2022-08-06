package types

import (
	"time"

	"github.com/google/uuid"
)

func (dto *CustomerDto) ToCustomer() *Customer {
	customer := Customer{
		FirstName:  dto.FirstName,
		LastName:   dto.LastName,
		Email:      dto.Email,
		BirthDate:  dto.BirthDate,
		Gender:     dto.Gender,
		Adress:     dto.Adress,
	}
	return &customer
}

func (dto *CustomerDto) ToAccountInformation(customerId uuid.UUID, passwordHash string) *AccountInformation {
	accountInformation := AccountInformation {
		CustomerId: customerId,
		PasswordHash: passwordHash,
		RegisterDate: time.Now(),
	}
	return &accountInformation
}