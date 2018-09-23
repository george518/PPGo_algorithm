/************************************************************
** @Description: heap
** @Author: haodaquan
** @Date:   2018-09-18 00:00
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-09-18 00:00
*************************************************************/
package heap

import "fmt"

type MaxHeap struct {
	Len      int   //元素个数
	Capacity int   //数组容量
	Heap     []int //数组
}

func (m *MaxHeap) Sort() {
	return
}

//初始化
func (m *MaxHeap) init(capacity int) {
	if capacity <= 0 {
		fmt.Println("Capacity must be greater than 1")
		return
	}
	m.Capacity = capacity
	m.Len = 1
	m.Heap = make([]int, 0, capacity)
	m.Heap = append(m.Heap, 0)
	fmt.Println(" maxHeap init completely", m.Heap)
}

func (m *MaxHeap) Insert(n int) {
	//fmt.Println("len i:", m.Len, len(m.Heap))
	m.Heap = append(m.Heap, n)
	m.ShiftUp(m.Len)
	m.Len += 1
	//fmt.Println("len i:", m.Len, len(m.Heap))
}

func (m *MaxHeap) ShiftUp(k int) {
	if m.Len > m.Capacity {
		fmt.Println("Excess capacity", m.Len, m.Capacity)
		return
	}
	for i := k; i > 1; i = i / 2 {
		if m.Heap[i/2] < m.Heap[i] {
			m.Heap[i/2], m.Heap[i] = m.Heap[i], m.Heap[i/2]
		}
	}
}

func (m *MaxHeap) GetMax() int {
	if m.Len > 0 {
		max := m.Heap[1]
		m.Heap[m.Len-1], m.Heap[1] = m.Heap[1], m.Heap[m.Len-1]
		m.Heap = m.Heap[0 : m.Len-1]
		m.Len -= 1
		//fmt.Println("len ", m.Len)
		m.ShiftDown(1)
		return max
	} else {
		return 0
	}
}
func (m *MaxHeap) ShiftDown(n int) {
	//fmt.Println("2*n:", 2*n, " m.Len:", m.Len)
	for 2*n < m.Len {
		j := 2 * n

		if j+1 < m.Len && m.Heap[j+1] > m.Heap[j] {
			j += 1
		}

		if m.Heap[n] >= m.Heap[j] {
			//fmt.Println("m.Heap[n]:", m.Heap[n], " m.Heap[j]:", m.Heap[j])
			break
		}

		m.Heap[n], m.Heap[j] = m.Heap[j], m.Heap[n]
		n = j
	}

}

func (m *MaxHeap) PrintHeap() {
	laynum := 0
	sum := 0
	//fmt.Println("print ", m.Len, len(m.Heap))

	for j := 1; ; j = 2 * j {
		sum += j
		laynum++
		if sum >= m.Len-1 {
			break
		}
	}
	//fmt.Println(sum, laynum)
	//i 为层数，laynum为总层数

	leftSpace := sum
	for i := 1; i <= laynum; i++ {
		n := pow(2, i-1) //每层元素个数

		//每层元素的的索引分为[n,2n)
		for k := n; k < 2*n && k < m.Len; k++ {
			if k == n {
				printSpace(leftSpace)
			} else {
				printSpace(2 * leftSpace)
			}
			fmt.Print("", m.Heap[k], "")
		}
		fmt.Println("")
		for k := n; k < 2*n; k++ {
			if k == n {
				printSpace(leftSpace - 1)
			} else {
				printSpace(2*leftSpace - 2)
			}
			if i != laynum {
				fmt.Print("/", "  \\")
			}

		}
		leftSpace = leftSpace / 2
		fmt.Println("")
	}
	//fmt.Println("print2 ", m.Len, len(m.Heap))
}

func printSpace(n int) {
	str := ""
	for i := 0; i < n; i++ {
		str += " "
	}
	fmt.Print(str)
}

func pow(x, n int) int {
	ret := 1 // 结果初始为0次方的值，整数0次方为1。如果是矩阵，则为单元矩阵。
	for n != 0 {
		if n%2 != 0 {
			ret = ret * x
		}
		n /= 2
		x = x * x
	}
	return ret
}
