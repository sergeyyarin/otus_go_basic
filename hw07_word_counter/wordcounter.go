package wordcounter

import (
	"strings"
	"unicode"
)

func CountWords(s string) (m map[string]int) {
	m = make(map[string]int)

	filter := func(r rune) bool {
		return unicode.IsControl(r) || unicode.IsSpace(r) || unicode.IsPunct(r)
	}

	split := strings.FieldsFunc(s, filter)

	for _, word := range split {
		m[word]++
	}

	return m
}
