package _137__N_th_Tribonacci_Number_

func tribonacci(n int) int {
	var one int = 0
	var two int = 1
	var three int = 1
	var sum int = 0

	if n == 0 {
		return 0
	}

	if n <= 2 {
		return 1
	}

	for ; n > 2; n-- {
		sum = one + two + three
		one = two
		two = three
		three = sum
	}
	return sum
}
