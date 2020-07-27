package algo

import (
	"testing"
)

func TestReverse(t *testing.T) {
	if maxArea([]int{7, 1, 1, 1, 1, 1, 10, 8}) != 49 {
		t.Fatal()
	}

}
