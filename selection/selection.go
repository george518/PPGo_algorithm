/************************************************************
** @Description: selection
** @Author: haodaquan
** @Date:   2018-06-02 00:44
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-06-02 00:44
*************************************************************/
package selection

import (
	. "github.com/george518/PPGo_algorithm/lib"
)

type Selection struct {
	FuncName string
	Se       *SortEntity
}

func (sl *Selection) Sort() {
	if sl.Se.Len <= 1 {
		return
	}
	for i := 0; i < sl.Se.Len; i++ {
		for j := i + 1; j < sl.Se.Len-1; j++ {
			if sl.Se.Arr[i] > sl.Se.Arr[j] {
				sl.Se.Arr[i], sl.Se.Arr[j] = sl.Se.Arr[j], sl.Se.Arr[i]
			}
		}
	}
	return
}

func (sl *Selection) GetFuncName() string {
	return sl.FuncName
}

func SlectionTest(n, m int) {
	se := &SortEntity{
		Arr:     RandArray(n, 1, n),
		Len:     n,
		RunTime: 0,
	}

	s := Selection{
		FuncName: "Selection Sort",
		Se:       se,
	}

	var Is ISort
	Is = &s
	RunSort(Is)
	s.Se.Print(0, m)
}
