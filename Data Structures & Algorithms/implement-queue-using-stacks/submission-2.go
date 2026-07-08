type MyQueue struct {
	inStk []int
	outStk []int
}

func Constructor() MyQueue {
	return MyQueue{
		inStk: make([]int, 0),
		outStk: make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {
	this.inStk = append(this.inStk, x)
}

func (this *MyQueue) Pop() int {
	top := this.Peek()
	this.outStk = this.outStk[0:len(this.outStk)-1]
	return top
}

func (this *MyQueue) Peek() int {
	if len(this.outStk) == 0 {
		// copy inStk to outStk
		i := len(this.inStk)-1
		for i >= 0 {
			this.outStk = append(this.outStk, this.inStk[i])
			i--
		}
		this.inStk = []int{}
	}
	return this.outStk[len(this.outStk)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.inStk) == 0 && len(this.outStk) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param2 := obj.Pop();
 * param3 := obj.Peek();
 * param4 := obj.Empty();
 */
