import (
	"slices"
)

type MyStack struct {
	items []int
}
func (s *MyStack) top() int {
	return s.items[len(s.items)-1]
}
func (s *MyStack) push(v int) {
	s.items = append(s.items, v)
}
func (s *MyStack) pop() int {
	n := s.top()
	s.items = s.items[:len(s.items)-1]
	return n
}

func evalRPN(tokens []string) int {
	stk := &MyStack{}
	var ch1, ch2 int
	for _, op := range tokens {
		if slices.Contains([]string{"+", "-", "*", "/"}, op) {
			ch2, ch1 = stk.pop(), stk.pop()
			if op == "+" {
				stk.push(ch1 + ch2)
			}
			if op == "-" {
				stk.push(ch1 - ch2)
			}
			if op == "*" {
				stk.push(ch1 * ch2)
			}
			if op == "/" {
				stk.push(ch1 / ch2)
			}
		} else {
			n, _ := strconv.Atoi(op)
			stk.push(n)
		}
		fmt.Println(stk.items)
	}
	return stk.top()
}
