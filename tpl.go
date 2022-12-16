package tpl

import (
	"errors"
	"regexp"
	"sort"
	"strings"
)

// Splits the string around each instance of one or more consecutive white space characters,
// as defined by unicode.IsSpace,
// returning a slice of substrings or an empty slice if s contains only white space.
func Split(t string) ([]string, error) {
	if len(t) == 0 {
		return nil, errors.New("Error! Empty string.")
	}

	text := strings.Fields(string(t))

	if len(text) == 0 {
		return nil, errors.New("Error! String contains only white space.")
	}

	return text, nil
}

// Removes special characters from the text.
func SpecCharRemover(text []string, mask string) []string {
	var re *regexp.Regexp
	switch mask {
	case "all":
		re = regexp.MustCompile(`[[:punct:]]`)
	case "quotes":
		// TODO: “”‘«»„“
	case "exclamation_mark":
		// TODO !
	case "question_mark":
		// TODO ?
	case "plus":
		// TODO +
	case "minus":
		// TODO -
	}

	for i, word := range text {
		text[i] = re.ReplaceAllString(word, "")
	}

	return text
}

func ToLowercase(text []string) []string {
	for i, word := range text {
		text[i] = strings.ToLower(word)
	}

	return text
}

func UniCounter(text []string) (c int, u []string) {
	text = SpecCharRemover(text, "all")
	text = ToLowercase(text)

	for _, word := range text {
		i := sort.Search(len(u), func(i int) bool { return word <= u[i] })
		if i < len(u) && u[i] == word {
			continue
		} else {
			u = append(u, word)
			c++
		}
	}

	return
}
