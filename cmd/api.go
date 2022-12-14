/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mehmetcekirdekci/WebApi/app/customer/application"
	customercontroller "github.com/mehmetcekirdekci/WebApi/app/customer/application/controller"
	"github.com/mehmetcekirdekci/WebApi/app/customer/domain/repositories"
	echoextention "github.com/mehmetcekirdekci/WebApi/pkg/echoExtention"
	"github.com/spf13/cobra"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	apiCmd.instance.Use(middleware.Logger())
	apiCmd.instance.Use(middleware.Recover())
	apiCmd.instance.GET("/swagger/*", echoSwagger.WrapHandler)

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

		mongoDb, err := MongoNewDatabase("mongodb://localhost:27017", "GolangApiExample")
		if err != nil {
			println(err)
			return err
		}

		customerRepo = repositories.NewCustomerRepository(db)

		accountInformationRepo = repositories.NewAccountInformationRepository(mongoDb)

		customerService = application.NewService(customerRepo, accountInformationRepo)

		customercontroller.MakeHandler(apiCmd.instance, customercontroller.NewController(customerService))

		apiCmd.instance.Logger.Fatal(apiCmd.instance.Start(":5000"))
		echoextention.Shutdown(apiCmd.instance, 2*time.Second)

		return nil
	}
}


func MongoNewDatabase(uri, databaseName string) (db *mongo.Database, err error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		println(err)
		return nil, err
	}
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	err = client.Connect(ctx)
	if err != nil {
		println(err)
		return nil, err
	}
	db = client.Database(databaseName)
	return db, nil
}
