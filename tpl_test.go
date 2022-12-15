package tpl

import (
	"fmt"
	"strings"
	"testing"
)

const (
	testString = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. " +
		"Etiam ac convallis risus. Ut accumsan urna sem, in placerat mi luctus a. " +
		"Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae; " +
		"Morbi eu massa in nulla rutrum maximus vitae id massa. Aenean venenatis, nunc nec cursus porta, ex lorem egestas erat, ut."

	testFile = "./testText.txt"
)

var punctuation = strings.Split("!\"#$%&'()*+,-./:;<=>?@[]^_`{|}~\\", "")

func TestNewFromSting(t *testing.T) {
	text, err := NewFromString(testString)
	if err != nil {
		t.Fatal(err)
	}

	if fmt.Sprintf("%T", text) != "*tpl.Text" {
		t.Fatal("Incorrect type")
	}

	if len(text.Text) <= 1 {
		t.Fatal("Text split error")
	}
}

func TestWrongNewFromString(t *testing.T) {
	_, err := NewFromString("Hello")
	if err == nil {
		t.Fatal("No error when getting a word")
	}

	_, err = NewFromString("")
	if err == nil {
		t.Fatal("No error when getting a empty input")
	}
}

func TestNewFromFile(t *testing.T) {
	text, err := NewFromFile(testFile)
	if err != nil {
		t.Fatal(err)
	}

	if fmt.Sprintf("%T", text) != "*tpl.Text" {
		t.Fatal("Incorrect type")
	}

	if len(text.Text) <= 1 {
		t.Fatal("Text split error")
	}
}

func TestSpecCharRemover(t *testing.T) {
	text, err := NewFromFile(testFile)
	if err != nil {
		t.Fatal(err)
	}

	textLength := len(text.Text)

	text.SpecCharRemover("all")

	for _, word := range text.Text {
		for _, symbol := range punctuation {
			if strings.Contains(word, symbol) {
				t.Fatal("The string contains a punctuation symbol")
			}
		}
	}

	if textLength != len(text.Text) {
		t.Fatal("The words are lost")
	}
}

func TestUniCounter(t *testing.T) {
	text, err := NewFromFile(testFile)
	if err != nil {
		t.Fatal(err)
	}

	count, unique := text.UniCounter()

	if count != len(unique) {
		t.Fatal("Error in calculating the quantity")
	}
}
