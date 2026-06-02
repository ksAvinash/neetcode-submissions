/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func newNode(v int) *ListNode {
	return &ListNode{
		Val: v,
		Next: nil,
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    carry, rem, count := 0, 0, 0
	dummy := &ListNode{}
	t := dummy
	for l1 != nil || l2 != nil {
		count = carry
		if l1 != nil {
			count += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			count += l2.Val
			l2 = l2.Next
		}
		rem = count % 10
		carry = (count - rem) / 10
		n := newNode(rem)
		t.Next = n
		t = t.Next
	}
	if carry > 0 {
		n := newNode(carry)
		t.Next = n
		t = t.Next
	}
	return dummy.Next
}
