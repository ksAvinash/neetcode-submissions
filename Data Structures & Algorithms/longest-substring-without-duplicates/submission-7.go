func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	max := 1
	for i := 0; i < len(s); i++ {
		j := i
		counts := map[string]int{}
		for j < len(s) {
			counts[string(s[j])]++
			if counts[string(s[j])] == 2 {
				break
			}
			j++
		}
		diff := j-i
		// fmt.Println("max from", string(s[i]), " = ", diff, "with counts", counts)
		if diff > max {
			max = diff
		}
	}
	return max
}
