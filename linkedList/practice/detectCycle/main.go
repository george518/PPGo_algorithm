/************************************************************
** @Description: detectCycle
** @Author: george hao
** @Date:   2018-12-14 15:58
** @Last Modified by:  george hao
** @Last Modified time: 2018-12-14 15:58
*************************************************************/
package main

import (
	"fmt"
)

func main() {
	l1 := &ListNode{2, &ListNode{1, &ListNode{5, &ListNode{3, &ListNode{9, &ListNode{7, nil}}}}}}

	//l1.Next.Next.Next.Next.Next.Next = l1.Next

	l := detectCycle(l1)
	Traverse(l)
}

func Traverse(l *ListNode) {
	if l == nil {
		return
	}
	node := l
	for ; node != nil; node = node.Next {
		fmt.Print(node.Val)
		if node.Next != nil {
			fmt.Print("->")
		}

	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	slow := head
	fast := head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			p := head
			for slow != nil {
				if p == slow {
					return slow
				}
				p = p.Next
				slow = slow.Next
			}
		}
	}
	return nil
}
