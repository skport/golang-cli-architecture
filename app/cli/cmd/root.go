// Package cmd is command controller with cobra.

package cmd

import (
	"fmt"
	"os"

	"webfetcher/core/app"
	"webfetcher/core/url"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"go.uber.org/dig"
)

var (
	diContainer dig.Container
)

var rootCmd = &cobra.Command{
	Use:   "golang-webfetcher",
	Short: "A simple web fetcher for golang",
	Long:  `Description`,
}

func init() {
	// Initialize Enviroment
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	env := os.Getenv("APP_ENV")
	if "" == env {
		env = "development"
	}

	// Initialize DI container
	diContainer = *dig.New()
	diContainer.Provide(app.NewApp)

	// Register UrlProvider to use with Container
	// Switch data source external or dummy
	if env == "production" {
		diContainer.Provide(url.NewWebProvider)
		return
	}
	diContainer.Provide(url.NewInMemDummyProvider)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
