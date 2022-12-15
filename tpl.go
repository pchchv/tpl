package tpl

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
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
func (t *Text) SpecCharRemover() {
	re := regexp.MustCompile(`[[:punct:]]`)
	for i, word := range t.Text {
		t.Text[i] = re.ReplaceAllString(word, "")
	}
	// TODO: Removing only necessary characters
}
