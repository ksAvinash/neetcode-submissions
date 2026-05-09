func isPalindrome(sub string, i int, j int) bool {
	for i < j {
		if sub[i] != sub[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func validPalindrome(s string) bool {
	i, j := 0, len(s)-1

	for i < j {
		if s[i] != s[j] {
			return isPalindrome(s, i, j-1) || isPalindrome(s, i+1, j)
		}
		i++
		j--
	}
	return true
}
