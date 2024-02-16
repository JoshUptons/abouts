package main

/* my solution for the addTwoNumbers LeetCode Question */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	remainder, firstVal := sum(0, l1.Val, l2.Val)
	head := &ListNode{
		Val: firstVal,
	}
	l1 = l1.Next
	l2 = l2.Next

	cur := head
	for l1 != nil && l2 != nil {
		rem, value := sum(remainder, l1.Val, l2.Val)
		cur.Next = &ListNode{
			Val: value,
		}
		remainder = rem
		// iterate lists
		cur = cur.Next
		l1 = l1.Next
		l2 = l2.Next
	}

	// add remaining numbers if the linked lists are not the same length
	for l1 != nil {
		rem, val := sum(remainder, l1.Val, 0)
		remainder = rem
		cur.Next = &ListNode{
			Val: val,
		}
		cur = cur.Next
		l1 = l1.Next
	}
	for l2 != nil {
		rem, val := sum(remainder, l2.Val, 0)
		remainder = rem
		cur.Next = &ListNode{
			Val: val,
		}
		cur = cur.Next
		l2 = l2.Next
	}
	if remainder > 0 {
		cur.Next = &ListNode{
			Val: remainder,
		}
	}
	return head
}

func sum(remainder, n1, n2 int) (int, int) {
	t := n1 + n2 + remainder
	if t >= 10 {
		return 1, t % 10
	}
	return 0, t % 10
}
