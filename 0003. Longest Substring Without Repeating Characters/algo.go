package algo

func lengthOfLongestSubstring(s string) int {
	if len(s) < 1 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}

	var (
		charMap               = make(map[int32]int)
		maxLength             = 0
		uniqueCharStringStart = 0
	)

	for index, each := range s {
		// 判断字符是否重复
		if at, ok := charMap[each]; ok {
			// 重复则将不重复字符串首字符位置指示符移动到重复字符后
			if at >= uniqueCharStringStart {
				uniqueCharStringStart = at + 1
			}
		}

		if index-uniqueCharStringStart+1 > maxLength {
			maxLength = index - uniqueCharStringStart + 1
		}

		charMap[each] = index
	}

	return maxLength
}
