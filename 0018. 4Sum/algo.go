package algo

import "sort"

func fourSum(nums []int, target int) [][]int {
	arrayLength := len(nums)
	if arrayLength < 4 {
		return nil
	}

	sort.Ints(nums)
	var result [][]int

	for a := 0; a < arrayLength-3; a++ {
		if a > 0 && nums[a] == nums[a-1] {
			continue
		}

		for x := a + 1; x < arrayLength-2; x++ {
			if x > a+1 && nums[x] == nums[x-1] {
				continue
			}

			z := arrayLength - 1
			sum := target - (nums[a] + nums[x])

			for y := x + 1; y < arrayLength-1; y++ {
				if y > x+1 && nums[y] == nums[y-1] {
					continue
				}

				for y < z && nums[y]+nums[z] > sum {
					z--
				}

				if y == z {
					break
				}

				if nums[y]+nums[z] == sum {
					result = append(result, []int{nums[a], nums[x], nums[y], nums[z]})
				}
			}
		}
	}

	return result
}
