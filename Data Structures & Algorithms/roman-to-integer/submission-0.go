func getVal(curr string, next string) (int, bool) {
	if curr == "I" {
		if next == "V" {
			return 4, true
		}
		if next == "X" {
			return 9, true
		}
		return 1, false
	}
	if curr == "X" {
		if next == "L" {
			return 40, true
		}
		if next == "C" {
			return 90, true
		}
		return 10, false
	}
	if curr == "C" {
		if next == "D" {
			return 400, true
		}
		if next == "M" {
			return 900, true
		}
		return 100, false
	}
	if curr == "V" {
		return 5, false
	}
	if curr == "L" {
		return 50, false
	}
	if curr == "D" {
		return 500, false
	}
	if curr == "M" {
		return 1000, false
	}
	return 0, false
}

func romanToInt(s string) int {
	count, i := 0, 0

	for i < len(s) {
		c := string(s[i])
		n := ""
		if i < len(s)-1 {
			n = string(s[i+1])
		}
		val, skip := getVal(c, n)
		count += val
		if skip {
			i++
		}
		i++
	}
	return count
}
