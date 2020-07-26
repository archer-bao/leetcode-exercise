func reverse(x int) int {
	if x > -10 && x < 10 {
		return x
	}

	var y int

	sign := 1
	if x < 0 {
		sign = -1
		x = -x
	}

	for x != 0 {
		// 需要判断是否溢出
		if y > (1<<31-1)/10 {
			return 0
		}
		// 不需要判断什么最后一位是7还是8，因为int32区间两端最大最小值是以2开头的
		// 能存下可能导致溢出的数只可能以1或者2开头

		y = y*10 + x%10
		x /= 10
	}

	return y * sign
}
