/************************************************************
** @Description: heap
** @Author: haodaquan
** @Date:   2018-09-18 23:19
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-09-18 23:19
*************************************************************/
package index_heap

import (
	"fmt"
	"testing"
)

func TestMaxHeap_Insert(t *testing.T) {
	var mh IndexMaxHeap
	mh.init(100)
	mh.Insert(2)
	mh.Insert(3)
	mh.Insert(1)
	mh.Insert(9)
	mh.Insert(4)
	mh.Insert(8)
	mh.Insert(7)
	mh.Insert(5)
	mh.Insert(6)

	index := make([]int, 0)
	for i := 0; i < mh.Len; i++ {
		index = append(index, i)
	}
	fmt.Println("索   引：", index)
	fmt.Println("索引数组：", mh.Index)
	fmt.Println("元   素：", mh.Heap)

	len := mh.Len
	fmt.Print("排序：")
	for i := 1; i < len; i++ {
		fmt.Print(mh.GetMax(), " ")
	}
}
