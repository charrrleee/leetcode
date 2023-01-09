package _05__Isomorphic_Strings

func isIsomorphic(s string, t string) bool {
	mapS := make(map[string]string)
	mapT := make(map[string]string)

	// assume s and t are the same length
	for i := 0; i < len(s); i++ {
		charS := string(s[i])
		charT := string(t[i])

		mapCharS, okS := mapS[charS]
		mapCharT, okT := mapT[charT]

		if okS && mapCharS != charT {
			return false
		} else if !okS {
			mapS[charS] = charT
		}

		if okT && mapCharT != charS {
			return false
		} else if !okS {
			mapT[charT] = charS
		}
	}
	return true
}
