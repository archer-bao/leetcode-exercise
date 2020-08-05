package algo

import (
	"testing"
)

func TestFunc(t *testing.T) {
	if len(fourSum([]int{1, 0, -1, 0, -2, 2}, 0)) != 3 {
		t.Fatal()
	}
}
