func plusOne(digits []int) []int {
	// if digits[len(digits)-1] < 9 {
	// 	digits[len(digits)-1] += 1
	// 	return digits
	// }


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
