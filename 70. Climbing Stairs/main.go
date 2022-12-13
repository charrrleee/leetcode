package _0__Climbing_Stairs

func climbStairs(n int) int {
	var value int = 3
	var lastValue int = 2
	var sum int = 0

	if n <= 3 {
		return n
	}

	for ; 3 < n; n-- {
		sum = value + lastValue
		lastValue = value
		value = sum
	}

	return sum
}
