func sortColors(nums []int) {
    counts := map[int]int{}
	for _, num := range nums {
		counts[num]++
	}

	i := 0
	for c := range []int{0, 1, 2} {
		count := counts[c]
		for count > 0 {
			nums[i] = c
			count--
			i++
		}
	}

}
