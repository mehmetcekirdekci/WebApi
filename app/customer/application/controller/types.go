package controller

import (
	"errors"
	"time"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/mehmetcekirdekci/WebApi/app/customer/domain/types"
)

type (
	CustomValidator struct {
		validator *validator.Validate
		translator *ut.Translator
	}

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

func (cv *CustomValidator) Validate(i interface{}) error {
	err := cv.validator.Struct(i)
	if err != nil {
		var translatedErrors string
		validationErrors := err.(validator.ValidationErrors).Translate((*cv.translator))
		for _, val := range validationErrors {
			translatedErrors += val + "."
		}
		return errors.New(translatedErrors)
	}
	return nil
}