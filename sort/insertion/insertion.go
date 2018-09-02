/************************************************************
** @Description: insertion
** @Author: haodaquan
** @Date:   2018-09-02 11:03
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-09-02 11:03
*************************************************************/
package insertion

func Sort(arr []int) []int {
	len := len(arr)
	if len < 2 {
		return arr
	}
	for i := 1; i < len; i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}

		}
	}
	return arr
}

//改进版 交换改为赋值，大大提高算法效率
func Sort2(arr []int) []int {
	len := len(arr)
	if len < 2 {
		return arr
	}
	for i := 1; i < len; i++ {
		e := arr[i]
		var ej int
		for j := i; j > 0; j-- {
			if arr[j-1] > e {
				arr[j] = arr[j-1]
			} else {
				ej = j
				break
			}
		}
		arr[ej] = e
	}
	return arr
}
