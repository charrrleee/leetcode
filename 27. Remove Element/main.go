package _7__Remove_Element

func removeElement(nums []int, val int) int {
	var index = 0
	for _, num := range nums {
		if num != val {
			nums[index] = num
			index = index + 1
		}
	}
	return index
}
