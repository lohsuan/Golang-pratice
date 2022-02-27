package main

import (
	"strings"
	"unicode"
)

func WordCount(s string) map[string]int {
	result := make(map[string]int)
	f := func(c rune) bool {
		return !unicode.IsLetter(c)
	}
	words := strings.FieldsFunc(s, f)
	for _, word := range words {
		for _, letter := range word {
			result[string(letter)]++
		}
	}
	return result
}
