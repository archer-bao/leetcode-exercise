package algo

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	s1 := "abcdabcd"
	s2 := "abcabc"
	s3 := "aaaaaa"
	s4 := "abcddcba"
	s5 := "abcdefg"
	s6 := ""
	s7 := " "
	s8 := "  "
	s9 := "dvdf"

	if lengthOfLongestSubstring(s1) != 4 {
		t.Fatal()
	}
	if lengthOfLongestSubstring(s2) != 3 {
		t.Fatal()
	}
	if lengthOfLongestSubstring(s3) != 1 {
		t.Fatal()
	}
	if lengthOfLongestSubstring(s4) != 4 {
		t.Fatal()
	}
	if lengthOfLongestSubstring(s5) != 7 {
		t.Fatal()
	}
	if lengthOfLongestSubstring(s6) != 0 {
		t.Fatal()
	}
	if lengthOfLongestSubstring(s7) != 1 {
		t.Fatal()
	}
	if lengthOfLongestSubstring(s8) != 1 {
		t.Fatal()
	}
	if lengthOfLongestSubstring(s9) != 3 {
		t.Fatal()
	}
}
