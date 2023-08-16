package leetcode

import "testing"
// import "leetcode.alisson/Golang"

func TestCaseBestTimeStock_2_1(t *testing.T) {
	var prices []int = []int{7,1,5,3,6,4}
	var result = maxProfit2(prices)
	var expected = 7
	if result != expected {
		t.Fatalf(`maxProfit2(%d) should have returned %d, but the result was %d.`, prices, expected, result)
	}
}

func TestCaseBestTimeStock_2_2(t *testing.T) {
	var prices []int = []int{1,2,3,4,5}
	var result = maxProfit2(prices)
	var expected = 4
	if result != expected {
		t.Fatalf(`maxProfit2(%d) should have returned %d, but the result was %d.`, prices, expected, result)
	}
}

func TestCaseBestTimeStock_2_3(t *testing.T) {
	var prices []int = []int{7,6,4,3,1}
	var result = maxProfit2(prices)
	var expected = 0 
	if result != expected {
		t.Fatalf(`maxProfit2(%d) should have returned %d, but the result was %d.`, prices, expected, result)
	}
}

func TestCaseBestTimeStock_2_4(t *testing.T) {
	var prices []int = []int{2,4,1}
	var result = maxProfit2(prices)
	var expected = 2
	if result != expected {
		t.Fatalf(`maxProfit2(%d) should have returned %d, but the result was %d.`, prices, expected, result)
	}
}