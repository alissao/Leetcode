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
	var expected = -1
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

func TestCaseFirstOcurrence_3(t *testing.T) {
	var needle string = "aaaa"
	var haystack string = "aaa"
	var result = fOcurr.StrStr(haystack, needle)
	var expected = -1
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

func TestCaseFirstOcurrence_4(t *testing.T) {
	var needle string = "issip"
	var haystack string = "mississippi"
	var result = fOcurr.StrStr(haystack, needle)
	var expected = 4
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

func TestCaseFirstOcurrence_5(t *testing.T) {
	var needle string = "issipi"
	var haystack string = "mississippi"
	var result = fOcurr.StrStr(haystack, needle)
	var expected = -1
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

func TestCaseFirstOcurrence_6(t *testing.T) {
	var needle string = "pi"
	var haystack string = "mississippi"
	var result = fOcurr.StrStr(haystack, needle)
	var expected = 9
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