package _09__Fibonacci_Number_

func fib(n int) int {
	var lastVal int = 1
	var val int = 1
	var sum int = 0

	if n == 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}

	for ; n > 2; n-- {
		sum = lastVal + val
		lastVal = val
		val = sum
	}
	return sum

}
