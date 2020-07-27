package algo

import "sort"

func ThreeSum(nums []int) [][]int {
	return threeSum(nums)
}

func threeSum(nums []int) [][]int {
	arrayLength := len(nums)
	if arrayLength < 3 {
		return nil
	}

	sort.Ints(nums)
	var result [][]int

	for x := 0; x < arrayLength; x++ {
		// 如果 nums[x] == nums[x-1]，那么该结果会重复
		// 因为 nums[x], nums[y], nums[z] == nums[x-1], nums[y], nums[z]
		// 此处判断 x>0，是防止nums[x-1]越界访问
		if x > 0 && nums[x] == nums[x-1] {
			continue
		}

		z := arrayLength - 1
		target := -nums[x]

		for y := x + 1; y < arrayLength; y++ {
			// 同上
			// 此处判断 y>x+1，是避免 num[y-1]==num[x] 时丢失组合
			if y > x+1 && nums[y] == nums[y-1] {
				continue
			}

			// 循环终止时，要么y、z指针重合，要么找到 ≤target 的组合
			for y < z && nums[y]+nums[z] > target {
				z--
			}

			// y、z指针重合，没希望了
			// 因为继续y++、z--和之前指针没区别
			// 只不过把y和z交换位置，根据加法交换律知道这是无用功
			if y == z {
				break
			}

			if nums[y]+nums[z] == target {
				result = append(result, []int{nums[x], nums[y], nums[z]})
			}
		}
	}

	return result
}
