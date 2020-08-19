package algo

import (
	"testing"
)

func TestFunc(t *testing.T) {
	nums := []int{1, 2, 2, 3, 3, 3, 4, 5}
	if removeElement(nums, 3) != 5 {
		t.Fatal()
	}
}
