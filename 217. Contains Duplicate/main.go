package _17__Contains_Duplicate

func containsDuplicate(nums []int) bool {
	mapping := make(map[int]bool)

	for _, num := range nums {
		_, ok := mapping[num]
		if ok {
			return true
		} else {
			mapping[num] = true
		}
	}
	return false
}
