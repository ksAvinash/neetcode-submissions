type MyQueue struct {
	inStack []int
	outStack []int
}

func Constructor() MyQueue {
	return MyQueue{
		inStack: make([]int, 0),
		outStack: make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {
	this.inStack = append(this.inStack, x)
}

func (this *MyQueue) Pop() int {
	top := this.Peek()
	this.outStack = this.outStack[0:len(this.outStack)-1]
	return top
}

func (this *MyQueue) Peek() int {
	if len(this.outStack) == 0 {
		i := len(this.inStack)-1
		for i >= 0 {
			this.outStack = append(this.outStack, this.inStack[i])
			i--
		}
		this.inStack = []int{}
	}
	return this.outStack[len(this.outStack)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.inStack) == 0 && len(this.outStack) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param2 := obj.Pop();
 * param3 := obj.Peek();
 * param4 := obj.Empty();
 */
