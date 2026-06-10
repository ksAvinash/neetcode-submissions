type MyStack struct {
	items []int
}
func (s *MyStack) empty() bool {
	return len(s.items) == 0
}
func (s *MyStack) push(v int) {
	s.items = append(s.items, v)
}
func (s *MyStack) top() int {
	return s.items[len(s.items)-1]
}
func (s *MyStack) pop() int {
	if s.empty() {
		return 0
	}
	t := s.top()
	s.items = s.items[:len(s.items)-1]
	return t
}


func asteroidCollision(asteroids []int) []int {
	stk := &MyStack{}

	for _, ast := range asteroids {
		if ast > 0 {
			stk.push(ast)
		} else {
			for {
				top := stk.pop()

				if top == 0 { // stack empty
					stk.push(ast)
					break
				}
				if top < 0 { // top is already -ve ast
					stk.push(top)
					stk.push(ast)
					break
				}
				if top > -ast { // top is bigger than ast
					stk.push(top)
					break
				}
				if top == -ast { // top & ast are same
					break
				}
			}
		}
	}

	return stk.items
}
