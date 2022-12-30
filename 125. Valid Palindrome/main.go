package _25__Valid_Palindrome

import (
	"strings"
)

func isPalindrome(s string) bool {
	newString := ""
	reversedStr := ""

	for _, bWord := range s {
		// filter out the non ascii alphanumeric characters
		if (bWord < 'a' || bWord > 'z') && (bWord < 'A' || bWord > 'Z') && (bWord < '0' || bWord > '9') {
			continue
		}

		newString = newString + string(bWord)
		reversedStr = string(bWord) + reversedStr
	}

	return strings.ToUpper(newString) == strings.ToUpper(reversedStr)
}
