package algo

import (
	"math/rand"
	"testing"
	"time"
)

func TestTwoSum(t *testing.T) {
	rand.Seed(int64(time.Now().Nanosecond()))
	for times := 0; times < 10; times++ {
		nums := []int{}
		for i := 0; i < 100; i++ {
			nums = append(nums, rand.Intn(100))
		}
		res := twoSum(nums, 128)
		if res == nil {
			t.Error("failed")
		}
	}
}

func TestTwoSum_BruteForce(t *testing.T) {
	rand.Seed(int64(time.Now().Nanosecond()))
	for times := 0; times < 10; times++ {
		nums := []int{}
		for i := 0; i < 100; i++ {
			nums = append(nums, rand.Intn(100))
		}
		res := twoSum_BruteForce(nums, 128)
		if res == nil {
			t.Error("failed")
		}
	}
}
