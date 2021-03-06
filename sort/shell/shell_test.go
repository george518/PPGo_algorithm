/************************************************************
** @Description: shell
** @Author: haodaquan
** @Date:   2018-09-04 23:06
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-09-04 23:06
*************************************************************/
package shell

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
