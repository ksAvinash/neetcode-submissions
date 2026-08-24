func longestConsecutive(nums []int) int {
	lookup := map[int]int{}
	for _, n := range nums {
		lookup[n] = 1
	}

	max := 0
	for _, n := range nums {
		streak, curr := 0, n
		for {
			v, ex := lookup[curr]
			if !ex {
				break
			}
			if v > 1 {
				streak += v
				break
			}
			curr++
			streak++
		}
		lookup[n] = streak
		if streak > max {
			max = streak
		}
	}
	

	return max
}
