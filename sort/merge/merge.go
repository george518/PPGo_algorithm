/************************************************************
** @Description: merge
** @Author: haodaquan
** @Date:   2018-09-02 18:16
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-09-02 18:16
*************************************************************/
package merge

import "github.com/george518/PPGo_algorithm/sort/insertion"

func merge(arr *[]int, l, mid, r int64) {
	len := r - l + 1
	newArr := make([]int, len)

	for i := l; i <= r; i++ {
		newArr[i-l] = (*arr)[i]
	}

	var i = l
	var j = mid + 1

	for k := l; k <= r; k++ {
		if i > mid {
			(*arr)[k] = newArr[j-l]
			j++
		} else if j > r {
			(*arr)[k] = newArr[i-l]
			i++
		} else if newArr[i-l] < newArr[j-l] {
			(*arr)[k] = newArr[i-l]
		} else {
			(*arr)[k] = newArr[j-l]
			i++
		}
	}

}

func sort(arr *[]int, l, r int64) {
	if l >= r {
		return
	}

	if r-l <= 15 {
		insertion.SortApart(arr, l, r)
		return
	}
	mid := (l + r) / 2
	sort(arr, l, mid)
	sort(arr, mid+1, r)
	if (*arr)[mid] > (*arr)[mid+1] { //优化点1
		merge(arr, l, mid, r)
	}

}

func Sort(arr []int) []int {
	len := int64(len(arr))
	sort(&arr, 0, len-1)
	return arr
}
