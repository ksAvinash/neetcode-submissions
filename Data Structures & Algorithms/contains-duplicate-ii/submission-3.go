func containsNearbyDuplicate(nums []int, k int) bool {
	pos := map[int]int{}
	for i, n := range nums {
		if j, ex := pos[n]; ex {
			if i - j <= k {
				return true
			}
		}
		pos[n] = i
	}
	return false
}
