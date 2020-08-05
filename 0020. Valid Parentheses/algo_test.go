package algo

import "testing"

func TestAlgo(t *testing.T) {
	if !isValid("()") {
		t.Fatal()
	}
	if !isValid("()[]{}") {
		t.Fatal()
	}
	if isValid("(]") {
		t.Fatal()
	}
	if isValid("(])") {
		t.Fatal()
	}
	if isValid("([)]") {
		t.Fatal()
	}
	if !isValid("{[]}") {
		t.Fatal()
	}
}
