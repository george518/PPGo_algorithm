/************************************************************
** @Description: mergerTwoLists
** @Author: george hao
** @Date:   2018-12-13 09:40
** @Last Modified by:  george hao
** @Last Modified time: 2018-12-13 09:40
*************************************************************/
package main

import (
	"fmt"
)

//两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
//输入：1->2->4, 1->3->4
//输出：1->1->2->3->4->4

//这里假设有序链表都是从小到大排列，合并成的有序链表也是从小到大排列

func main() {
	l1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}

	l1 = &ListNode{10, nil}
	l2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}

	l := mergeTwoLists(l1, l2)

	Traverse(l)
}

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
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
