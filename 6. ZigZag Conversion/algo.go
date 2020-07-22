func convert(s string, numRows int) string {
	var size = len(s)

	if size < 2 || numRows == 1 {
		return s
	}

	var (
		res  = ""
		step = 2*numRows - 2
	)

	for i := 0; i < numRows; i++ {
		for j := 0; j+i < size; j += step {
			res += s[j+i : j+i+1]
			if i != 0 && i != numRows-1 && j+step-i < size {
				res += s[j+step-i : j+step-i+1]
			}
		}
	}

	return res
}
