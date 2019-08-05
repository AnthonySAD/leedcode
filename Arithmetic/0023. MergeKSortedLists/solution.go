package  main

import "fmt"


type ListNode struct {
	Val int
	Next *ListNode
}

func main(){
	var l1 *ListNode
	l1 = nil
	l2 := &ListNode{
		Val: 2,
	}

	
	lists := []*ListNode{l1, l2}

	ans := mergeKLists(lists)

	fmt.Println("ans:")

	for ans != nil{
		print(ans.Val)
		print("->")
		ans = ans.Next
	}

}

func mergeKLists(lists []*ListNode) *ListNode {
    ans := &ListNode {
		Val: 0,
		Next: nil,
	}
	temp := ans

	count := len(lists)
	
	for key := 0; count > 0; count --{
		if lists[key] == nil {
			lists = append(lists[:key], lists[key + 1:]...)
		}else{
			key ++
		}
	}

	for len(lists) != 0{
		if len(lists) == 1 {
			temp.Next = lists[0]
			break
		}
		
		var minValue, minKey int
		for key, value := range lists{

			if key == 0 {
				minKey = key
				minValue = value.Val
			}else{
				if value.Val < minValue {
					minKey = key
					minValue = value.Val
				}
			}
		}
		temp.Next = &ListNode{
			Val: minValue,
			Next: nil,
		}
		temp = temp.Next
		if lists[minKey].Next == nil {
			lists = append(lists[:minKey], lists[minKey + 1:]...)
		}else{
			lists[minKey] = lists[minKey].Next
		}
	}

	return ans.Next
}