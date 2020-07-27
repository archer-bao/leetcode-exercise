package algo

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	reverseNum := 0
	for x > reverseNum {
		reverseNum = reverseNum*10 + x%10
		x /= 10
	}

	if x == reverseNum || x == reverseNum/10 {
		return true
	}

	return false
}
