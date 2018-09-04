/************************************************************
** @Description: shell
** @Author: haodaquan
** @Date:   2018-09-04 22:59
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-09-04 22:59
*************************************************************/
package shell

func Sort(arr []int) []int {
	len := len(arr)
	if len <= 1 {
		return arr
	}
	var i, j, gap int

	for gap = len / 2; gap > 0; gap = gap / 2 {
		for i = gap; i < len; i++ {
			for j = i - gap; j >= 0 && arr[j] > arr[j+gap]; j = j - gap {
				arr[j], arr[j+gap] = arr[j+gap], arr[j]
			}
		}
	}

	return arr
}
