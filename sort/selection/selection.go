/************************************************************
** @Description: selection
** @Author: haodaquan
** @Date:   2018-08-30 22:18
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-08-30 22:18
*************************************************************/
package selection

func Sort(arr []int) []int {
	len := len(arr)
	if len <= 1 {
		return arr
	}

	for i := 0; i < len; i++ {
		for j := i + 1; j < len; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}

	return arr
}
