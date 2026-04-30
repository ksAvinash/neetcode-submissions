func hasDuplicate(nums []int) bool {
    counts := map[int]int{}
	
	for _, v := range nums {
		_, ex := counts[v]
		if ex {
			return true
		} else {
			counts[v] = 1
		}
	}
	return false
}
