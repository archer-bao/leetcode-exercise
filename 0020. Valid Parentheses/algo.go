package algo

func isValid(s string) bool {
	bracket := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}'}

	var stack = make([]byte, 0, len(s)/2)
	for index := range s {
		if _, ok := bracket[s[index]]; ok {
			stack = append(stack, s[index])
		} else {
			l := len(stack)
			if l == 0 || bracket[stack[l-1]] != s[index] {
				return false
			}
			stack = stack[:l-1]
		}
	}

	return len(stack) == 0
}
