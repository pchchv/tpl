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

// Gets an array of strings and a special characters mask,
// removes special characters from each string.
// Returns the modified array
// Mask options:
//    "all" - Remove all special characters (!"#$%&'()*+,-./:;<=>?@[]^_`{|}~\)
//    "quotes" - Remove quotes (”‘«»„)
//    "exclamation_mark" - Remove exclamation marks (!)
//    "question_mark" - Remove question marks (?)
//    "plus" - Remove pluses (+)
//    "minus" - Remove minuses (-)
//    "brackets" - Remove brackets ((){}[])
func SpecCharRemover(text []string, mask string) ([]string, error) {
	var re *regexp.Regexp
	switch mask {
	case "all":
		re = regexp.MustCompile(`[[:punct:]]`)
	case "quotes":
		re = regexp.MustCompile(`[“”‘«»„“]`)
	case "exclamation_mark":
		re = regexp.MustCompile(`!`)
	case "question_mark":
		re = regexp.MustCompile(`\?`)
	case "plus":
		re = regexp.MustCompile(`\+`)
	case "minus":
		re = regexp.MustCompile(`-`)
	case "brackets":
		re = regexp.MustCompile(`[\(\)\{\}\[\]]`)
	default:
		return nil, errors.New("Enter the correct character set")

	}

	for i, word := range text {
		text[i] = re.ReplaceAllString(word, "")
	}

	return text, nil
}

// Gets an array of strings, converts all strings to lowercase.
// Returns the modified array
func ToLowercase(text []string) []string {
	for i, word := range text {
		text[i] = strings.ToLower(word)
	}

	return text
}

// Gets an array of strings, counts the number of unique strings.
// Returns the number of unique strings and an array of them.
func UniCounter(text []string) (c int, u []string) {
	text, err := SpecCharRemover(text, "all")
	if err != nil {
		panic(err)
	}
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

// Gets an array of strings, joins them separated by spaces.
// Returns a string
func Build(t []string) string {
	var text string

	for _, word := range t {
		text += word + " "
	}

	return strings.TrimSpace(text)
}
