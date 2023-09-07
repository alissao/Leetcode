package longestcommonprefix

import (
	"sort"
)

func LongestCommonPrefix(strs []string) string {
  if (len(strs[0]) == 0) {
		return ""
	}
	strsLen := len(strs)
	
	if (strsLen == 1) {
		return strs[0]
	}

	sort.Strings(strs)

	for i := range(strs[0]) {
		if (strs[0][i] != strs[strsLen-1][i]) {
			return strs[0][:i]
		}
	}

	return strs[0]
}