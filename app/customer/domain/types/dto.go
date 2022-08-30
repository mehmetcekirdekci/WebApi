package types

import (
	"time"

	"github.com/google/uuid"
)

type (
	CustomerDto struct {
		CustomerId uuid.UUID
		FirstName  string
		LastName   string
		Email      string
		BirthDate  time.Time
		Gender GenderTypeEnum
		Adress *string
		Password string
		IsActive bool
	}
)