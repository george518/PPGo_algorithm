/************************************************************
** @Description: circularLinkedList
** @Author: george hao
** @Date:   2018-12-11 13:45
** @Last Modified by:  george hao
** @Last Modified time: 2018-12-11 13:45
*************************************************************/
package main

import (
	"errors"
	"fmt"
)

func main() {

	list := NewList()

	for i := 1; i < 10; i++ {
		linkData := LinkData{i, 100 + i}
		list.Append(&LinkNode{linkData, nil})
	}
	fmt.Println("Append")
	list.Traverse()
	fmt.Println()

	fmt.Println("Insert:")
	list.Insert(1, &LinkNode{
		LinkData{44, 444}, nil})
	list.Traverse()
	//
	fmt.Println("DeleteIndex:")
	list.DeleteIndex(1)
	list.Traverse()
	fmt.Println()

	fmt.Println("Search:")
	index := list.Search(&LinkNode{LinkData{8, 108}, nil})
	list.Traverse()
	fmt.Println("index :", index)
	fmt.Println()

	fmt.Println("DeleteNode:")
	list.Delete(&LinkNode{LinkData{9, 109}, nil})
	list.Traverse()
}

type LinkData struct {
	Flag    int
	Element interface{}
}

func Equal(ld LinkData, ld2 LinkData) bool {
	return ld.Flag == ld2.Flag
}

type LinkNode struct {
	Data LinkData
	Next *LinkNode
}

type List struct {
	tail *LinkNode
	size int
}

//初始化一个空单循环链表
func NewList() (list *List) {
	list = new(List)
	list.tail = nil
	list.size = 0
	return
}

func (list *List) Append(node *LinkNode) error {
	if node == nil {
		return errors.New("failure")
	}
	//如果链表长度为0，将node赋值一下
	if list.size == 0 {
		node.Next = node
	} else {
		//如果长度大于0，首先寻找要插入节点的上一个节点
		preData := list.tail.Next
		node.Next = preData
		list.tail.Next = node
	}
	list.tail = node
	list.size++
	return nil
}

func (list *List) Insert(index int, listNode *LinkNode) error {
	if listNode == nil {
		return errors.New("listNode is nil")
	}

	if list.size == 0 || list.size == index {
		list.Append(listNode)
	} else {
		var preData *LinkNode

		if index == 0 {
			preData = list.tail
		} else {
			preData = list.GetData(index - 1)
			if list.size == index {
				list.tail = listNode
			}
		}

		listNode.Next = preData.Next
		preData.Next = listNode
		list.size++
	}
	return nil
}

func (list *List) Delete(linkNode *LinkNode) error {
	node := list.tail
	size := list.size
	for i := 0; i < size; i++ {
		if Equal(node.Data, linkNode.Data) {
			if i == 0 {
				list.DeleteIndex(list.size - 1)
			} else {
				list.DeleteIndex(i)
			}
			return nil
		}

		node = node.Next
	}
	return nil
}

func (list *List) DeleteIndex(index int) error {

	if list.size == 0 {
		return nil
	}

	if list.size == 1 {
		list.tail = nil
		list.size = 0
		return nil
	} else {

		var preNode, currentNode *LinkNode
		if index == 0 {
			preNode = list.tail
		} else {
			preNode = list.GetData(index - 1)
		}
		currentNode = preNode.Next
		preNode.Next = currentNode.Next

		//头尾相连
		if index == list.size-1 {
			list.tail = preNode
		}
		currentNode.Next = nil
		currentNode = nil
	}

	list.size--

	return nil
}

func (list *List) Search(listNode *LinkNode) int {
	node := list.tail
	size := list.size
	for i := 0; i < size; i++ {

		if Equal(node.Data, listNode.Data) {
			return i
		}

		node = node.Next
	}

	return -1
}

func (list *List) GetData(index int) (linkNode *LinkNode) {
	linkNode = list.tail
	for i := 0; i <= index; i++ {
		linkNode = linkNode.Next
	}
	return
}

func (list *List) Traverse() {

	point := list.tail
	for i := 0; i < list.size; i++ {
		fmt.Printf("(%d,%d)  ", point.Data.Flag, point.Data.Element)
		point = point.Next
	}
	fmt.Println()

}
