package url

import (
	"net/http"
	"regexp"
	"io"
)

// The struct UrlService is a DomainService.
type UrlService struct {
	url *Url
}

func NewUrlService(url *Url) *UrlService {
	i := new(UrlService)
	i.url = url
	return i
}

func (us *UrlService) FetchSummary() ([]string, error) {
	addr := us.url.Addr()

	re, err := http.Get(addr)
	if err != nil {
		return nil, err
	}
	defer re.Body.Close()

	body, err := io.ReadAll(re.Body)
	if err != nil {
		return nil, err
	}
	bodyStr := string(body)

	var summaries []string

	funcClearTag := func(s string) string {
		re := regexp.MustCompile(`<.*?>`)
		return re.ReplaceAllString(s, "")
	}

	// Find <title>
	rgTitle := regexp.MustCompile(`(?i)<\s*title.*>.+<\s*/title\s*>`)
	if rgTitle.MatchString(bodyStr) {
		s := rgTitle.FindString(bodyStr)
		summaries = append(summaries, "title :"+funcClearTag(s))
	}

	// Find <h1>
	rgH1 := regexp.MustCompile(`(?i)<\s*h1.*>.+<\s*/h1\s*>`)
	if rgH1.MatchString(bodyStr) {
		s := rgH1.FindString(bodyStr)
		summaries = append(summaries, "H1 :"+funcClearTag(s))
	}

	return summaries, nil
}
