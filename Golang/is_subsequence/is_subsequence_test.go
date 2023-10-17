package issubsequence_test

import (
	"testing"

	palindrome "leetcode.alisson/Golang/is_subsequence"
)

func TestCaseIsSubsequence_1(t *testing.T) {
	var firstString string = "abc"
	var secondString string = "ahbgdc"
	var result = palindrome.IsSubsequence(firstString, secondString)
	var expected = true
	if result != expected {
		t.Fatalf(
			`IsSubsequence(%s, %s) should have returned %t, but the result was %t.`,
			firstString,
			secondString,
			expected,
			result,
		)
	}
}

func TestCaseIsSubsequence_2(t *testing.T) {
	var firstString string = "axc"
	var secondString string = "ahbgdc"
	var result = palindrome.IsSubsequence(firstString, secondString)
	var expected = false
	if result != expected {
		t.Fatalf(
			`IsSubsequence(%s, %s) should have returned %t, but the result was %t.`,
			firstString,
			secondString,
			expected,
			result,
		)
	}
}

func TestCaseIsSubsequence_3(t *testing.T) {
	var firstString string = ""
	var secondString string = "ahbgdc"
	var result = palindrome.IsSubsequence(firstString, secondString)
	var expected = true
	if result != expected {
		t.Fatalf(
			`IsSubsequence(%s, %s) should have returned %t, but the result was %t.`,
			firstString,
			secondString,
			expected,
			result,
		)
	}
}