package algo

import (
	"testing"
)

func TestFunc(t *testing.T) {
	var head = &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	// head.Next.Next.Next = &ListNode{Val: 4}
	res := F(head)

	if res == nil {
		println("nil")
		t.Fatal()
	}

	for res.Next != nil {
		print(res.Val, "->")
		res = res.Next
	}
	println(res.Val)
}
