package solution

type ListNode struct {
	 Val int
	 Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	current := head
	for l1 != nil && l2 != nil{
		temp := new(ListNode)
		if l1.Val < l2.Val{
			temp.Val = l1.Val
			l1 = l1.Next
		}else{
			temp.Val = l2.Val
			l2 = l2.Next
		}
		current.Next = temp
		current = current.Next
	}

	if l1 == nil{
		current.Next = l2
	}else{
		current.Next = l1
	}

	return head.Next
}