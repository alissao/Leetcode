package validpalindrome

import (
	"regexp"
	"strings"
)

func IsPalindrome(s string) bool {
	
	treatedString := strings.ToLower(
		strings.ReplaceAll(
			regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(s, ""), 
			" ", 
			"",
		),
	)
	
	if (len(treatedString) == 0) { return true }

	pointer1 := 0
	pointer2 := len(treatedString) -1
	isPalindrome := true

	for pointer1 < pointer2 {
		if (treatedString[pointer1] != treatedString[pointer2]) {
			isPalindrome = false
			break
		}
		pointer1++
		pointer2--
	}
  
	return isPalindrome
}