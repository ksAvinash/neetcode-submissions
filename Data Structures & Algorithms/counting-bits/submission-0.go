func countBits(n int) []int {
	res := []int{}
	for i := range n + 1 {
		count := 0
		for i != 0 {
			count += i & 1
			i >>= 1
		}
		res = append(res, count)
	}
	return res
}
