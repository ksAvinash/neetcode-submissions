func findDuplicate(nums []int) int {
    counts := map[int]int{}
	for _, v := range nums {
		counts[v]++
		if counts[v] > 1 {
			return v
		}
	}
	return 0
}
