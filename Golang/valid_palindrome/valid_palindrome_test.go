package validpalindrome_test

import (
	"testing"

	palindrome "leetcode.alisson/Golang/valid_palindrome"
)

func TestCaseIsPalindrome_1(t *testing.T) {
	var somePhrase string = "A man, a plan, a canal: Panama"
	var result = palindrome.IsPalindrome(somePhrase)
	var expected = true
	if result != expected {
		t.Fatalf(
			`IsPalindrome(%s) should have returned %t, but the result was %t.`,
			somePhrase,
			expected,
			result,
		)
	}
}

func TestCaseIsPalindrome_2(t *testing.T) {
	var somePhrase string = "race a car"
	var result = palindrome.IsPalindrome(somePhrase)
	var expected = false
	if result != expected {
		t.Fatalf(
			`IsPalindrome(%s) should have returned %t, but the result was %t.`,
			somePhrase,
			expected,
			result,
		)
	}
}

func TestCaseIsPalindrome_3(t *testing.T) {
	var somePhrase string = " "
	var result = palindrome.IsPalindrome(somePhrase)
	var expected = true
	if result != expected {
		t.Fatalf(
			`IsPalindrome(%s) should have returned %t, but the result was %t.`,
			somePhrase,
			expected,
			result,
		)
	}
}
