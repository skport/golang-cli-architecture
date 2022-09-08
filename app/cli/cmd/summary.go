// Controller Layer : Cmd Summary

package cmd

import (
	"fmt"

	"webfetcher/core/app"

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

	// Application Logic
	// Create app Instance via DI Container
	diContainer.Invoke(
		func(a *app.App) {
			a.CmdSummary(args)
		},
	)
}
