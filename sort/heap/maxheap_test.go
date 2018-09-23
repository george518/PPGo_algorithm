/************************************************************
** @Description: heap
** @Author: haodaquan
** @Date:   2018-09-18 23:19
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-09-18 23:19
*************************************************************/
package heap

import (
	"fmt"
	"testing"
)

func TestMaxHeap_Insert(t *testing.T) {
	var mh MaxHeap
	mh.init(100)
	mh.Insert(10)
	mh.Insert(19)
	mh.Insert(11)
	mh.Insert(17)
	mh.Insert(13)
	mh.Insert(14)
	mh.Insert(1)
	mh.Insert(15)
	mh.Insert(16)
	mh.Insert(18)
	mh.Insert(122)
	mh.Insert(133)
	mh.Insert(21)
	mh.Insert(121)

	fmt.Println(mh)
	mh.PrintHeap()

	len := mh.Len
	for i := 1; i < len; i++ {
		fmt.Print(mh.GetMax(), " ")
	}
	mh.PrintHeap()
}
