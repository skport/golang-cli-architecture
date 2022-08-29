// DomainService : Url

package url

import (
	"bufio"
	"net/http"
	"regexp"
)

type UrlService struct{}

func NewUrlService() *UrlService {
	i := new(UrlService)
	return i
}

func (s *UrlService) FetchSummary(url *Url) ([]string, error) {
	addr := url.GetAddr()

	re, err := http.Get(addr)
	if err != nil {
		return nil, err
	}
	defer re.Body.Close()

	funcClearTag := func(s string) string {
		re := regexp.MustCompile(`<.*?>`)
		return re.ReplaceAllString(s, "")
	}

	var summaries []string

	// Read the body line by line
	bf := bufio.NewScanner(re.Body)
	rgTitle := regexp.MustCompile(`(?i)<\s*title.*>.+<\s*/title\s*>`)
	rgH1 := regexp.MustCompile(`(?i)<\s*h1.*>.+<\s*/h1\s*>`)
	for bf.Scan() {
		line := bf.Text()
		// Extract <title>
		if rgTitle.MatchString(line) {
			s := rgTitle.FindString(line)
			summaries = append(summaries, "title :"+funcClearTag(s))
		}
		// Extract <h1>
		if rgH1.MatchString(line) {
			s := rgH1.FindString(line)
			summaries = append(summaries, "H1 :"+funcClearTag(s))
		}
	}
	if err := bf.Err(); err != nil {
		return nil, err
	}

	return summaries, nil
}
