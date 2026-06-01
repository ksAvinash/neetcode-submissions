/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func print(head *ListNode) {
	t := head
	for t != nil {
		fmt.Print(t.Val, " -> ")
		t = t.Next
	}
	fmt.Println("NULL")
}

func reorderList(head *ListNode) {
    pos := map[int]*ListNode{}

	l, t := 0, head
	for t != nil {
		pos[l] = t
		l++
		t = t.Next
	}

	dummy := &ListNode{
		Next: nil,
	}
	i, j, t := 0, l-1, dummy
	for i <= j {
		t.Next = pos[i]
		t = t.Next
		t.Next = pos[j]
		t = t.Next
		i++
		j--
	}
	t.Next = nil
	dummy.Next = head
}
