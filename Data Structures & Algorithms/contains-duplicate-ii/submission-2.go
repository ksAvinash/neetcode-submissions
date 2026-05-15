func containsNearbyDuplicate(nums []int, k int) bool {
	pos := map[int]int{}

	for i, n := range nums {
		_, e := pos[n]
		if e {
			if (i - pos[n]) <= k {
				return true
			}
		}
		pos[n] = i
	}

	return false
}
