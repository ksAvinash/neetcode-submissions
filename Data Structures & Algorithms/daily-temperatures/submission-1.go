type MyStk struct {
	items []int
}
func (s *MyStk) push(v int) {
	s.items = append(s.items, v)
}
func (s *MyStk) pop() int {
	if s.empty() {
		return 0
	}
	t := s.top()
	s.items = s.items[:len(s.items)-1]
	return t
}
func (s *MyStk) top() int {
	if s.empty() {
		return 0
	}
	return s.items[len(s.items)-1]
}
func (s *MyStk) empty() bool {
	return len(s.items) == 0
}


func dailyTemperatures(temperatures []int) []int {
	stk := MyStk{}
	res := make([]int, len(temperatures))

	for today, temp := range temperatures {
		for !stk.empty() && temp > temperatures[stk.top()] {
			prev := stk.pop()
			res[prev] = today - prev
		}
		stk.push(today)
	}
	
	return res
}
