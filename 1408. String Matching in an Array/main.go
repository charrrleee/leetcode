package _408__String_Matching_in_an_Array

import "strings"

func stringMatching1(words []string) []string {
	var slice []string
	for i, w1 := range words {
		for j, w2 := range words {
			if i == j {
				continue
			}
			originalSize := len(w1)
			newString := strings.ReplaceAll(w1, w2, "")
			newSize := len(newString)
			if originalSize != newSize && !contains(slice, w2) {
				slice = append(slice, w2)
			}
		}
	}
	return slice
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func stringMatching2(words []string) []string {
	var slice []string
	var data = make(map[string]bool)
	for i, w1 := range words {
		for j, w2 := range words {
			if i == j {
				continue
			}

			if strings.Contains(w1, w2) {
				_, ok := data[w2]
				if !ok {
					slice = append(slice, w2)
					data[w2] = true

				}
			}
		}
	}
	return slice
}
