func productExceptSelf(nums []int) []int {
	n := len(nums)
	pref, suff := make([]int, len(nums)), make([]int, len(nums))
	pref[0], suff[len(nums)-1] = 1, 1 // nothing left of pref & nothing right of suff

	for i := 1; i < n; i++ {
		pref[i] = nums[i-1] * pref[i-1]
	}
	for i := n - 2; i >=0; i-- {
		suff[i] = nums[i+1] * suff[i+1]
	}

	res := make([]int, len(nums))
	for i := 0; i < n; i++ {
		res[i] = pref[i] * suff[i]
	}
	return res
}
