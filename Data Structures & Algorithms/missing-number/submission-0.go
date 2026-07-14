func missingNumber(nums []int) int {
	sum1 := 0
	for i := range len(nums) + 1 {
		sum1 += i
	}

	sum2 := 0
	for _, v := range nums {
		sum2 += v
	}
	return sum1 - sum2
}
