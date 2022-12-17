package tpl

import (
	"strings"
	"testing"
)

const testString = "Lorem ipsum dolor sit amet, consectetur adipiscing elit! " +
	"Etiam ac convallis risus. Ut accumsan urna sem, in placerat mi luctus a. " +
	"Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae; " +
	"Morbi eu massa in nulla rutrum maximus vitae id massa. Aenean venenatis, nunc nec cursus porta, ex lorem egestas erat, ut."

var punctuation = strings.Split("!\"#$%&'()*+,-./:;<=>?@[]^_`{|}~\\", "")

func TestSplit(t *testing.T) {
	text, err := Split(testString)
	if err != nil {
		t.Fatal(err)
	}

	if len(text) <= 1 {
		t.Fatal("Text split error")
	}
}

func TestSpecCharRemoverAll(t *testing.T) {
	text, err := Split(testString)
	if err != nil {
		t.Fatal(err)
	}

	text, err = SpecCharRemover(text, "all")
	if err != nil {
		t.Fatal(err)
	}

	for _, word := range text {
		for _, symbol := range punctuation {
			if strings.Contains(word, symbol) {
				t.Fatal("The string contains a punctuation symbol")
			}
		}
	}
}

func TestSpecCharRemoverWrongInput(t *testing.T) {
	text, err := Split(testString)
	if err != nil {
		t.Fatal(err)
	}

	_, err = SpecCharRemover(text, "test")
	if err == nil {
		t.Fatal("Attempting to return a result with incorrect input")
	}
}

func TestSpecCharRemoverExclamation(t *testing.T) {
	text, err := Split(testString)
	if err != nil {
		t.Fatal(err)
	}

	text, err = SpecCharRemover(text, "exclamation_mark")
	if err != nil {
		t.Fatal(err)
	}

	for _, word := range text {
		if strings.Contains(word, "!") {
			t.Fatal("The string contains a exclamation mark")
		}
	}
}

func TestUniCounter(t *testing.T) {
	text, err := Split(testString)
	if err != nil {
		t.Fatal(err)
	}

	count, unique := UniCounter(text)

	if count != len(unique) {
		t.Fatal("Error in calculating the quantity.")
	}
}

func TestToLowercase(t *testing.T) {
	text, err := Split(testString)
	if err != nil {
		t.Fatal(err)
	}

	ToLowercase(text)

	if text[0] == "Lorem" {
		t.Fatal("The word reductions didn't work.")
	}
}

func TestBuild(t *testing.T) {
	text, err := Split(testString)
	if err != nil {
		t.Fatal(err)
	}

	stringText := Build(text)

	if len(stringText) == 0 {
		t.Fatal("Build error. Empty string in result")
	}

	if testString != stringText {
		t.Fatal("Build Error. testString is different from result string")
	}
}
