
type Stack struct {
	items []int
}
func(s *Stack) empty() bool {
	return len(s.items) == 0
}
func(s *Stack) top() int {
	if s.empty() {
		return 0
	}
	return s.items[len(s.items)-1]
}
func(s *Stack) push(ch int) {
	s.items = append(s.items, ch)
}
func(s *Stack) pop() int {
	if s.empty() {
		return 0
	}
	ch := s.top()
	s.items = s.items[:len(s.items)-1]
	return ch
}


func calPoints(operations []string) int {
	stk := &Stack{}

	for _, op := range operations {
		if op == "C" {
			stk.pop()
		} else if op == "D" {
			n := stk.pop()
			stk.push(n)
			stk.push(n * 2)
		} else if op == "+" {
			n1, n2 := stk.pop(), stk.pop()
			stk.push(n2); stk.push(n1)
			stk.push(n1 + n2)
		} else {
			n, _ := strconv.Atoi(op)
			stk.push(n)
		}
		// fmt.Println(stk.items)
	}
	sum := 0
	for _, c := range stk.items {
		sum += c
	}
	return sum
}
