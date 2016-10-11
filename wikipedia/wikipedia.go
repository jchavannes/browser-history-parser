package wikipedia

import (
	"regexp"
	"strings"
)

func IsWikipediaUrl(url string) bool {
	regular := strings.HasPrefix(url, "https://en.wikipedia.org")
	mobile := strings.HasPrefix(url, "https://en.m.wikipedia.org")
	return regular || mobile
}

func ArticleNameFromUrl(url string) string {
	re := regexp.MustCompile("wiki/([^#]*)")
	matches := re.FindStringSubmatch(url)
	if len(matches) >= 2 {
		return matches[1]
	}
	return ""
}
