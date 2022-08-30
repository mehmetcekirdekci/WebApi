package controller

import (
	"fmt"
	"net/http"
	"github.com/labstack/echo/v4"
)

func MakeHandler(instance *echo.Echo, s *resource) {
	
	instance.GET("/", func(context echo.Context) error {
		return context.JSON(http.StatusOK, "Service is up.")
	})
	g := instance.Group("")
	baseUrl := "api/v1/customer"

	g.POST(fmt.Sprintf("%s", baseUrl), s.registerCustomer)
	g.POST(fmt.Sprintf("%s/login", baseUrl), s.login)
}