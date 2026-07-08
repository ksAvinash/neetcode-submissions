type MyStack struct {
	queue []int
}

func Constructor() MyStack {
	return MyStack{
		queue: make([]int, 0),
	}
}

func (this *MyStack) Push(x int) {
	this.queue = append(this.queue, x)
	// fmt.Println(this.queue)
}

func (this *MyStack) Pop() int {
	top := this.Top()
	this.queue = this.queue[0:len(this.queue)-1]
	// fmt.Println(this.queue)
	return top
}

func (this *MyStack) Top() int {
	// fmt.Println(this.queue)
	return this.queue[len(this.queue)-1]
}

func (this *MyStack) Empty() bool {
	// fmt.Println(this.queue)
	return len(this.queue) == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param2 := obj.Pop();
 * param3 := obj.Top();
 * param4 := obj.Empty();
 */
