package main

type ListNode struct {
	Val int
	Next *ListNode
}

func main(){

}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	count := 1
	temp := head
	for temp.Next != nil {
		count ++
		temp = temp.Next
	}

	if n == count{
		return head.Next
	}

	temp = head
	for i := 1; i < count - n; i ++ {
		temp = temp.Next
	}

	temp.Next = temp.Next.Next

	return head
}