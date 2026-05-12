type Stack struct {
	items []string
}
func (s *Stack) empty() bool {
	return len(s.items) == 0
}
func (s *Stack) push(ch string) {
	s.items = append(s.items, ch)
}
func (s *Stack) top() string {
	if s.empty() {
		return ""
	}
	return s.items[len(s.items)-1]
}
func (s *Stack) pop() string {
	ch := s.top()
	s.items = s.items[:len(s.items)-1]
	return ch
}



func isValid(s string) bool {
	stk := Stack{}

    for _, v := range s {
		ch := string(v)
		if ch == "[" || ch == "{" || ch == "(" {
			stk.push(ch)
		} else if ch == "]" {
			if stk.top() != "[" {
				return false
			}
			stk.pop()
		} else if ch == "}" {
			if stk.top() != "{" {
				return false
			}
			stk.pop()
		} else if ch == ")" {
			if stk.top() != "(" {
				return false
			}
			stk.pop()
		}
	}
	return stk.empty()
}
