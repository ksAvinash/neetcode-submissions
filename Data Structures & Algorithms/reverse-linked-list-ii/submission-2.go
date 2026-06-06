/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 
func reverseBetween(head *ListNode, left int, right int) *ListNode {
    if left == right {
		return head
	}

	var preList, listSt, listEn, postList *ListNode
	i := 1
	preList = &ListNode{
		Next: head,
	}
	
	for i < left {
		preList = preList.Next
		i++
	}
	

	var curr, prev, next *ListNode
	listEn = preList.Next

	curr = preList.Next
	for i <= right {
		next = curr.Next
		curr.Next = prev

		prev = curr
		curr = next
		i++
	}
	listSt = prev
	postList = curr

	preList.Next = listSt
	if preList.Val == 0 {
		head = preList.Next
	}
	listEn.Next = postList

	return head
}
