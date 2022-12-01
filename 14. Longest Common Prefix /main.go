package _4__Longest_Common_Prefix_

func longestCommonPrefix(strs []string) string {
	// match index key
	for i, keyStr := range strs[0] {
		for _, str := range strs {
			if i == len(str) || byte(keyStr) != str[i] {
				return strs[0][0:i]
			}
		}
	}
	return strs[0]
}
