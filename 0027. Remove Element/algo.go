package algo

func removeElement(nums []int, val int) int {
	length := len(nums)
	if length < 1 {
		return -1
	}

	i := 0
	for j := 0; j < length; j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}

	return i
}
