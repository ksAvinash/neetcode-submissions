func isAlnum(k byte) bool {
	return unicode.IsDigit(rune(k)) || unicode.IsLetter(rune(k))
}

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i <= j {
		if !isAlnum(s[i]) {
			i++
			continue
		}
		if !isAlnum(s[j]) {
			j--
			continue
		}
		if unicode.ToLower(rune(s[i])) != unicode.ToLower(rune(s[j])) {
			return false
		}
		i++
		j--
	}
	return true
}
