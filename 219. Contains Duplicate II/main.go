package _19__Contains_Duplicate_II

func containsNearbyDuplicate(nums []int, k int) bool {
	mapping := make(map[int]int)

	for i, num := range nums {
		j, ok := mapping[num]

		if !ok {
			mapping[num] = i
			continue
		}

		diff := i - j

		if k >= diff {
			return true
		}
		mapping[num] = i
	}
	return false
}
