package challenges

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	currentList := new(ListNode)
	result := currentList
	value := 0

	for l1 != nil || l2 != nil {
		if l1 != nil {
			value += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			value += l2.Val
			l2 = l2.Next
		}
		currentList.Next = &ListNode{
			value % 10,
			nil,
		}
		currentList = currentList.Next
		value = value / 10
	}

	if value > 0 {
		currentList.Next = &ListNode{value, nil}
	}

	return result.Next
}
