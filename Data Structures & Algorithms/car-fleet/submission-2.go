type MyStack struct {
	items []float64
}
func (s *MyStack) empty() bool {
	return len(s.items) == 0
}

func (s *MyStack) top() float64 {
	if s.empty() {
		return -1
	}
	return s.items[len(s.items)-1]
}

func (s *MyStack) pop() float64 {
	v := s.top()
	s.items = s.items[0:len(s.items)-1]
	return v
}

func (s *MyStack) push(v float64) {
	s.items = append(s.items, v)
}


func carFleet(target int, position []int, speed []int) int {
	cars := map[int]float64{} // pos: time
	for i, p := range position {
		cars[p] = float64(target-p) / float64(speed[i])
	}
	// fmt.Println(cars)

	stk := &MyStack{}
	i := target // start from target
	for i >= 0 {
		time, ex := cars[i]
		if ex {
			if time > stk.top() {
				stk.push(time)
			} else {
				stk.push(stk.top())
			}
			// fmt.Println(stk.items)
		}
		i--
	}

	res := 0
	uniques := map[float64]int{}
	for _, v := range stk.items {
		_, ex := uniques[v]
		if !ex {
			uniques[v] = 1
			res += 1
		}
	}
	// fmt.Println(uniques)


	return res
}
