package main

import (
	"regexp"
	"strconv"
)

func Check(original string) bool {
	re, err := regexp.Compile("\\D")
	if err != nil {
		panic(err)
	}

	cleaned := re.ReplaceAllString(original, "")
	sum := 0
	alt := false
	size := len(cleaned)

	if size < 13 || size > 19 {
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
