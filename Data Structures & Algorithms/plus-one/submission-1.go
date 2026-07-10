func plusOne(digits []int) []int {
	ok := false
	j := len(digits) - 1
	for j >= 0 {
		if digits[j] < 9 {
			digits[j] += 1
			ok = true
			break
		} else {
			digits[j] = 0
		}
		j--
	}
	if !ok {
		digits = append([]int{1}, digits...)
	}
	return digits
}
