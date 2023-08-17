package leetcode

import (
	"testing"
)

func TestCaseRomanToInt_1(t *testing.T) {
	var romanNum string = "III"
	var result = romanToInt(romanNum)
	var expected = 3
	if result != expected {
		t.Fatalf(`romanToInt(%s) should have returned %d, but the result was %d.`, romanNum, expected, result)
	}
}

func TestCaseRomanToInt_2(t *testing.T) {
	var romanNum string = "LVIII"
	var result = romanToInt(romanNum)
	var expected = 58
	if result != expected {
		t.Fatalf(`romanToInt(%s) should have returned %d, but the result was %d.`, romanNum, expected, result)
	}
}

func TestCaseRomanToInt_3(t *testing.T) {
	var romanNum string = "MCMXCIV"
	var result = romanToInt(romanNum)
	var expected = 1994
	if result != expected {
		t.Fatalf(`romanToInt(%s) should have returned %d, but the result was %d.`, romanNum, expected, result)
	}
}