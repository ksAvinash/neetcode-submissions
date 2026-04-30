func longestCommonPrefix(strs []string) string {
	res := ""
	for i, ch := range strs[0] { // take 1st string to pick chars
		for _, s := range strs {
			if i >= len(s) || rune(s[i]) != ch {
				return res
			}
		}
		res += string(ch)
	}
	return res
}
