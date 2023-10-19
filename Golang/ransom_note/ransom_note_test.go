package ransomnote_test

import (
	"testing"

	ransomNote "leetcode.alisson/Golang/ransom_note"
)

func TestCaseRansomNote_1(t *testing.T) {
	var ransomNoteString string = "a"
	var magazine string = "b"
	var result = ransomNote.CanConstruct(ransomNoteString, magazine)
	var expected = false
	if result != expected {
		t.Fatalf(
			`IsSubsequence(%s, %s) should have returned %t, but the result was %t.`,
			magazine,
			ransomNoteString,
			expected,
			result,
		)
	}
}
func TestCaseRansomNote_2(t *testing.T) {
	var ransomNoteString string = "aa"
	var magazine string = "ab"
	var result = ransomNote.CanConstruct(ransomNoteString, magazine)
	var expected = false
	if result != expected {
		t.Fatalf(
			`IsSubsequence(%s, %s) should have returned %t, but the result was %t.`,
			magazine,
			ransomNoteString,
			expected,
			result,
		)
	}
}
func TestCaseRansomNote_3(t *testing.T) {
	var ransomNoteString string = "aa"
	var magazine string = "aab"
	var result = ransomNote.CanConstruct(ransomNoteString, magazine)
	var expected = true
	if result != expected {
		t.Fatalf(
			`IsSubsequence(%s, %s) should have returned %t, but the result was %t.`,
			magazine,
			ransomNoteString,
			expected,
			result,
		)
	}
}

func TestCaseRansomNote_4(t *testing.T) {
	var ransomNoteString string = "bg"
	var magazine string = "efjbdfbdgfjhhaiigfhbaejahgfbbgbjagbddfgdiaigdadhcfcj"
	var result = ransomNote.CanConstruct(ransomNoteString, magazine)
	var expected = true
	if result != expected {
		t.Fatalf(
			`IsSubsequence(%s, %s) should have returned %t, but the result was %t.`,
			magazine,
			ransomNoteString,
			expected,
			result,
		)
	}
}

