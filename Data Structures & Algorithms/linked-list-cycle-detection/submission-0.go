/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
    maps := map[*ListNode]int{}
	for head != nil {
		_, ex := maps[head]
		if ex {
			return true
		}
		maps[head] = head.Val
		head = head.Next
	}
	return false
}
