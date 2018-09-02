/************************************************************
** @Description: insertion
** @Author: haodaquan
** @Date:   2018-09-02 15:22
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-09-02 15:22
*************************************************************/
package insertion

import (
	"testing"

	"github.com/george518/PPGo_algorithm/common"
)

//go test -v
func TestSort(t *testing.T) {
	arr := common.RandArray(10, 1000, 9999)
	sortArr := Sort(arr)
	if !common.CheckArray(&sortArr) {
		t.Error("failure")
	} else {
		t.Log("success")
	}
}

//go test -test.bench ".*"
func BenchmarkSort(b *testing.B) {
	b.StopTimer() //调用该函数停止压力测试的时间计数
	min, max := 1, b.N
	len := 10000
	arr := common.RandArray(len, min, max)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Sort(arr)
	}
}

//go test -v
func TestSort2(t *testing.T) {
	arr := common.RandArray(10, 1000, 9999)
	sortArr := Sort2(arr)
	if !common.CheckArray(&sortArr) {
		t.Error("failure")
	} else {
		t.Log("success")
	}
}

//go test -test.bench ".*"
func BenchmarkSort2(b *testing.B) {
	b.StopTimer() //调用该函数停止压力测试的时间计数
	min, max := 1, b.N
	len := 10000
	arr := common.RandArray(len, min, max)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Sort2(arr)
	}
}

//100000	     10126 ns/op
//30	  34741500 ns/op
//20	  55466659 ns/op
//20	  55212608 ns/op
