/************************************************************
** @Description: sortList
** @Author: george hao
** @Date:   2018-12-14 12:30
** @Last Modified by:  george hao
** @Last Modified time: 2018-12-14 12:30
*************************************************************/
package main

import (
	"fmt"
)

func main() {
	l1 := &ListNode{2, &ListNode{1, &ListNode{5, &ListNode{3, &ListNode{9, &ListNode{7, nil}}}}}}
	l := sortList(l1)
	Traverse(l)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//递归
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	list1 := head
	list2 := getMid(head)

	list1 = sortList(list1)
	list2 = sortList(list2)

	return mergeTwoLists(list1, list2)
}

func getMid(list *ListNode) *ListNode {

	var fast = list.Next
	var slow = list.Next
	var prev = list

	for {
		if fast == nil {
			break
		}

		fast = fast.Next
		if fast == nil {
			break
		}

		fast = fast.Next

		prev = slow
		slow = slow.Next
	}

	prev.Next = nil

	return slow
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	//处理临界情况
	if l1 == nil && l2 != nil {
		return l2
	} else if l2 == nil && l1 != nil {
		return l1
	} else if l1 == nil && l2 == nil {
		return nil
	}
	node1 := l1
	node2 := l2
	var l = new(ListNode)
	node := l
	for {
		if node1 == nil && node2 != nil {
			node.Val = node2.Val
			node2 = node2.Next
		} else if node1 != nil && node2 == nil {
			node.Val = node1.Val
			node1 = node1.Next
		} else if node1.Val >= node2.Val {
			node.Val = node2.Val
			node2 = node2.Next
		} else if node1.Val < node2.Val {
			node.Val = node1.Val
			node1 = node1.Next
		} else {
			//fmt.Println("no status")
		}

		if node1 == nil && node2 == nil {
			//结束循环条件
			break
		} else {
			node.Next = new(ListNode)
			node = node.Next
		}
	}

	return l
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
