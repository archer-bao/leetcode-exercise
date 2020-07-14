package algo

func twoSum(nums []int, target int) []int {
	if len(nums) < 2 {	//数组长度<2时无意义
		return nil
	}
	
	elemMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if j, ok := elemMap[target-nums[i]]; ok {
			return []int{j, i}	//index顺序为先j后i，因为j是之前搜索时放进去的，因此j一定在i前
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
