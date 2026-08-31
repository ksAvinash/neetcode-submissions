/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
    pos := map[int]*ListNode{}
	
	i, temp := 0, head
	for temp != nil {
		pos[i] = temp
		i++
		temp = temp.Next
	}


	dummy := &ListNode{}
	temp = dummy
	i, j := 0, len(pos)-1
	for i <= j {
		temp.Next = pos[i]
		temp = temp.Next
		i++

		temp.Next = pos[j]
		temp = temp.Next
		j--
	}
	temp.Next = nil

	head = dummy.Next
}
