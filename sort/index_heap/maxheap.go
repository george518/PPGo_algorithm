/************************************************************
** @Description: heap
** @Author: haodaquan
** @Date:   2018-09-18 00:00
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-09-18 00:00
*************************************************************/
package index_heap

import "fmt"

type IndexMaxHeap struct {
	Len      int   //元素个数
	Capacity int   //数组容量
	Heap     []int //元素数组
	Index    []int //索引数组
}

func (m *IndexMaxHeap) Sort() {
	return
}

//初始化
func (m *IndexMaxHeap) init(capacity int) {
	if capacity <= 0 {
		fmt.Println("Capacity must be greater than 1")
		return
	}
	m.Capacity = capacity
	m.Len = 1
	m.Heap = make([]int, 0, capacity)
	m.Heap = append(m.Heap, 0)
	m.Index = make([]int, 0, capacity)
	m.Index = append(m.Index, 0)
	fmt.Println(" maxHeap init completely", m.Heap, m.Index)
}

func (m *IndexMaxHeap) Insert(n int) {
	//fmt.Println("len i:", m.Len, len(m.Heap))
	m.Heap = append(m.Heap, n)
	m.Index = append(m.Index, m.Len)
	m.ShiftUp(m.Len)
	m.Len += 1
	//fmt.Println("len i:", m.Len, len(m.Heap))
}

func (m *IndexMaxHeap) ShiftUp(k int) {
	if m.Len > m.Capacity {
		fmt.Println("Excess capacity", m.Len, m.Capacity)
		return
	}
	for i := k; i > 1; i = i / 2 {
		if m.Heap[m.Index[i/2]] < m.Heap[m.Index[i]] {
			m.Index[i/2], m.Index[i] = m.Index[i], m.Index[i/2]
		}
	}

	//fmt.Println(m.Index)
	//fmt.Println(m.Heap)
}

func (m *IndexMaxHeap) ShiftDown(n int) {
	//fmt.Println("2*n:", 2*n, " m.Len:", m.Len)
	for 2*n < m.Len {
		j := 2 * n

		if j+1 < m.Len && m.Heap[m.Index[j+1]] > m.Heap[m.Index[j]] {
			j += 1
		}

		if m.Heap[m.Index[n]] >= m.Heap[m.Index[j]] {
			//fmt.Println("m.Heap[n]:", m.Heap[n], " m.Heap[j]:", m.Heap[j])
			break
		}

		m.Index[n], m.Index[j] = m.Index[j], m.Index[n]
		n = j
	}

}

func (m *IndexMaxHeap) GetMax() int {
	if m.Len > 0 {
		max := m.Heap[m.Index[1]]
		m.Index[m.Len-1], m.Index[1] = m.Index[1], m.Index[m.Len-1]
		m.Index = m.Index[0 : m.Len-1]
		m.Len -= 1
		//fmt.Println("len ", m.Len)
		m.ShiftDown(1)

		return max
	} else {
		return 0
	}
}
