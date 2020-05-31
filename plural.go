package plural

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

var re = regexp.MustCompile(".*[^aiueo]y")

func Convert(word string) string {
	switch {
	case strings.HasSuffix(word, "o"),
		strings.HasSuffix(word, "s"),
		strings.HasSuffix(word, "x"),
		strings.HasSuffix(word, "ch"),
		strings.HasSuffix(word, "sh"):
		return word + "es"
	case strings.HasSuffix(word, "f"):
		length := utf8.RuneCountInString(word)
		return string([]rune(word)[:length-1]) + "ves"
	case strings.HasSuffix(word, "fe"):
		length := utf8.RuneCountInString(word)
		return string([]rune(word)[:length-2]) + "ves"
	case re.MatchString(word):
		length := utf8.RuneCountInString(word)
		return string([]rune(word)[:length-1]) + "ies"
	default:
		return word + "s"
	}
}
