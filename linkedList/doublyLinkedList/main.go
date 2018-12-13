/************************************************************
** @Description: doublyLinkedList //双向链表
** @Author: george hao
** @Date:   2018-12-11 17:22
** @Last Modified by:  george hao
** @Last Modified time: 2018-12-11 17:22
*************************************************************/
package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func main() {

	list := NewList()

	for i := 1; i < 10; i++ {
		linkData := LinkData{i, 100 + i}
		list.Append(linkData)
	}
	fmt.Println("Append")
	list.Traverse()

	fmt.Println("InsertNext")
	node := list.Search(LinkData{2, 102})
	list.InsertNext(node, LinkData{33, 3333})
	list.Traverse()
	fmt.Println()

	fmt.Println("InsertPre")
	node = list.Search(LinkData{4, 104})
	list.InsertNext(node, LinkData{44, 444})
	list.Traverse()
	fmt.Println()

	fmt.Println("Remove")
	node = list.Search(LinkData{2, 102})
	list.Remove(node)
	list.Traverse()

	fmt.Println("GetHead:", list.GetHead())
	fmt.Println("GetSize:", list.GetSize())
	fmt.Println("GetTail:", list.GetTail())

}

type LinkedList interface {
	//公共接口
	GetSize() int
	GetHead() *DoublyLinkNode
	GetTail() *DoublyLinkNode
	Append(data LinkData) error
	InsertNext(node *DoublyLinkNode, data LinkData) error
	InsertPrev(node *DoublyLinkNode, data LinkData) error
	Remove(node *DoublyLinkNode) *LinkData
	Search(linkData LinkData) *DoublyLinkNode

	//私有接口
	isHead(node *DoublyLinkNode) bool
	isTail(node *DoublyLinkNode) bool

	//打印
	Traverse()
}

type NodeItf interface {
	//节点接口
	GetData() LinkData
	GetPrev() *DoublyLinkNode
	GetNext() *DoublyLinkNode
}

type LinkData struct {
	Flag    int
	Element interface{}
}

func Equal(ld LinkData, ld2 LinkData) bool {
	return ld.Flag == ld2.Flag
}

type DoublyLinkNode struct {
	Data LinkData
	Next *DoublyLinkNode
	Pre  *DoublyLinkNode
}

type List struct {
	Size int
	Head *DoublyLinkNode
	Tail *DoublyLinkNode
}

type Te struct{}

func (t *Te) GetSize() int {
	return 0
}

func NewList() (list *List) {
	list = new(List)
	list.Size = 0
	list.Head = nil
	list.Tail = nil
	return
}

//私有函数
func (list *List) isHead(node *DoublyLinkNode) bool {
	return Equal(list.GetTail().Data, node.Data)
}

func (list *List) isTail(node *DoublyLinkNode) bool {
	return Equal(list.GetHead().Data, node.Data)
}

//链表函数
func (list *List) GetTail() *DoublyLinkNode {
	return list.Tail
}

func (list *List) GetHead() *DoublyLinkNode {
	return list.Head
}

func (list *List) GetSize() int {
	return list.Size
}

func (list *List) Search(linkData LinkData) *DoublyLinkNode {
	if list.GetSize() == 0 {
		return nil
	}

	node := list.GetHead()
	for ; node != nil; node = node.GetNext() {
		if Equal(node.Data, linkData) {
			break
		}
	}
	return node
}

func (list *List) Append(data LinkData) error {
	newNode := new(DoublyLinkNode)
	newNode.Data = data

	if list.GetSize() == 0 {
		newNode.Pre = nil
		newNode.Next = nil
		list.Head = newNode
		list.Tail = newNode
	} else {
		newNode.Pre = list.Tail
		newNode.Next = nil
		list.Tail.Next = newNode
		list.Tail = newNode
	}
	list.Size++
	return nil
}

func (list *List) InsertNext(node *DoublyLinkNode, data LinkData) error {
	if node == nil {
		return errors.New("node is nil")
	}

	if list.isTail(node) {
		list.Append(data)
	} else {
		newNode := new(DoublyLinkNode)

		newNode.Data = data
		newNode.Pre = node
		newNode.Next = node.Next

		node.Next = newNode
		newNode.Next.Pre = newNode

		list.Size++
	}
	return nil
}

func (list *List) InsertPrev(node *DoublyLinkNode, data LinkData) error {
	if node == nil {
		return errors.New("node is nil")
	}

	if list.isHead(node) {
		newNode := new(DoublyLinkNode)
		newNode.Data = data
		newNode.Next = list.GetHead()
		newNode.Pre = nil

		list.Head.Pre = newNode
		list.Head = newNode
		list.Size++
		return nil
	}

	prev := node.Pre
	return list.InsertNext(prev, data)
}

func (list *List) Remove(node *DoublyLinkNode) *LinkData {
	if node == nil {
		return nil
	}

	prev := node.Pre
	next := node.Next

	if list.isHead(node) {
		list.Head = next
	} else {
		prev.Next = next
	}

	if list.isTail(node) {
		list.Tail = prev
	} else {
		next.Pre = prev
	}

	list.Size--
	linkData := node.GetData()
	return &linkData
}

//节点相关
func (dln *DoublyLinkNode) GetData() LinkData {
	return dln.Data
}

func (dln *DoublyLinkNode) GetPrev() *DoublyLinkNode {
	return dln.Pre
}

func (dln *DoublyLinkNode) GetNext() *DoublyLinkNode {
	return dln.Next
}

func (list *List) Traverse() {

	node := list.GetHead()
	for ; node != nil; node = node.Next {
		//if node.Pre != nil {
		//	fmt.Print(node.Pre.Data, " ")
		//} else {
		//	fmt.Print("{nil,nil} ")
		//}
		fmt.Print(node.Data, " ")

		//if node.Next != nil {
		//	fmt.Print(node.Next.Data, " ")
		//} else {
		//	fmt.Print("{nil,nil} ")
		//}
		fmt.Println()
	}
}
