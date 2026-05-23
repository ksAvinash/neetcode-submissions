func productExceptSelf(nums []int) []int {
	res := []int{}
	for i, _ := range nums {
		pr := 1
		for j, k := range nums {
			if i != j {
				pr *= k
			}
		}
		res = append(res, pr)
	}
	return res
}
