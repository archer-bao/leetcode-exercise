package algo

import (
	"math/rand"
	"strings"
	"testing"
)

func TestTwoNumbers(t *testing.T) {
	linkList1 := makeList()
	linkList2 := makeList()

	res := addTwoNumbers(linkList1, linkList2)
	if res == nil {
		t.Fatal("empty result")
	}

	print("第一链表: ")
	traverse(linkList1)

	print("第二链表: ")
	traverse(linkList2)

	print("结果链表: ")
	traverse(res)

	println(strings.Repeat("=", 30))

	sameList := makeList()
	res = addTwoNumbers(sameList, sameList)
	if res == nil {
		t.Fatal("empty result")
	}

	print("同一链表: ")
	traverse(sameList)

	print("结果链表: ")
	traverse(res)
}

// 创建单链表
func makeList() *ListNode {
	var listNode = &ListNode{Val: rand.Intn(9) + 1}
	cnt := rand.Intn(6) + 3
	for i := 1; i < cnt; i++ {
		l := listNode
		for {
			if l.Next == nil {
				l.Next = &ListNode{Val: rand.Intn(9) + 1}
				break
			}
			l = l.Next
		}
	}
	return listNode
}

// 遍历单链表
func traverse(linkList *ListNode) {
	iter := linkList
	for {
		if iter != nil {
			if iter.Next == nil {
				print(iter.Val)
			} else {
				print(iter.Val, "->")
			}
			iter = iter.Next
		} else {
			break
		}
	}
	println()
}
