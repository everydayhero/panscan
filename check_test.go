package main

import (
	"testing"
)

func TestCheck(t *testing.T) {
	var cases = map[string]bool{
		"3720-7956-0813-168":  true,
		"3615-7344-819-566":   true,
		"6011 9479 8460 8576": true,
		"2149 1844 5574 804":  true,
		"3088412336401191":    true,
		"5481223188541697":    true,
		"4539524757964623":    true,
		"869934845122647":     true,
		"6706661850803758147": true,
		"1111111111111111":    false,
		"1234123412341234":    false,
		"0000000000000000":    false,
		"376031615608":        false,
		"8888888888888888":    false,
	}

	for value, expected := range cases {
		if actual := Check(value); actual != expected {
			t.Errorf("%s failed. Expected %b but received %b instead.", value, expected, actual)
		}
	}
}
