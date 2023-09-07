package firstOcurrence_test

import (
	"testing"
	fOcurr "leetcode.alisson/Golang/find_the_index_first_ocurrence_string"
)

func TestCaseFirstOcurrence_1(t *testing.T) {
	var needle string = "sad"
	var haystack string = "sadbutsad"
	var result = fOcurr.StrStr(haystack, needle)
	var expected = 0
	if result != expected {
		t.Fatalf(
			`StrStr(%s, %s) should have returned %d, but the result was %d.`, 
			haystack, 
			needle, 
			expected, 
			result,
		)
	}
}

func TestCaseFirstOcurrence_2(t *testing.T) {
	var needle string = "leeto"
	var haystack string = "leetcode"
	var result = fOcurr.StrStr(haystack, needle)
	var expected = 0
	if result != expected {
		t.Fatalf(
			`StrStr(%s, %s) should have returned %d, but the result was %d.`, 
			haystack, 
			needle, 
			expected, 
			result,
		)
	}
}