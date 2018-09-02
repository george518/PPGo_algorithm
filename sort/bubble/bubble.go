/************************************************************
** @Description: bubble_sort
** @Author: haodaquan
** @Date:   2018-08-29 22:01
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-08-29 22:01
*************************************************************/
package bubble

//冒泡排序
func Sort(arr []int) []int {
	len := len(arr)
	if len < 2 {
		return arr
	}
	for i := 1; i < len; i++ {
		for j := 0; j < len-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
			//fmt.Println(s)
		}
	}
	return arr
}
