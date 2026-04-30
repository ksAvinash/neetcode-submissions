func twoSum(nums []int, target int) []int {
    pos := map[int]int{}

	for i, v := range nums {
		rem := target - v
		r1, ex := pos[rem]
		if ex {
			return []int{r1, i}
		} else {
			pos[v] = i
		}
	}
	return []int{}
}
