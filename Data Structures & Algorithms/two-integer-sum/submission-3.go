func twoSum(nums []int, target int) []int {
    seen := make(map[int]int)
	for i, num := range nums {
		rem := target - num
		val, ex := seen[rem]
		if ex {
			return []int{val, i}
		} else {
			seen[num] = i
		}
	}
	return []int{}
}
