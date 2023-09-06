package lengthOfLastWord_test

import (
	"testing"
	lastWLen "leetcode.alisson/Golang/length_of_last_word"
)

func TestCaseLengthOfLastWord_1(t *testing.T) {
	var stringWithWords string = "Hello World"
	var result = lastWLen.LengthOfLastWord(stringWithWords)
	var expected = 5
	if result != expected {
		t.Fatalf(
			`LengthOfLastWord(%s) should have returned %d, but the result was %d.`, 
			stringWithWords, 
			expected, 
			result,
		)
	}
}
func TestCaseLengthOfLastWord_2(t *testing.T) {
	var stringWithWords string = "   fly me   to   the moon  "
	var result = lastWLen.LengthOfLastWord(stringWithWords)
	var expected = 4
	if result != expected {
		t.Fatalf(
			`LengthOfLastWord(%s) should have returned %d, but the result was %d.`, 
			stringWithWords, 
			expected, 
			result,
		)
	}
}
func TestCaseLengthOfLastWord_3(t *testing.T) {
	var stringWithWords string = "luffy is still joyboy"
	var result = lastWLen.LengthOfLastWord(stringWithWords)
	var expected = 6
	if result != expected {
		t.Fatalf(
			`LengthOfLastWord(%s) should have returned %d, but the result was %d.`, 
			stringWithWords, 
			expected, 
			result,
		)
	}
}