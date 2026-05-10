func search(nums []int, target int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		mid := int((i+j)/2)
		
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	return -1
}
