package __Palindrome_Number

import "strconv"

func isPalindrome1(x int) bool {
	var newStr string

	var xStr = strconv.Itoa(x)
	for _, s := range xStr {
		newStr = string(s) + newStr
	}
	newVal, _ := strconv.Atoi(newStr)
	return x == newVal
}

func isPalindrome2(x int) bool {
	if x < 0 {
		return false
	}

	var reversedValue int = 0
	var target int = x

	for target != 0 {
		reversedValue = reversedValue * 10
		reversedValue = reversedValue + target%10
		target = target / 10
	}

	if x == reversedValue {
		return true
	}
	return false
}
