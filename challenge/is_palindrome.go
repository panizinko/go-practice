package challenge

import (
	"regexp"
	"strings"
)

func IsPalindrome(str string) bool {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		return false
	}

	cleanStr := reg.ReplaceAllString(strings.ToLower(str), "")
	return cleanStr == ReverseString(cleanStr)
}
