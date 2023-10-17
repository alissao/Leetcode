package twosumii_test

import (
	"testing"
	"reflect"

	twosumii "leetcode.alisson/Golang/two_sum_ii"
)

func TestCaseTwoSumII_1(t *testing.T) {
	var numbers []int = []int{2,7,11,15}
	var target int = 9
	var result = twosumii.TwoSum(numbers, target)
	expected := []int{1,2}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf(
			`TwoSum(%d, %d) should have returned %d, but the result was %d.`,
			numbers,
			target,
			expected,
			result,
		)
	}
}


func TestCaseTwoSumII_2(t *testing.T) {
	var numbers []int = []int{2,3,4}
	var target int = 6
	var result = twosumii.TwoSum(numbers, target)
	expected := []int{1,3}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf(
			`TwoSum(%d, %d) should have returned %d, but the result was %d.`,
			numbers,
			target,
			expected,
			result,
		)
	}
}


func TestCaseTwoSumII_3(t *testing.T) {
	var numbers []int = []int{-1,0}
	var target int = -1
	var result = twosumii.TwoSum(numbers, target)
	expected := []int{1,2}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf(
			`TwoSum(%d, %d) should have returned %d, but the result was %d.`,
			numbers,
			target,
			expected,
			result,
		)
	}
}

