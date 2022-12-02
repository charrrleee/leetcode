package _8__Length_of_Last_Word_

import "strings"

func lengthOfLastWord1(s string) int {
	split := strings.Split(strings.TrimSpace(s), " ")
	size := len(split)
	lastWord := split[size-1]
	return len(lastWord)
}

func lengthOfLastWord2(s string) int {
	split := strings.Fields(s) //tokenize
	size := len(split)
	return len(split[size-1])
}

func lengthOfLastWord3(s string) int {
	word := ""
	l := len(s)

	for i, _ := range s {
		index := l - i - 1
		str := string(s[index])

		if str == " " && word != "" {
			return len(word)
		}
		if str != " " {
			word = str + word
		}
		if str == " " {
			continue
		}
	}
	return len(word)
}
