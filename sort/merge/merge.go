/************************************************************
** @Description: merge
** @Author: haodaquan
** @Date:   2018-09-02 18:16
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-09-02 18:16
*************************************************************/
package merge

import (
	"github.com/george518/PPGo_algorithm/sort/insertion"
)

func merge(arr *[]int, l, mid, r int64) {
	//计算数组长度
	len := r - l + 1
	//重新开辟一个数组空间，长度和arr一样，并做赋值
	newArr := make([]int, len)
	for i := l; i <= r; i++ {
		newArr[i-l] = (*arr)[i]
	}

	//设定 i 为左半部分的起始索引位置
	//j为右半部分的起始索引位置
	var i = l
	var j = mid + 1

	//迭代，从l到r;
	for k := l; k <= r; k++ {

		//i-l表示左半部分没有处理过的下标
		//j-l表示有半部分没有处理过的下标
		switch {
		case i > mid:
			//说明左半部分已经处理完毕，处理右边剩余的
			(*arr)[k] = newArr[j-l]
			j++
		case j > r:
			//说明右半部分已经处理完毕，处理左边剩余的
			(*arr)[k] = newArr[i-l]
			i++
		case newArr[i-l] > newArr[j-l]:
			//比较两边部分的领头元素,如果右半部分领头元素小，则赋值给k索引
			(*arr)[k] = newArr[j-l]
			j++
		default:
			(*arr)[k] = newArr[j-l]
			j++
		}
	}

}

func sort(arr *[]int, l, r int64) {
	//如果数组元素小于1个，直接return
	if l >= r {
		return
	}

	//优化点2：如果元素小于15，使用插入排序法效率更高一些
	if r-l <= 15 {
		insertion.SortApart(arr, l, r)
		return
	}

	//取数组中部的一个点作为分组，分别是arr[l,mid]和arr[mid+1,r]
	mid := (l + r) / 2
	//左边部分递归
	sort(arr, l, mid)
	//右边部分递归
	sort(arr, mid+1, r)

	//因为经过上一步的递归，左边和右边部分的数组是顺序的，
	// 因此如果左边最后一个元素arr[mid],是左边数组最大的一个，而arr[mid+1]是右边数组最小的一个，
	// arr[mid]<=arr[mid+1] 说明数组已经排好序，不需要merge了

	if (*arr)[mid] > (*arr)[mid+1] {
		merge(arr, l, mid, r)
	}

}

func Sort(arr []int) []int {
	len := int64(len(arr))
	sort(&arr, 0, len-1)
	return arr
}
