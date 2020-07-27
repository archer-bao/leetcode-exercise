package algo

import (
	"testing"
)

func TestFunc(t *testing.T) {
	if len(threeSum([]int{-1, 0, 1, 2, -1, -4})) != 2 {
		t.Fatal()
	}

}
