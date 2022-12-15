package tpl

import (
	"errors"
	"strings"
)

type Text struct {
	Text []string
}

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

func NewFromFile(file string) (*Text, error) {
	s := file // TODO: Getting text from a file

	return NewFromString(s)
}
