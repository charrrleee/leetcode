package _71__Jewels_and_Stones

func numJewelsInStones(jewels string, stones string) int {
	stoneMap := make(map[string]int)
	for _, stone := range stones {
		stoneMap[string(stone)]++
	}

	counter := 0

	for _, jewel := range jewels {
		v, ok := stoneMap[string(jewel)]
		if ok {
			counter = counter + v
		}
	}
	return counter
}
