package _6__Remove_Duplicates_from_Sorted_Array

func removeDuplicates(nums []int) int {
	var index int = 0
	for i, num := range nums {
		if nums[index] != num {
			index = index + 1
			nums[index] = nums[i]
		}
	}
	return index + 1
}
