/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func newNode(v int) *Node {
	return &Node{
		Val: v,
	}
}

func insertAtHead(head *Node, v int) *Node {
	n := newNode(v)
	n.Next = head
	head = n
	return head
}

func getItem(head *Node, k int) *Node {
	for k > 0 {
		head = head.Next
		k--
	}
	return head
}

func copyRandomList(head *Node) *Node {
    items := map[int][]int{} // <index>: [val, <index> of random item]
	t := head
	for t != nil {
		v := []int{t.Val}
		if t.Random != nil {
			t2, k := head, 0
			for t2 != nil {
				if t2 == t.Random {
					v = append(v, k)
					break
				}
				t2 = t2.Next
				k++
			}
		}
		items[len(items)] = v
		t = t.Next
	}
	// fmt.Println(items)

	// create initial list
	i := len(items)-1
	var head2 *Node
	for i >= 0 {
		head2 = insertAtHead(head2, items[i][0])
		i--
	}

	// search and update random links
	i, t = 0, head2
	for i < len(items) {
		if len(items[i]) == 2 {
			t.Random = getItem(head2, items[i][1])
			// fmt.Println(t)
		}
		t = t.Next
		i++
	}

	return head2
}
