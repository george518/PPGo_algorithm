/************************************************************
** @Description: singleLinkedList
** @Author: george hao
** @Date:   2018-12-10 17:49
** @Last Modified by:  george hao
** @Last Modified time: 2018-12-10 17:49
*************************************************************/
package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func main() {

	var head = &LinkNode{LinkData{0, 100}, nil}

	for i := 0; i < 10; i++ {
		linkData := LinkData{i + 1, 100 + i + 1}
		head.Append(linkData)
	}
	fmt.Println("Append")
	head.Traverse()

	fmt.Println("Insert:")
	head.Insert(1, LinkData{44, 444})
	head.Traverse()

	fmt.Println("Delete:")
	head.Delete(1)
	head.Traverse()

	var linkData = LinkData{6, 106}
	index := head.Search(linkData)
	fmt.Println("Search:", linkData)
	fmt.Println(index)

	ld, err := head.GetData(8)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("GetData:")
		fmt.Println(ld)
	}

}

type LinkedList interface {
	Append(data LinkData) error
	Insert(index int, data LinkData) error
	Delete(index int) error
	GetLen() (int, error)
	Search(data LinkData) int
	GetData(index int) (*LinkData, error)
	Traverse()
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

func (head *LinkNode) Append(data LinkData) error {
	point := head
	for point.Next != nil {
		point = point.Next
	}
	var node = new(LinkNode)
	node.Data = data
	node.Next = nil
	point.Next = node
	return nil
}

func (head *LinkNode) Insert(index int, data LinkData) error {

	var linkNode = new(LinkNode)
	point := head
	//如果是首节点
	if index == 0 {
		return errors.New("insert head node is not allowed")
	}

	len, err := head.GetLen()

	i := 0
	if err != nil {
		return err
	}

	if index < 0 || index > len-1 {
		return errors.New("index cross boundary ")
	}

	for point.Next != nil {
		if i+1 >= len-1 {
			return errors.New("index can not find")
		}
		if index == i+1 {
			//insert
			linkNode.Next = point.Next
			linkNode.Data = data
			point.Next = linkNode
			return nil
		}

		i++
		point = point.Next

	}
	return errors.New("failure")
}

//删除复杂度为O(1),查找复杂度为O(n)
func (head *LinkNode) Delete(index int) error {
	len, err := head.GetLen()
	if err != nil {
		return err
	}

	//如果是首节点
	if index == 0 {
		return errors.New("delete head node is not allowed")
	}
	if index < 0 || index > len {
		return errors.New("out of index")
	}

	point := head
	for i := 0; i < index-1; i++ {
		point = point.Next
	}
	point.Next = point.Next.Next

	return nil
}

func (head *LinkNode) GetLen() (int, error) {
	point := head
	var len int
	len = 0
	for point.Next != nil {
		len++
		point = point.Next
	}
	return len, nil
}

func (head *LinkNode) Search(data LinkData) int {
	point := head
	index := 0

	len, err := head.GetLen()
	if err != nil {
		return -1
	}

	for point.Next != nil {
		if Equal(data, point.Data) {
			return index
		}

		point = point.Next
		index++

		if index > len-1 {
			return -1
		}
	}

	return -1
}

func (head *LinkNode) GetData(index int) (*LinkData, error) {
	point := head
	len, err := head.GetLen()
	if err != nil {
		return nil, err
	}
	if index < 0 || index > len {
		return nil, errors.New("out of index")
	}

	for i := 0; i < index; i++ {
		point = point.Next
	}

	return &point.Data, nil
}

func (head *LinkNode) Traverse() {
	point := head
	index := 0
	for point.Next != nil {
		fmt.Printf("(%d %d %d)  ", index, point.Data.Flag, point.Data.Element)
		point = point.Next
		index++
	}
	fmt.Println()
}
