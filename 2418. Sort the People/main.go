package _418__Sort_the_People

import "sort"

func sortPeople(names []string, heights []int) []string {
	var dict map[int]string = make(map[int]string, len(names))

	for index, name := range names {
		height := heights[index]
		dict[height] = name
	}

	var newList []int

	for height, _ := range dict {
		newList = append(newList, height)
	}

	// todo dont use library
	sort.Slice(newList, func(i, j int) bool {
		return newList[j] < newList[i]
	})

	var newNameList []string

	for _, height := range newList {
		newNameList = append(newNameList, dict[height])
	}

	return newNameList
}
