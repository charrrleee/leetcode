package _0__Valid_Parentheses

func isValid(s string) bool {
	var size int = len(s)
	stack := make([]rune, 0, len(s))
	hashMap := map[rune]rune{
		'}': '{',
		')': '(',
		']': '[',
	}

	if size%2 == 1 {
		return false
	}

	for _, bracket := range s {
		openBracket, ok := hashMap[bracket]
		if !ok {
			stack = append(stack, bracket)
		} else {
			if len(stack) == 0 || (stack[len(stack)-1] != openBracket) {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}
