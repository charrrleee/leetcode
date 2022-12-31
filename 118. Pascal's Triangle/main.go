package _18__Pascal_s_Triangle

func generate(numRows int) [][]int {
	var wrapperList [][]int

	// set base case
	if 1 <= numRows {
		wrapperList = append(wrapperList, []int{1})
	}

	if 2 <= numRows {
		wrapperList = append(wrapperList, []int{1, 1})
	}

	if numRows <= 2 && 0 <= numRows {
		return wrapperList
	}

	for i := 2; i < numRows; i++ {
		// create new slice
		var list []int
		for j := 0; j <= i; j++ {

			// base case if the row start/end must be 1
			if j == 0 || j == i {
				list = append(list, 1)
				continue
			}
			value := wrapperList[i-1][j-1] + wrapperList[i-1][j]
			list = append(list, value)
		}
		wrapperList = append(wrapperList, list)
	}
	return wrapperList
}
