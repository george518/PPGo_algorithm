/************************************************************
** @Description: binary_search_tree
** @Author: haodaquan
** @Date:   2018-11-04 22:41
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-11-04 22:41
*************************************************************/
package binary_search_tree

import (
	"fmt"
	"testing"
)

func TestBst_Insert(t *testing.T) {
	bst := new(Bst)
	bst.Nodes = &Node{9, 10, nil, nil}
	Insert(bst.Nodes, 7, 1)
	Insert(bst.Nodes, 2, 3)
	Insert(bst.Nodes, 4, 5)
	Insert(bst.Nodes, 6, 7)
	Insert(bst.Nodes, 8, 9)
	Insert(bst.Nodes, 10, 11)
	//fmt.Println(bst.Nodes)
	bst.PrintBst()

	node1 := Find(bst.Nodes, 8)
	fmt.Println(node1)

	node2 := Find(bst.Nodes, 11)
	fmt.Println(node2)
}
