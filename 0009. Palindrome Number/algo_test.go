package algo

import (
	"testing"
)

func TestReverse(t *testing.T) {
	if isPalindrome(12345) {
		t.Fatal()
	}

	if isPalindrome(-12345) {
		t.Fatal()
	}

	if isPalindrome(54321) {
		t.Fatal()
	}

	if !isPalindrome(787) {
		t.Fatal()
	}

	if isPalindrome(7890) {
		t.Fatal()
	}

	if !isPalindrome(12344321) {
		t.Fatal()
	}

	if isPalindrome(5656) {
		t.Fatal()
	}

	if !isPalindrome(56565) {
		t.Fatal()
	}
}
