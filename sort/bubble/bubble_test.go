/************************************************************
** @Description: bubble
** @Author: haodaquan
** @Date:   2018-08-29 22:13
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-08-29 22:13
*************************************************************/
package bubble

import (
	"testing"

	"fmt"

	"github.com/george518/PPGo_algorithm/common"
)

//go test -v
func TestSort(t *testing.T) {
	arr := common.RandArray(10, 1000, 9999)
	sortArr := Sort(arr)
	fmt.Println(sortArr)
	if !common.CheckArray(&sortArr) {
		t.Error("failure")
	} else {
		t.Log("success")
	}
}

//go test -v
func TestSort2(t *testing.T) {
	arr := common.RandArray(10, 1000, 9999)
	sortArr := Sort2(arr)
	fmt.Println(sortArr)
	if !common.CheckArray(&sortArr) {
		t.Error("failure")
	} else {
		t.Log("success")
	}
}

//go test -v
func TestSort3(t *testing.T) {
	arr := common.RandArray(100, 1, 1000)

	sortArr := Sort3(arr)
	if !common.CheckArray(&sortArr) {
		t.Error("failure")
	} else {
		t.Log("success")
	}
}

//go test -v
func TestSort4(t *testing.T) {
	arr := common.RandArray(100, 1, 1000)

	sortArr := Sort4(arr)
	if !common.CheckArray(&sortArr) {
		t.Error("failure")
	} else {
		t.Log("success")
	}
}

//go test -v
func TestSort5(t *testing.T) {
	arr := common.RandArray(100, 1, 1000)

	sortArr := Sort5(arr)
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

//go test -test.bench ".*"
func BenchmarkSort3(b *testing.B) {
	b.StopTimer() //调用该函数停止压力测试的时间计数
	min, max := 1, b.N
	len := 10000
	arr := common.RandArray(len, min, max)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Sort3(arr)
	}
}

//go test -test.bench ".*"
func BenchmarkSort4(b *testing.B) {
	b.StopTimer() //调用该函数停止压力测试的时间计数
	min, max := 1, b.N
	len := 10000
	arr := common.RandArray(len, min, max)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Sort4(arr)
	}
}

//go test -test.bench ".*"
func BenchmarkSort5(b *testing.B) {
	b.StopTimer() //调用该函数停止压力测试的时间计数
	min, max := 1, b.N
	len := 10000
	arr := common.RandArray(len, min, max)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Sort5(arr)
	}
}
