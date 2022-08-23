package core

import (
	"flag"
	"fmt"

	"webfetcher/core/fetcher"
)

func App() {
	// Input parameters of command line
	var (
		addr = flag.String("a", "https://github.com", "Target URL.")
	)
	flag.Parse()

	// Create Fetcher Object
	f, err := fetcher.NewFetcher(*addr)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Show summary of page
	err = f.PrintSummary()
	if err != nil {
		fmt.Println(err)
		return
	}
}
