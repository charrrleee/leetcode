package _9__Sqrt_x_

func mySqrt(x int) int {
	var min int = 0
	var max int = x

	for min <= max {
		var mid int = (min + max) / 2
		if mid*mid == x {
			return x
		} else if mid*mid < x {
			min = mid + 1
		} else {
			max = mid - 1
		}
	}
	return max
}

// I think it is not a better answer cause if x is a large number, it is not a good idea to loop through the value
func mySqrt1(x int) int {
	var value = 0

	for value*value < x {
		value++
	}

	if value*value > x {
		value--
	}

	return value
}

// Copy from leetcode which is easy to understand but first u should know Newton's method
func mySqrt2(x int) int {
	var value int = x
	for value*value > x {
		value = (value + x/value) / 2
	}
	return value
}
