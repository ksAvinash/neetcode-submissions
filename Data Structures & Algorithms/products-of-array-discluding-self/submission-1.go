func productExceptSelf(nums []int) []int {
	prod, zeros := 1, 0
	for _, n := range nums {
		if n == 0 {
			zeros++
		} else {
			prod *= n
		}
	}
	// fmt.Println(prod, zeros)

	res := make([]int, len(nums))
	if zeros > 1 {
		return res
	}

	for i, n := range nums {
		if zeros == 0 {
			res[i] = prod / n
		} else {
			if n == 0 {
				res[i] = prod
			} else {
				res[i] = 0
			}
		}
	}
	return res
}
