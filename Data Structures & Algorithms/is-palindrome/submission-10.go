func isPalindrome(s string) bool {
	r := []rune(s)
	i, j := 0, len(r)-1
	for i < j {
		// skip non-alpha forward
		for i < j && !unicode.IsLetter(r[i]) && !unicode.IsDigit(r[i]) {
			i++
		}

		// skip non-alpha backward
		for i < j && !unicode.IsLetter(r[j]) && !unicode.IsDigit(r[j]) {
			j--
		}

		if unicode.ToLower(r[i]) != unicode.ToLower(r[j]) {
			return false
		}
		i++
		j--
	}
	return true
}
