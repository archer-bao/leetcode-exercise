package algo

import (
	"testing"
)

func TestFunc(t *testing.T) {
	if threeSumClosest([]int{-1, 0, 1, 1, 55}, 3) != 2 {
		t.Fatal()
	}

}
