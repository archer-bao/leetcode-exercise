package algo

func twoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		return nil
	}
	elemMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if j, ok := elemMap[target-nums[i]]; ok {
			return []int{j, i}
		}
		elemMap[nums[i]] = i
	}
	return nil
}

func twoSum_BruteForce(nums []int, target int) []int {
	if len(nums) < 2 {
		return nil
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
