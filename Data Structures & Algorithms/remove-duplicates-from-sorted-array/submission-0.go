func removeDuplicates(nums []int) int {
	counts := map[int]int{}
	for _, n := range nums {
		nums[len(counts)] = n
		counts[n] = 1
	}
	return len(counts)
}
