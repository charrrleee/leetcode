package _3__Roman_to_Integer_

func romanToInt1(s string) int {
	var l []int
	var sum int = 0
	var size int

	for _, str := range s {
		l = append(l, transform(string(str)))
	}

	size = len(l)
	for i, str := range l {
		if i == size-1 || str >= l[i+1] {
			sum = sum + str
		} else {
			sum = sum - str
		}
	}
	return sum
}

func romanToInt2(s string) int {
	var sum int = 0
	var currentValue int = 0
	var lastValue int = 0

	for i := len(s) - 1; i >= 0; i-- {
		currentValue = transform(string(s[i]))
		if currentValue < lastValue {
			sum -= currentValue
		} else {
			sum += currentValue
		}
		lastValue = currentValue
	}
	return sum
}

func transform(s string) int {
	switch s {
	case "I":
		return 1
	case "V":
		return 5
	case "X":
		return 10
	case "L":
		return 50
	case "C":
		return 100
	case "D":
		return 500
	case "M":
		return 1000
	default:
		return 0

	}
}
