package algo

func maxArea(height []int) int {
	area := 0
	left, right := 0, len(height)-1

	for left <= right {
		tmp := 0
		if height[left] > height[right] {
			tmp = (right - left) * height[right]
			right--
		} else {
			tmp = (right - left) * height[left]
			left++
		}

		if tmp > area {
			area = tmp
		}
	}

	return area
}
