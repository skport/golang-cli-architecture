package service

import (
	"bufio"
	"fmt"
	"net/http"
	"regexp"

	"webfetcher/core/url"
)

type Service struct {
	url url.Url
}

func NewService(addr string) (*Service, error) {
	s := new(Service)

	s.url = *url.NewUrl(addr)
	err := s.url.Validate()
	if err != nil {
		return s, err
	}

	return s, nil
}

func (s *Service) Execute() error {
	addr := s.url.GetAddr()

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
