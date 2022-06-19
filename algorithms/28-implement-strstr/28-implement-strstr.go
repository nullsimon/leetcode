package implementstrstr

func strStr(haystack string, needle string) int {

	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[0] {
			if len(haystack)-i < len(needle) {
				return -1
			}
			not := false
			for j := 0; j < len(needle); j++ {

				if haystack[i+j] != needle[j] {
					not = true
					break
				}

			}
			if not {
				continue
			} else {
				return i

			}
		}
	}
	return -1
}
