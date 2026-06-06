/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || left == right {
        return head
    }

    dummy := &ListNode{Next: head}
    pre := dummy
    for i := 1; i < left; i++ {
        pre = pre.Next
    }

    curr := pre.Next
    var prev *ListNode
    for i := left; i <= right; i++ {
        next := curr.Next
        curr.Next = prev
        prev = curr
        curr = next
    }

    tail := pre.Next
    pre.Next = prev
    tail.Next = curr

    return dummy.Next
}
