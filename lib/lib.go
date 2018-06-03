/************************************************************
** @Description: lib
** @Author: haodaquan
** @Date:   2018-06-01 22:27
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-06-01 22:27
*************************************************************/
package lib

import (
	"fmt"
	"math/rand"
	"time"
)

type ISort interface {
	Sort()
	GetFuncName() string
}

type SortEntity struct {
	Arr     []int
	Len     int
	RunTime time.Duration
}

//生成随机数组
func RandArray(n, min, max int) []int {

	arr := make([]int, n)
	if n <= 0 || min > max {
		return arr
	}
	rand.Seed(time.Now().UnixNano()) //利用当前时间的UNIX时间戳初始化rand包
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(max)%(max-min+1) + min
	}
	return arr
}

//测试排序性能
func RunSort(sortFuc ISort) {
	t1 := time.Now()
	sortFuc.Sort()
	elapsed := time.Since(t1)
	fmt.Println(sortFuc.GetFuncName(), " elapsed: ", elapsed)
}

func (se *SortEntity) Print(start, end int) {
	if se.Len < end || start < 0 {
		fmt.Println("error：wrong range")
	}
	for i := start; i <= end; i++ {
		fmt.Printf(" %v ", se.Arr[i])
	}
}
