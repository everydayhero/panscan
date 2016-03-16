package main

import (
	"regexp"
	"strconv"
)

func Check(original string) bool {
	re := regexp.MustCompile("\\D")

	cleaned := re.ReplaceAllString(original, "")
	sum := 0
	alt := false
	size := len(cleaned)

	if size < 13 || size > 19 || len(condense(cleaned)) <= 1 {
		return false
	}

	for i := size - 1; i > -1; i-- {
		mod, _ := strconv.Atoi(string(cleaned[i]))
		if alt {
			mod *= 2
			if mod > 9 {
				mod = (mod % 10) + 1
			}
		}
		alt = !alt
		sum += mod
	}

	return sum > 0 && sum%10 == 0
}

func condense(s string) string {
	b := []byte{}

	for i := 0; i < len(s); i++ {
		l := len(b)
		if l == 0 || s[i] != b[l-1] {
			b = append(b, s[i])
		}
	}

	return string(b)
}
