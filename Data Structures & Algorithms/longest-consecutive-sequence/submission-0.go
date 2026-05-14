func longestConsecutive(nums []int) int {
	pos := map[int]int{}
	for i, v := range nums {
		pos[v] = i
	}
	fmt.Println(pos)

	i, max := 0, 0
	for i < len(nums) {
		n, c := nums[i], 0
		for {
			_, ex := pos[n]
			if ex {
				n++
				c++
			} else {
				break
			}
		}
		if c > max {
			max = c
		}
		i++
	}
	return max
}
