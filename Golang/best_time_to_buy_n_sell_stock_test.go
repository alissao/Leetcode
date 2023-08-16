package leetcode

import "testing"

func TestCaseBestTimeStock1(t *testing.T) {
	var prices []int = []int{7,1,5,3,6,4}
	var result = maxProfit(prices)
	var expected = 5 
	if result != expected {
		t.Fatalf(`maxProfit(%d) should have returned %d, but the result was %d.`, prices, expected, result)
	}
}

func TestCaseBestTimeStock2(t *testing.T) {
	var prices []int = []int{7,6,4,3,1}
	var result = maxProfit(prices)
	var expected = 0
	if result != expected {
		t.Fatalf(`maxProfit(%d) should have returned %d, but the result was %d.`, prices, expected, result)
	}
}

func TestCaseBestTimeStock3(t *testing.T) {
	var prices []int = []int{2,4,1}
	var result = maxProfit(prices)
	var expected = 2 
	if result != expected {
		t.Fatalf(`maxProfit(%d) should have returned %d, but the result was %d.`, prices, expected, result)
	}
}

func TestCaseBestTimeStock4(t *testing.T) {
	var prices []int = []int{2,1,2,1,0,1,2}
	var result = maxProfit(prices)
	var expected = 2
	if result != expected {
		t.Fatalf(`maxProfit(%d) should have returned %d, but the result was %d.`, prices, expected, result)
	}
}

func TestCaseBestTimeStock5(t *testing.T) {
	var prices []int = []int{4,11,2,7,1}
	var result = maxProfit(prices)
	var expected = 7
	if result != expected {
		t.Fatalf(`maxProfit(%d) should have returned %d, but the result was %d.`, prices, expected, result)
	}
}