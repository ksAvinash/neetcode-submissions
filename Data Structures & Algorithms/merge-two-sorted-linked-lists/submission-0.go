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


func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	var head *ListNode
	if list1.Val <= list2.Val {
		head = newNode(list1.Val)
		list1 = list1.Next
	} else {
		head = newNode(list2.Val)
		list2 = list2.Next
	}

	temp := head
	for list1 != nil && list2 != nil {
		if list1.Val == list2.Val {
			n, k := newNode(list1.Val), newNode(list2.Val)
			temp.Next = n
			temp = temp.Next
			temp.Next = k
			temp = temp.Next
			list1 = list1.Next
			list2 = list2.Next
		} else if list1.Val <= list2.Val {
			n := newNode(list1.Val)
			temp.Next = n
			temp = temp.Next
			list1 = list1.Next
		} else {
			n := newNode(list2.Val)
			temp.Next = n
			temp = temp.Next
			list2 = list2.Next
		}
	}
	for list1 != nil {
		n := newNode(list1.Val)
		temp.Next = n
		temp = temp.Next
		list1 = list1.Next
	}
	for list2 != nil {
		n := newNode(list2.Val)
		temp.Next = n
		temp = temp.Next
		list2 = list2.Next
	}
	return head
}
