package algo

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		return nil
	}

	var (
		sumList     = &ListNode{}
		currentNode = sumList
		nextL1      = l1
		nextL2      = l2
		carry       = 0
	)

	for nextL1 != nil || nextL2 != nil {
		var sum, x, y int

		if nextL1 != nil {
			x = nextL1.Val
			nextL1 = nextL1.Next
		} else {
			x = 0
		}

		if nextL2 != nil {
			y = nextL2.Val
			nextL2 = nextL2.Next
		} else {
			y = 0
		}

		sum = x + y + carry
		carry = sum / 10

		// 不要使用 currentNode.Val = sum % 10; currentNode.Next = &ListNode{};
		// 否则会导致最后多余一个ListNode{}结点
		currentNode.Next = &ListNode{Val: sum % 10}
		currentNode = currentNode.Next

	}

	// 最高位相加导致的进位，如 753+864 中 7+8导致的进位
	if carry > 0 {
		currentNode.Next = &ListNode{Val: carry}
	}

	return sumList.Next
}
