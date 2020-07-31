package algo

import (
	"math"
	"sort"
)

func F(nums []int, target int) int {
	return threeSumClosest(nums, target)
}

func threeSumClosest(nums []int, target int) int {
	arrayLength := len(nums)
	if arrayLength < 3 {
		return 1<<31 - 1
	}

	sort.Ints(nums)

	var best = 1<<31 - 1
	var update = func(num int) {
		if math.Abs(float64(num-target)) < math.Abs(float64(best-target)) {
			best = num
		}
	}

	for x := 0; x < arrayLength; x++ {
		if x > 0 && nums[x] == nums[x-1] {
			continue
		}

		y := x + 1
		z := arrayLength - 1
		for y < z {
			sum := nums[x] + nums[y] + nums[z]
			if sum == target {
				return target
			}

			update(sum)

			if sum > target {
				z--
				for y < z && nums[z] == nums[z+1] {
					z--
				}
			} else {
				y++
				for y < z && nums[y] == nums[y-1] {
					y++
				}
			}
		}
	}

	return best
}
