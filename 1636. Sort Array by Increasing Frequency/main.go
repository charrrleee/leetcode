package _636__Sort_Array_by_Increasing_Frequency

import (
	"sort"
)

type Pair struct {
	Key   int
	Value int
}

func frequencySort(nums []int) []int {
	var dict = make(map[int]int, len(nums))

	var sorted []Pair = []Pair{}

	for _, num := range nums {
		_, ok := dict[num]
		if ok {
			dict[num] = dict[num] + 1
		} else {
			dict[num] = 1
		}
	}

	for num, counter := range dict {
		sorted = append(sorted, Pair{Key: num, Value: counter})
	}

	sort.Slice(sorted, func(i, j int) bool {
		pair1 := sorted[j]
		pair2 := sorted[i]
		if pair1.Value == pair2.Value && pair1.Key < pair2.Key {
			return true
		}
		return pair1.Value > pair2.Value
	})

	var newSortedList []int = []int{}
	for _, pair := range sorted {
		for pair.Value > 0 {
			newSortedList = append(newSortedList, pair.Key)
			pair.Value--
		}
	}

	return newSortedList

}
