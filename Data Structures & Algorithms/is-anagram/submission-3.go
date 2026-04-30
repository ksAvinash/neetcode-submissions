func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	counts := map[rune]int{}
	for _, v := range s {
		counts[v]++
	}

	for _, v := range t {
		if counts[v] == 0 {
			return false
		}
		counts[v]--
	}
	return true
}
