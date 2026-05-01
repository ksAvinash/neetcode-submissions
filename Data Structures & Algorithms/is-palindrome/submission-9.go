func pickAlphaNumeric(s string) string {
	re := regexp.MustCompile(`[a-zA-Z0-9]`)
	matches := re.FindAllString(s, -1)
	s = strings.ToLower(strings.Join(matches, ""))
	return s
}

func isPalindrome(s string) bool {
	s = pickAlphaNumeric(s)
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
