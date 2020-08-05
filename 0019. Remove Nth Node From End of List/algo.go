package algo

type ListNode struct {
	Val  int
	Next *ListNode
}

// 题目保证n有效
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}

	var (
		sentinel    = &ListNode{Val: 0, Next: head}
		tailPointer = sentinel
		destPointer = sentinel
	)

	for n >= 0 {
		tailPointer = tailPointer.Next
		n--
	}

	for tailPointer != nil {
		tailPointer = tailPointer.Next
		destPointer = destPointer.Next
	}

	destPointer.Next = destPointer.Next.Next

	return sentinel.Next
}
