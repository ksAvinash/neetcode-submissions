func lengthOfLongestSubstring(s string) int {
	pos := map[byte]int{}

	l, res := 0, 0
	for r := 0; r < len(s); r++ {
		if i, ex := pos[s[r]]; ex {
			l = max(i+1, l)
		}
		pos[s[r]] = r
		window := r - l + 1
		if window > res {
			res = window
		}
	}
	return res
}
