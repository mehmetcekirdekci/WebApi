package controller

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/mehmetcekirdekci/WebApi/app/customer/application"
)

type resource struct {
	service application.Service
}

func NewController(service application.Service) *resource {
	return &resource{
		service: service,
	}
}

// registerCustomer godoc
// @Summary Register the customer
// @Tags customer
// @Param request body RegisterCustomerRequest true "RegisterCustomerRequest"
// @Success 201 {object} BaseCustomerResponse
// @Failure 400 {object} BaseCustomerResponse
// @Router /api/v1/customer	[post]
func (receiver *resource) registerCustomer(c echo.Context) error  {
	c.Echo().Validator = &CustomValidator{validator: validator.New()}
	request := new(RegisterCustomerRequest)
	result := BaseCustomerResponse {
		Success: false,
	}
	err := c.Bind(request)
	if err != nil {
		result.ResponseMessage = err.Error()
		return echo.NewHTTPError(http.StatusBadRequest, result)
	}
	err = c.Validate(request)	
	if err != nil {
		result.ResponseMessage = err.Error()
		return echo.NewHTTPError(http.StatusBadRequest, result)
	}
	// TODO: Some validation will add.
	dto := request.ToDto()
	err = receiver.service.Register(dto)
	if err != nil {
		result.ResponseMessage = err.Error()
		return echo.NewHTTPError(http.StatusBadRequest, result)
	}
	result.Success = true
	result.ResponseMessage = application.RegisterSuccessMessage
	return c.JSON(http.StatusCreated, result)
}

// registerCustomer godoc
// @Summary Login
// @Tags customer
// @Param request body LoginCustomerRequest true "LoginCustomerRequest"
// @Success 201 {object} BaseCustomerResponse
// @Failure 400 {object} BaseCustomerResponse
// @Router /api/v1/customer/login	[post]
func (receiver *resource) login(c echo.Context) error {
	c.Echo().Validator = &CustomValidator{validator: validator.New()}
	request := new(LoginCustomerRequest)
	result := BaseCustomerResponse {
		Success: false,
	}
	err := c.Bind(request)
	if err != nil {
		result.ResponseMessage = err.Error()
		return echo.NewHTTPError(http.StatusBadRequest, result)
	}
	err = c.Validate(request)
	if err != nil {
		result.ResponseMessage = err.Error()
		return echo.NewHTTPError(http.StatusBadRequest, result)
	}
	err = receiver.service.Login(request.Email, request.Password)
	if err != nil {
		result.ResponseMessage = err.Error()
		return echo.NewHTTPError(http.StatusBadRequest, result)
	}
	result.Success = true
	result.ResponseMessage = application.LoginSuccessMessage
	return c.JSON(http.StatusOK, result)
}