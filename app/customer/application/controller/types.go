package controller

import (
	"time"
	"net/http"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/mehmetcekirdekci/WebApi/app/customer/domain/types"
)

type (
	CustomValidator struct {
		validator *validator.Validate
	}

	BaseCustomerResponse struct {
		Success         bool   `json:"success"`
		ResponseMessage string `json:"responseMessage"`
	}

	RegisterCustomerRequest struct {
		Email     string               `json:"email" validate:"required"`
		Password  string               `json:"password" validate:"required"`
		FirstName string               `json:"firstName" validate:"required"`
		LastName  string               `json:"lastName" validate:"required"`
		BirthDate time.Time            `json:"birthDate" validate:"required"`
		Gender    types.GenderTypeEnum `json:"gender" validate:"gte=0"`
		Adress    *string              `json:"adress"`
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
	  // Optionally, you could return the error to give each route more control over the status code
	  return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
  }