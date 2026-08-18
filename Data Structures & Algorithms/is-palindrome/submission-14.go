func isAlnum(k byte) bool {
	return unicode.IsDigit(rune(k)) || unicode.IsLetter(rune(k))
}

func toLower(k byte) rune {
	return unicode.ToLower(rune(k))
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
		if toLower(s[i]) != toLower(s[j]) {
			return false
		}
		i++
		j--
	}
	return true
}
