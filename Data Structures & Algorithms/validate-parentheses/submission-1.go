func isValid(s string) bool {
    pairs := map[rune]rune{'}': '{', ']': '[', ')': '('}
	stack := []rune{}

	for _, v := range s {
		switch v {
			case '(', '[', '{':
				stack = append(stack, v)
			case ']', '}', ')':
				if len(stack) == 0 || stack[len(stack)-1] != pairs[v] {
					return false
				}
				stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
