/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
	"os"

	"github.com/mehmetcekirdekci/WebApi/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type RootCmd struct {
	configFile string
	AppConfig *config.Configuration
	cobraCommand *cobra.Command
}

var RootCommand = RootCmd{
	cobraCommand: &cobra.Command{
		Use: "webapi",
		Short: "Service is up.",
		Long: "Golang api example service is up.",
	},
	AppConfig: &config.Configuration{},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := RootCommand.cobraCommand.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCommand.cobraCommand.PersistentFlags().StringVarP(&RootCommand.configFile, "config", "c", "config.qa.yaml", "")
}

func initConfig() {
	if RootCommand.configFile != "" {
		viper.SetConfigFile(RootCommand.configFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("webapi.json")
	}

	viper.Set("Verbose", true)
	viper.SetConfigType("json")

	err := viper.ReadInConfig()
	if err != nil {
		RootCommand.cobraCommand.Help()
	} else {
		err = viper.Unmarshal(RootCommand.AppConfig)
		if err != nil {
			errors.New(viper.GetViper().ConfigFileUsed())
		}
	}
}

func (c *RootCmd) AddCommand(cmd *cobra.Command)  {
	c.cobraCommand.AddCommand(cmd)
}