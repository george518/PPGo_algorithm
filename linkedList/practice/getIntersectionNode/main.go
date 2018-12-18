/************************************************************
** @Description: getIntersectionNode
** @Author: george hao
** @Date:   2018-12-14 17:33
** @Last Modified by:  george hao
** @Last Modified time: 2018-12-14 17:33
*************************************************************/
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{2, &ListNode{2, &ListNode{4, nil}}}

	l2 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	l := getIntersectionNode(l1, l2)
	fmt.Println(l)
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {

	if headA == nil || headB == nil {
		return nil
	}
	node1 := headA
	node2 := headB
	for node1 != node2 {
		if node1 != nil {
			node1 = headB
		} else {
			node1 = node1.Next
		}

		if node2 != nil {
			node2 = headA
		} else {
			node2 = node1.Next
		}
	}
	return node1
}
