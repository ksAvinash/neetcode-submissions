func searchInsert(nums []int, target int) int {
	i, j := 0, len(nums)-1

	for i <=j {
		m := int((i+j)/2)
		if nums[m] == target {
			return m
		}
		if nums[m] < target {
			i = m + 1
		} else {
			j = m - 1
		}
	}
	return i
}
