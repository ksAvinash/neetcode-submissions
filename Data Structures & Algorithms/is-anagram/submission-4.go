func isAnagram(s string, t string) bool {
	if len(s) == 0 || len(s) != len(t) {
		return false
	}

	counts := map[rune]int{}
	for _, v := range s {
		counts[v]++
	}
	for _, v := range t {
		counts[v]--
		if counts[v] == 0 {
			delete(counts, v)
		}
	}
	return len(counts) == 0
}
