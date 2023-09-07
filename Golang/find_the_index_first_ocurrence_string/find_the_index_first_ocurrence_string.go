package firstOcurrence

func StrStr(haystack string, needle string) int {
	if len(needle) > len(haystack) {
		return -1
	}

	nIndex := 0
	firstAppear := -1

	for i := 0; i < len(haystack) && nIndex < len(needle); i++ {
		if haystack[i] == needle[nIndex] {
			if firstAppear == -1 {
				firstAppear = i
			}
			nIndex++
		} else {
			if firstAppear != -1 {
				i = firstAppear
				firstAppear = -1
			}
			nIndex = 0
		}
	}

	if (len(haystack) - firstAppear < len(needle)) {
		firstAppear = -1
	}

	return firstAppear
}
