package helper

import (
	"regexp"
)

func FilterHtmlTag(filter_tag, html string) string {
	if filter_tag == "" {
		filter_tag = "p|div|h[1-6]|blockquote|pre|table|dl|ol|ul|script|noscript|form|fieldset|iframe|math|ins|del"
	}

	return regexp.MustCompile(`<[/]?(`+filter_tag+`)[^>]*>[\n]?`).ReplaceAllString(html, "")
}
