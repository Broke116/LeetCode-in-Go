package problem0002

// ListNode is the linked list represantation
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	temp := result
	sum := 0
	crr := 0

	for l1 != nil || l2 != nil || crr > 0 {
		sum = crr
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		crr = sum / 10
		temp.Next = &ListNode{Val: sum % 10}
		temp = temp.Next
	}

	return result.Next
}
