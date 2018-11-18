/************************************************************
** @Description: two_sum
** @Author: haodaquan
** @Date:   2018-11-12 21:41
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-11-12 21:41
*************************************************************/
package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

var bits = 0

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ln := new(ListNode)
	if bits > 0 {
		ln.Val = bits
	}
	//bits = 0

	if l1 == nil && l2 == nil {
		if bits > 0 {
			bits = 0
			return ln
		}
		return nil
	}

	if l1 != nil && l2 == nil {
		ln.Val += l1.Val
		bits = 0
		if ln.Val >= 10 {
			ln.Val -= 10
			bits = 1
		}
		ln.Next = addTwoNumbers(l1.Next, nil)
	}

	if l1 == nil && l2 != nil {
		ln.Val += l2.Val
		bits = 0
		if ln.Val >= 10 {
			ln.Val -= 10
			bits = 1
		}
		ln.Next = addTwoNumbers(nil, l2.Next)
	}

	if l1 != nil && l2 != nil {
		ln.Val += l1.Val + l2.Val
		bits = 0
		if ln.Val >= 10 {
			ln.Val -= 10
			bits = 1
		}
		ln.Next = addTwoNumbers(l1.Next, l2.Next)
	}

	return ln
}

func main() {
	Test_addTwoNumbers()
}

func Test_addTwoNumbers() {
	listNode1 := new(ListNode)
	listNode1.Val = 9
	//listNode1.Next = new(ListNode)
	//listNode1.Next.Val = 9
	//listNode1.Next.Next = new(ListNode)
	//listNode1.Next.Next.Val = 9

	listNode1.PrintLn()
	fmt.Println()

	listNode2 := new(ListNode)
	listNode2.Val = 9
	//listNode2.Next = new(ListNode)
	//listNode2.Next.Val = 8
	//listNode2.Next.Next = new(ListNode)
	//listNode2.Next.Next.Val = 4

	listNode2.PrintLn()
	fmt.Println()

	listNode := addTwoNumbers(listNode1, listNode2)

	listNode.PrintLn()
}

func (ln *ListNode) PrintLn() {
	if ln != nil {
		fmt.Print(ln.Val, " ")
	}

	if ln.Next != nil {
		ln.Next.PrintLn()
	}
}
