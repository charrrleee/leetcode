package _5__Search_Insert_Position

func searchInsert1(nums []int, target int) int {
	var lastIdx = 0
	for i, num := range nums {
		if num >= target {
			return i
		}
		lastIdx = i
	}
	return lastIdx + 1
}

func searchInsert2(nums []int, target int) int {
	var left int = 0
	var right int = len(nums) - 1

	for left <= right {
		mid := (left + right) / 2
		num := nums[mid]

		if target < num {
			right = mid - 1
		} else {
			left = mid + 1
		}

		if target == num {
			return mid
		}
	}
	return left
}
