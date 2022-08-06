package controller

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

)

func MakeHandler(instance *echo.Echo, s *resource) {

	instance.Use(middleware.Logger())
	instance.Use(middleware.Recover())
	instance.GET("/swagger/*", echoSwagger.WrapHandler)

	instance.GET("/", func(context echo.Context) error {
		return context.JSON(http.StatusOK, "Service is up.")
	})
	g := instance.Group("")
	baseUrl := "api/v1/customer"

	g.POST(fmt.Sprintf("%s", baseUrl), s.registerCustomer)
	instance.Logger.Fatal(instance.Start(":5000"))
}