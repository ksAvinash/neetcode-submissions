type MyStack struct {
	items []int
}
func (s *MyStack) push(ch string) {
	v, _ := strconv.Atoi(ch)
	s.items = append(s.items, v)
}
func (s *MyStack) pop() int {
	if s.empty() {
		return 0
	}
	t := s.top()
	s.items = s.items[:len(s.items)-1]
	return t
}
func (s *MyStack) top() int {
	if s.empty() {
		return 0
	}
	return s.items[len(s.items)-1]
}
func (s *MyStack) empty() bool {
	return len(s.items) == 0
}


func evalRPN(tokens []string) int {
	stk := MyStack{}

	var res, op1, op2 int
	for _, ch := range tokens {
		switch ch {
		case "+":
			op2, op1 = stk.pop(), stk.pop()
			res = op1 + op2
			stk.push(strconv.Itoa(res))
		case "*":
			op2, op1 = stk.pop(), stk.pop()
			res = op1 * op2
			stk.push(strconv.Itoa(res))
		case "-":
			op2, op1 = stk.pop(), stk.pop()
			res = op1 - op2
			stk.push(strconv.Itoa(res))
		case "/":
			op2, op1 = stk.pop(), stk.pop()
			res = op1 / op2
			stk.push(strconv.Itoa(res))
		default:
			stk.push(ch)
		}
		// fmt.Println(stk.items)
	}
	return stk.top()
}
