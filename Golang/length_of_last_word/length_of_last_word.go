package lengthOfLastWord

import "strings"

func LengthOfLastWord(s string) int {
  var trimmedS = strings.Trim(s, " ")
	var splittedTrimmedS = strings.Split(trimmedS, " ")
	return len(splittedTrimmedS[len(splittedTrimmedS)-1])
}