package algo

import "testing"

func TestAtoi(t *testing.T) {
	if myAtoi("words and 987") != 0 {
		t.Fatal()
	}
	if myAtoi("-12345") != -12345 {
		t.Fatal()
	}
	if myAtoi("12345") != 12345 {
		t.Fatal()
	}
	if myAtoi("012345") != 12345 {
		t.Fatal()
	}
	if myAtoi("9876543210") != 1<<31-1 {
		t.Fatal()
	}
}
