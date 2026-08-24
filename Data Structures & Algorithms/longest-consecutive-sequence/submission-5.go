func longestConsecutive(nums []int) int {
	lookup := make(map[int]struct{}, len(nums))
	for _, n := range nums {
		lookup[n] = struct{}{}
	}
	max := 0
	for _, n := range nums {
		if _, ex := lookup[n-1]; ex {
			continue // because; n is not beg of sequence
		}

		streak, curr := 0, n
		for {
			if _, ex := lookup[curr]; !ex {
				break
			}
			curr++
			streak++
		}
		if streak > max {
			max = streak
		}
	}
	return max
}
