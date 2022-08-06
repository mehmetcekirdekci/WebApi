package controller

import (
	"time"

	"github.com/mehmetcekirdekci/WebApi/app/customer/domain/types"
)

type (
	BaseCustomerResponse struct {
		Success         bool   `json:"success"`
		ResponseMessage string `json:"responseMessage"`
	}

	RegisterCustomerRequest struct {
		Email     string               `json:"email"`
		Password  string               `json:"password"`
		FirstName string               `json:"firstName"`
		LastName  string               `json:"lastName"`
		BirthDate time.Time            `json:"birthDate"`
		Gender    types.GenderTypeEnum `json:"gender"`
		Adress    *string              `json:"adress"`
	}
)