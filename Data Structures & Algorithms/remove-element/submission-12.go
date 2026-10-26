func removeElement(nums []int, val int) int {
    count := 0
	for _, n := range nums {
		if n == val {
			continue
		} else {
			nums[count] = n
			count++
		}
	}
	return count
}
