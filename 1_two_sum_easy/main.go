package __two_sum_easy

func twoSum1(nums []int, target int) []int {
	for i, value1 := range nums {
		for j, value2 := range nums {
			if i == j {
				continue
			}
			if value1+value2 == target {
				return []int{i, j}
			}
		}
	}
	return []int{0, 0}
}

func twoSum2(nums []int, target int) []int {
	m := make(map[int]int)

	for index, num := range nums {
		if value, ok := m[target-num]; ok {
			return []int{value, index}
		}
		m[num] = index
	}
	return []int{0, 0}
}
