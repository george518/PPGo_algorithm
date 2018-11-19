/************************************************************
** @Description: 二分搜索树
** @Author: haodaquan
** @Date:   2018-11-02 01:14
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-11-02 01:14
*************************************************************/
package binary_search_tree

import (
	"fmt"
)

type Node struct {
	Key   int
	Value int
	Left  *Node
	Right *Node
}
type Bst struct {
	Nodes *Node //节点
	Count int   //节点数
}

func Insert(node *Node, key int, value int) *Node {
	if node == nil {
		return &Node{key, value, nil, nil}
	}
	switch {
	case node.Key == key:
		node.Value = value
		return node
	case node.Key > key:
		node.Left = Insert(node.Left, key, value)
		return node
	case node.Key < key:
		node.Right = Insert(node.Right, key, value)
		return node
	}
	return node
}

func Find(node *Node, key int) *Node {
	if node == nil {
		return nil
	}
	switch {
	case node.Key == key:
		return node
	case node.Key > key:
		return Find(node.Left, key)
	case node.Key < key:
		return Find(node.Right, key)
	}
	return nil
}

func (b *Bst) PrintBst() {
	b.Nodes.printNode()
}

func (n *Node) printNode() {
	fmt.Println(n.Key, n.Value, n.Left, n.Right)
	if n.Left != nil {
		n.Left.printNode()
	}

	if n.Right != nil {
		n.Right.printNode()
	}
}
