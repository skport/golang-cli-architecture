// Package cmd is command controller with cobra.

package cmd

import (
	"fmt"
	"os"

	"webfetcher/core/app"
	"webfetcher/core/url"

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
	// Initialize DI container
	diContainer = *dig.New()
	diContainer.Provide(app.NewApp)

	// Register UrlProvider to use with Container
	diContainer.Provide(url.NewWebProvider)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
