/************************************************************
** @Description: reverseList
** @Author: george hao
** @Date:   2018-12-13 14:52
** @Last Modified by:  george hao
** @Last Modified time: 2018-12-13 14:52
*************************************************************/
package main

import (
	"fmt"
)

//反转一个单链表。
//输入: 1->2->3->4->5->NULL
//输出: 5->4->3->2->1->NULL
//
//你可以迭代或递归地反转链表。你能否用两种方法解决这道题？

func main() {

	l1 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	l := reverseList(l1)
	Traverse(l)
}

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

//迭代
func reverseList(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}
	var preNode, nextNode *ListNode
	nextNode = new(ListNode)
	nextNode = nil
	preNode = nil

	for {
		if head == nil {
			break
		}
		nextNode = head.Next
		head.Next = preNode
		preNode = head
		head = nextNode
	}

	return preNode
}

//递归
func reverseList2(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	var preNode, nextNode *ListNode
	nextNode = nil
	preNode = nil

	for {
		if head == nil {
			break
		}
		nextNode = head.Next
		head.Next = preNode
		preNode = head
		head = nextNode
	}

	return preNode
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
