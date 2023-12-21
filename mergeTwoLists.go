package main

import "fmt"

func main() {
	n1 := ListNode{0, nil}
	n2 := ListNode{1, nil}
	n3 := ListNode{2, nil}
	n4 := ListNode{2, nil}
	n5 := ListNode{2, nil}

	n6 := ListNode{1, nil}
	n7 := ListNode{1, nil}
	n8 := ListNode{1, nil}
	n9 := ListNode{2, nil}
	n10 := ListNode{5, nil}

	n1.Next = &n2
	n2.Next = &n3
	n3.Next = &n4
	n4.Next = &n5

	n6.Next = &n7
	n7.Next = &n8
	n8.Next = &n9
	n9.Next = &n10
	list1 := n1
	list2 := n6
	l := mergeTwoLists(&list1, &list2)
	printList(l)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func printList(l *ListNode) {
	cur := l
	for cur.Next != nil {
		fmt.Println(cur.Val)
		cur = cur.Next
	}
	fmt.Println(cur.Val)
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	var cur *ListNode
	if list1.Val >= list2.Val {
		cur = list2
		list2 = list2.Next
	} else {
		cur = list1
		list1 = list1.Next
	}

	cur.Next = mergeTwoLists(list1, list2)

	return cur

}
