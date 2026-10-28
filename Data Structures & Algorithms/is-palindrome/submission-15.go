func isAlnum(ch byte) bool {
	r := rune(ch)
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}

func toLower(ch byte) rune {
	r := rune(ch)
	return unicode.ToLower(r)
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
