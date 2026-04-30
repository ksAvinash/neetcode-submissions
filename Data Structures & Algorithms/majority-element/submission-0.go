func majorityElement(nums []int) int {
    counts := map[int]int{}
	
	for _, n := range nums {
		counts[n]++
		if counts[n] > int(len(nums) / 2) {
			return n
		}
	}
	return 0
}
