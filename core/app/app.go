// Package app is Business Logic layer.
// Receives instructions from the controller and performs application logic.
package app

import (
	"fmt"
	"regexp"

	"webfetcher/core/url"
)

type App struct{
	urlProvider url.Provider
}

func NewApp() *App {
	a := new(App)
	a.init()
	return a
}

func (a *App) init() {
	// Select url provider to use
	a.urlProvider = url.NewWebProvider()
} 

func (a *App) CmdSummary(args []string) {
	addr := args[0]

	// Create Url Instance
	u, err := url.NewUrl(addr)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Load Body from Url
	body, err := a.urlProvider.ReadBody(u)
	if err != nil {
		fmt.Println(err)
		return
	}

	var summaries []string

	funcClearTag := func(s string) string {
		re := regexp.MustCompile(`<.*?>`)
		return re.ReplaceAllString(s, "")
	}

	// Find <title>
	rgTitle := regexp.MustCompile(`(?i)<\s*title.*>.+<\s*/title\s*>`)
	if rgTitle.MatchString(body) {
		s := rgTitle.FindString(body)
		summaries = append(summaries, "title :"+funcClearTag(s))
	}

	// Find <h1>
	rgH1 := regexp.MustCompile(`(?i)<\s*h1.*>.+<\s*/h1\s*>`)
	if rgH1.MatchString(body) {
		s := rgH1.FindString(body)
		summaries = append(summaries, "H1 :"+funcClearTag(s))
	}

	// Print Summary
	fmt.Println(summaries)
}
