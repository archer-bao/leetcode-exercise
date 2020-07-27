package algo

import (
	"testing"
)

func TestFunc(t *testing.T) {
	if longestCommonPrefix([]string{"1234", "123", "1234"}) != "123" {
		t.Fatal()
	}

}
