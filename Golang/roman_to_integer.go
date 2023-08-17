package leetcode

import (
	"math"
	"strings"
)

type RomanNumber int

const (
	I RomanNumber = 1
	V RomanNumber = 5
	X RomanNumber = 10
	L RomanNumber = 50
	C RomanNumber = 100
	D RomanNumber = 500
	M RomanNumber = 1000
)

var RomanNumberMappings = map[string]RomanNumber{
	"I": I,
	"V": V,
	"X": X,
	"L": L,
	"C": C,
	"D": D,
	"M": M,
}

func romanToInt(s string) int {
	chars := strings.Split(s, "")
	total := 0

	for x, el := range chars {
		var currNum, next = int(RomanNumberMappings[el]), math.MaxInt
		if x + 1 < len(s) {
			next = int(RomanNumberMappings[chars[x+1]])
		}else if x == len(s) - 1 {
			next = 0
		}
		if currNum < next {
			total -= currNum
		} else {
			total += currNum
		}
	}

	return total
}