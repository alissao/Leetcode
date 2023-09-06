package longestcommonprefix_test

import (
	"testing"
	longestCommon "leetcode.alisson/Golang/longest_common_prefix"
)

func TestCaseLongestCommonPrefix_1(t *testing.T) {
	var stringArray []string = []string{"flower","flow","flight"}
	var result = longestCommon.LongestCommonPrefix(stringArray)
	var expected = "fl"
	if result != expected {
		t.Fatalf(
			`LongestCommonPrefix(%s) should have returned %s, but the result was %s.`, 
			stringArray, 
			expected, 
			result,
		)
	}
}

func TestCaseLongestCommonPrefix_2(t *testing.T) {
	var stringArray []string = []string{"dog","racecar","car"}
	var result = longestCommon.LongestCommonPrefix(stringArray)
	var expected = ""
	if result != expected {
		t.Fatalf(
			`LongestCommonPrefix(%s) should have returned %s, but the result was %s.`, 
			stringArray, 
			expected, 
			result,
		)
	}
}