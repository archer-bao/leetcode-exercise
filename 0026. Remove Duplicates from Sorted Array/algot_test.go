package algo

import (
	"testing"
)

func TestFunc(t *testing.T) {
	nums := []int{1, 2, 2, 3, 4, 5, 5}
	if removeDuplicates(nums) != 5 {
		t.Fatal()
	}
}
