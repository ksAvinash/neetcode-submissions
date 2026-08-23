/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
    nodes := map[*ListNode]bool{}
	for head != nil {
		_, ex := nodes[head]; if ex {
			return true
		}
		nodes[head] = true
		head = head.Next
	}
	return false
}
