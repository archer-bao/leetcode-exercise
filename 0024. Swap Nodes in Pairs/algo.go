package algo

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	var dummy = &ListNode{Next: head}
	var cur = dummy

	for head != nil && head.Next != nil {
		secondNode := head.Next

		cur.Next = secondNode
		head.Next = secondNode.Next
		secondNode.Next = head

		cur = head
		head = head.Next
	}

	return dummy.Next
}
