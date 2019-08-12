package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}


func main (){
	l1 := &ListNode{
		Val: 1,
	}

	l2 := &ListNode{
		Val: 2,
	}

	l1.Next = l2

	ans := swapPairs(l1)
	fmt.Println(ans)
}

func swapPairs(head *ListNode) *ListNode {
    pre := &ListNode{
		Next: head,
	}

    head = pre;

	for head.Next != nil && head.Next.Next != nil {
		start := head.Next
        end := head.Next.Next
        head.Next = end
        start.Next = end.Next
        head.Next.Next = start
        head = head.Next.Next
	}
	
	return pre.Next
}