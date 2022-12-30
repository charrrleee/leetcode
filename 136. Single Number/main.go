package _36__Single_Number

func singleNumber(nums []int) int {
	mapping := make(map[int]int)

	for _, num := range nums {
		_, ok := mapping[num]
		if ok {
			delete(mapping, num)
		} else {
			mapping[num] = 1
		}
	}

	for key, value := range mapping {
		if value == 1 {
			return key
		}
	}
	return 0
}
