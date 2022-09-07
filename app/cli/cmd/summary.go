// Controller Layer : Cmd Summary

package cmd

import (
	"fmt"

	"webfetcher/core/app"
	"webfetcher/core/url"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(SummaryCmd)
}

var SummaryCmd = &cobra.Command{
	Use:   "summary [URL]",
	Short: "Print a summary from a web",
	Long:  `Print a summary from a web`,
	Run:   SummaryCmdRun,
}

func SummaryCmdRun(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		fmt.Println("invalid URL specified")
		return
	}

	// Select Url Provider
	urlProvider := url.NewWebProvider()

	// Application Logic
	a := app.NewApp(urlProvider) // DI
	a.CmdSummary(args)
}
