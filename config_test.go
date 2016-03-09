package main

import (
	"testing"
)

func TestConfigCheck(t *testing.T) {
	var cases = map[string]bool{
		"foo\nbar":              false,
		"foo372079560813168bar": true,
	}

	c := Config{Regex: DefaultRegex}
	for value, expected := range cases {
		if actual := c.Check(value); actual != expected {
			t.Errorf("%s failed. Expected %b but received %b instead.", value, expected, actual)
		}
	}
}
