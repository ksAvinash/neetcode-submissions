type MyQueue struct {
	elements []int
}
func (this *MyQueue) Empty() bool {
	return len(this.elements) == 0
}
func (this *MyQueue) Enqueue(x int) {
	this.elements = append(this.elements, x)
}
func (this *MyQueue) Peek() int {
	return this.elements[0]
}
func (this *MyQueue) Dequeue() int {
	top := this.Peek()
	this.elements = this.elements[1:]
	return top
}

type MyStack struct {
	queue MyQueue
}

func Constructor() MyStack {
	return MyStack{
		queue: MyQueue{
			elements: make([]int, 0),
		},
	}
}

func (this *MyStack) Push(x int) {
	this.queue.Enqueue(x)
	size := len(this.queue.elements)
	for range size-1 {
		v := this.queue.Dequeue()
		this.queue.Enqueue(v)
	}
}

func (this *MyStack) Pop() int {
	top := this.queue.Dequeue()
	return top
}

func (this *MyStack) Top() int {
	return this.queue.Peek()
}

func (this *MyStack) Empty() bool {
	return this.queue.Empty()
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param2 := obj.Pop();
 * param3 := obj.Top();
 * param4 := obj.Empty();
 */
