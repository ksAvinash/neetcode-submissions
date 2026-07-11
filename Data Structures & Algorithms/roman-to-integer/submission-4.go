func romanToInt(s string) int {
	count, i := 0, 0
	values := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

	for i < len(s) {
		v, n := values[s[i]], 0
		if i < len(s)-1 {
			n = values[s[i+1]]
		}
		if v < n {
			count += n - v
			i++
		} else {
			count += v
		}
		i++
	}
	return count
}
