package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	k := 2

	node := ListNode{
		Val: 1,
	}

	node1 := ListNode{
		Val: 2,
	}

	node2 := ListNode{
		Val: 3,
	}

	node3 := ListNode{
		Val: 4,
	}

	node4 := ListNode{
		Val: 5,
	}

	node.Next = &node1
	node1.Next = &node2
	node2.Next = &node3
	node3.Next = &node4

	pringNodeList(&node)
	res := reverseKGroup(&node, k)
	pringNodeList(res)

}

func reverseKGroup(head *ListNode, k int) *ListNode {

	first := ListNode{
		Val:  0,
		Next: head,
	}
	tempHead := &first
	temp := head
	next := head.Next
	tempArr := make([]*ListNode, k)
	current := 1
	key := 0

	for {
		if temp == nil {
			tempHead.Next = next
			break
		}

		tempArr[key] = temp
		key++

		//如果等于0则反转
		if current%k == 0 {
			next = temp.Next
			temp = temp.Next
			key = 0
			for i := k - 1; i >= 0; i-- {
				tempHead.Next = tempArr[i]
				tempHead = tempHead.Next

			}
		} else {
			temp = temp.Next
		}

		current++
	}

	return first.Next
}

func pringNodeList(list *ListNode) {

	fmt.Print("[")
	for list != nil {
		fmt.Print(list.Val)
		fmt.Print(" ")
		list = list.Next
	}
	fmt.Print("]")
}
