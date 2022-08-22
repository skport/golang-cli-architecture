package fetcher

import (
	"bufio"
	"fmt"
	"net/http"
	"regexp"

	"fetchweb/core/addr"
)

type Fetcher struct {
	addr addr.Addr
}

func NewFetcher(url string) (*Fetcher, error) {
	f := new(Fetcher)

	f.addr = *addr.NewAddr(url)
	err := f.addr.Validate()
	if err != nil {
		return f, err
	}

	return f, nil
}

func (f *Fetcher) PrintSummary() error {
	addr := f.addr.GetAddr()

	re, err := http.Get(addr)
	if err != nil {
		return err
	}
	defer re.Body.Close()

	funcClearTag := func(s string) string {
		re := regexp.MustCompile(`<.*?>`)
		return re.ReplaceAllString(s, "")
	}

	// Read the body line by line
	bf := bufio.NewScanner(re.Body)
	rgTitle := regexp.MustCompile(`(?i)<\s*title.*>.+<\s*/title\s*>`)
	rgH1 := regexp.MustCompile(`(?i)<\s*h1.*>.+<\s*/h1\s*>`)
	for bf.Scan() {
		line := bf.Text()
		// Extract <title>
		if rgTitle.MatchString(line) {
			s := rgTitle.FindString(line)
			fmt.Println("title :", funcClearTag(s))
		}
		// Extract <h1>
		if rgH1.MatchString(line) {
			s := rgH1.FindString(line)
			fmt.Println("H1 :", funcClearTag(s))
		}
	}
	if err := bf.Err(); err != nil {
		return err
	}

	return nil
}
