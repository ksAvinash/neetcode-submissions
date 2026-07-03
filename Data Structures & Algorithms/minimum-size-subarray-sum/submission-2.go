func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	res := n+1

	l, total := 0, 0
	for r := 0; r < n; r++ {
		total += nums[r]
		for total >= target {
			res = min(res, r-l+1)
			total -= nums[l]
			l++
		}
	}

	if res > n {
		return 0
	}
	return res
}
