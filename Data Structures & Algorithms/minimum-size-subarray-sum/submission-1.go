func minSubArrayLen(target int, nums []int) int {
	window := 1
	for window <= len(nums) {
		// fmt.Println("window", window)
		i, j := 0, window-1
		for j < len(nums) {
			// fmt.Print(i, j)
			sum := 0
			k := i
			for k <= j {
				sum += nums[k]
				k++
			}
			if sum >= target {
				return window
			}
			// fmt.Print(" ", sum, "\n")
			j++
			i++
		}
		window++
		// fmt.Println(">>>---->>>>---->>>>")
	}
	return 0
}
