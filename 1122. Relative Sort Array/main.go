package _122__Relative_Sort_Array

import "sort"

func relativeSortArray(arr1 []int, arr2 []int) []int {
	var dict = make(map[int]int)

	for _, val := range arr1 {
		_, ok := dict[val]
		if ok {
			dict[val] = dict[val] + 1
		} else {
			dict[val] = 1
		}
	}

	var newList []int

	for _, val := range arr2 {
		counter, _ := dict[val]
		for counter > 0 {
			newList = append(newList, val)
			counter--
		}
		delete(dict, val)
	}
	var sortRemainList []int

	for key, _ := range dict {
		sortRemainList = append(sortRemainList, key)
	}

	// todo better do not use library
	sort.Ints(sortRemainList)

	for _, key := range sortRemainList {
		counter := dict[key]
		for counter > 0 {
			newList = append(newList, key)
			counter--
		}
		delete(dict, key)
	}

	return newList
}
