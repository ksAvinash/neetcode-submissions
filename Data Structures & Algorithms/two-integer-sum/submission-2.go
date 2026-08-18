func twoSum(nums []int, target int) []int {
    pos := map[int]int{}

	for i, v := range nums {
		rem := target - v
		val, ex := pos[rem]
		if ex {
			return []int{val, i}
		}
		pos[v] = i
	}
	return []int{}
}
