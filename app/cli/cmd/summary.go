package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"webfetcher/core/service"
)

func init() {
	rootCmd.AddCommand(cobraCmd)
}

var cobraCmd = &cobra.Command{
	Use:   "summary",
	Short: "Print a summary from a web",
	Long:  `Print a summary from a web`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("invalid URL specified")
			return
		}

		url := args[0]

		// Create Service Object
		s, err := service.NewService(url)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Show summary of page
		err = s.Execute()
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}
