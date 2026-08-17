func hasDuplicate(nums []int) bool {
    counts := map[int]int{}

	for _, v := range nums {
		counts[v]++
		if counts[v] > 1 {
			return true
		}
	}
	return false
}
