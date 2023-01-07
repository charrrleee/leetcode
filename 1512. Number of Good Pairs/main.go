package _512__Number_of_Good_Pairs

// if a number show more than once, the value will be 1 + 2 + x == 0 + 1 + x
// if 1 shows 3 times, the pair counter will be 2
// if 1 shows 4 times, the pair counter wil be 3 + 2 + 1

func numIdenticalPairs(nums []int) int {
	mapping := make(map[int]int)
	var counter int

	for _, num := range nums {
		counter = counter + mapping[num]
		mapping[num]++
	}
	return counter
}
