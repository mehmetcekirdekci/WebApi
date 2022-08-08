/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/mehmetcekirdekci/WebApi/app/customer/application"
	customercontroller "github.com/mehmetcekirdekci/WebApi/app/customer/application/controller"
	"github.com/mehmetcekirdekci/WebApi/app/customer/domain/repositories"
	echoextention "github.com/mehmetcekirdekci/WebApi/pkg/echoExtention"
	"github.com/spf13/cobra"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type api struct {
	instance *echo.Echo
	command *cobra.Command
	Port string
}

// apiCmd represents the api command
var apiCmd = &api{
	command: &cobra.Command{
		Use: "api",
		Short: "Service is up.",
		Long: "Golang api project is up.",
	},
	Port: "5000",
}

func init() {
	RootCommand.AddCommand(apiCmd.command)
	apiCmd.instance = echo.New()

	apiCmd.command.RunE = func(cmd *cobra.Command, args []string) error {
		var err error
		var db *gorm.DB
		var dsn string
		var customerRepo repositories.CustomerRepository
		var accountInformationRepo repositories.AccountInformationRepository
		var customerService application.Service
		dsn = "host=localhost user=postgres password=5492 dbname=GolangApiExample port=5432 sslmode=disable"

		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			println(err)
			return err
		}

		customerRepo = repositories.NewCustomerRepository(db)

		accountInformationRepo = repositories.NewAccountInformationRepository(db)

		customerService = application.NewService(customerRepo, accountInformationRepo)

		customercontroller.MakeHandler(apiCmd.instance, customercontroller.NewController(customerService))

		apiCmd.instance.GET("/swagger/*", echoSwagger.WrapHandler)

		go func ()  {
			err := apiCmd.instance.Start(fmt.Sprint("%s", apiCmd.Port))
			if err != nil {
				println(err)
			}
		}()
		echoextention.Shutdown(apiCmd.instance, 2*time.Second)

		return nil
	}
}
