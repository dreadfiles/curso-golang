package greetings

import (
	"regexp"
	"strings"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Adrian"
	expectedRegex := regexp.MustCompile(`\b` + name + `\b`)
	response, err := Hello("Adrian")
	if !expectedRegex.MatchString(response) || err != nil {
		t.Fatalf("no match strings")
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf("this must be empty and error")
	}
}

func TestTableDriven(t *testing.T) {
	var tests = []struct {
		name  string
		input string
	}{
		{"name ok", "Adrian"},
		{"name empty", ""},
	}

	for _, item := range tests {
		t.Run(item.name, func(t *testing.T) {
			response, err := Hello(item.input)

			if err == nil && !strings.Contains(response, item.input) {
				t.Error("not match string")
			}

			if err != nil && item.input != "" {
				t.Error("not error controlled")
			}
		})
	}
}
