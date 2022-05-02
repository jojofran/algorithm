package array

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

//https://leetcode-cn.com/problems/add-two-numbers/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	dummy := &ListNode{Next: &ListNode{}}
	l3 := dummy

	carry := 0
	val := 0

	for l1 != nil || l2 != nil {
		l3 = l3.Next

		if l1 != nil && l2 == nil {
			val = l1.Val + carry
			l1 = l1.Next

		} else if l1 == nil && l2 != nil {
			val = l2.Val + carry
			l2 = l2.Next

		} else {
			val = l1.Val + l2.Val + carry
			l1 = l1.Next
			l2 = l2.Next
		}

		l3.Val = val % 10
		carry = val / 10

		if l1 != nil || l2 != nil || carry > 0 {
			l3.Next = &ListNode{}
		}
	}

	if carry > 0 {
		l3.Next = &ListNode{Val: carry}
	}

	return dummy.Next
}
