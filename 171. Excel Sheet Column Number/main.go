package _71__Excel_Sheet_Column_Number

// start from the first
func titleToNumber(columnTitle string) int {
	sum := 0

	for i := 0; i < len(columnTitle); i++ {
		// "A" >> getValue("A")
		// "AB" >> 26 * getValue("A")  + getValue("B")
		// "ZY" >> 26 * getValue("Z") + getValue("Y")
		// AAY >> 26 * 26 * getValue("A") + 26 * getValue("A") + getValue("Y")

		column := columnTitle[i]

		//  ASCII A = 65
		tempSum := int(column) - 64

		// if columnTitle is in 10 digit
		for j := len(columnTitle) - 1; j > i; j-- {
			tempSum = tempSum * 26
		}
		sum = sum + tempSum
	}

	return sum
}

// start from last
func titleToNumber2(columnTitle string) int {
	//  ASCII A = 65
	sum := 0
	for _, column := range columnTitle {
		v := int(column) - 64
		sum = sum * 26
		sum = sum + v
	}
	return sum
}

