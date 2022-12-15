package tpl

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

type Text struct {
	Text []string
}

// Gets a string, splits it into words (any whitespace characters).
// Returns the Text structure with the Text field as a list of strings consisting of words.
func NewFromString(s string) (*Text, error) {
	t := &Text{
		Text: strings.Fields(s),
	}

	if len(t.Text) == 1 {
		return nil, errors.New("Error! Expected: text, received: word.")
	} else if len(t.Text) == 0 {
		return nil, errors.New("Error! Expected: text, received: empty string.")
	}

	return t, nil
}

// Gets the path to the file, creates its abstract representation, returns the NewFromString function.
func NewFromFile(file string) (*Text, error) {
	filePath, err := filepath.Abs(file)
	if err != nil {
		return nil, errors.New("Error! Incorrect path.")
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("Error when reading a file: %v", err)
	}

	return NewFromString(string(data))
}

// Removes special characters from the text.
func (t *Text) SpecCharRemover(mask string) {
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

	for i, word := range t.Text {
		t.Text[i] = re.ReplaceAllString(word, "")
	}
}

func (t *Text) UniCounter() (c int, u []string) {
	text := &Text{
		Text: t.Text,
	}

	text.SpecCharRemover("all")
	// TODO: Reducing all words to lowercase

	for _, word := range t.Text {
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
