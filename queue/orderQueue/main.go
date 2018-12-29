/************************************************************
** @Description: orderQueue 有序队列
** @Author: george hao
** @Date:   2018-12-24 12:53
** @Last Modified by:  george hao
** @Last Modified time: 2018-12-24 12:53
*************************************************************/
package main

import "fmt"

func main() {

	oq := NewOrderQueue(2)
	fmt.Println(oq.Pop())
	oq.Push("hao")
	fmt.Println(oq.Pop())
	oq.Push("da")
	oq.Push("quan")

	fmt.Println(oq)

}

type orderQueue struct {
	data     []string
	size     int
	capacity int
}

func NewOrderQueue(capacity int) *orderQueue {
	return &orderQueue{
		data:     make([]string, 0),
		size:     0,
		capacity: capacity,
	}
}

func (oq *orderQueue) Push(str string) bool {
	if oq.size+1 > oq.capacity {
		return false
	}
	oq.data = append(oq.data, str)
	oq.size++
	return true
}

func (oq *orderQueue) Pop() string {
	if oq.size == 0 {
		return ""
	}

	str := oq.data[0]
	oq.data = oq.data[1:]
	oq.size--
	return str

}
