/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	l, t := 0, head
	for t != nil {
		l++
		t = t.Next
	}
	r, t := l - n, head

	if r == 0 {
		head = head.Next
	} else {
		for r > 1 {
			r--
			t = t.Next
		}

		rem := t.Next
		if rem != nil {
			t.Next = rem.Next
		} else {
			t.Next = nil
		}
	}
	return head
}
