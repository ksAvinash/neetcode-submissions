func hasDuplicate(nums []int) bool {
    present := map[int]bool{}
	for _, n := range nums {
		_, ex := present[n]
		if ex {
			return true
		}
		present[n] = true
	}
	return false
}
