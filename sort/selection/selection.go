/************************************************************
** @Description: selection
** @Author: haodaquan
** @Date:   2018-08-30 22:18
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-08-30 22:18
*************************************************************/
package selection

//每次查找剩余元素的最小值，交换位置
func Sort(arr []int) []int {
	len := len(arr)
	if len <= 1 {
		return arr
	}

	for i := 0; i < len; i++ {
		var minIndex = i
		for j := i + 1; j < len; j++ {
			if arr[minIndex] > arr[j] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]

	}

	return arr
}
