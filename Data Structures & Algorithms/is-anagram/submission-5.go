func isAnagram(s string, t string) bool {
	seen := make(map[rune]int)
	for _, c := range s {
		seen[c]++
	}
	
	for _, c := range t {
		seen[c]--
		if seen[c] < 0 {
			return false
		}
		if seen[c] == 0 {
			delete(seen, c)
		}
	}
	return len(seen) == 0
}
