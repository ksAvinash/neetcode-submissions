func getVal(v string) int {
	switch v {
		case "I": return 1
		case "V": return 5
		case "X": return 10
		case "L": return 50
		case "C": return 100
		case "D": return 500
		case "M": return 1000
		default:
			return 0
	}
}

func romanToInt(s string) int {
	count, i := 0, 0

	for i < len(s) {
		v, n := getVal(string(s[i])), 0
		if i < len(s)-1 {
			n = getVal(string(s[i+1]))
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
