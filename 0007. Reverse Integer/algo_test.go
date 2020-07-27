package algo

import (
	"testing"
)

func TestReverse(t *testing.T) {
	if reverse(12345) != 54321 {
		t.Fatal()
	}

	if reverse(-12345) != -54321 {
		t.Fatal()
	}

	if reverse(54321) != 12345 {
		t.Fatal()
	}

	if reverse(789) != 987 {
		t.Fatal()
	}

	if reverse(7890) != 987 {
		t.Fatal()
	}

	if reverse(1234567899) != 0 {
		t.Fatal()
	}

	if reverse(-2147483412) != -2143847412 {
		t.Fatal()
	}
}
