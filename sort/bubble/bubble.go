/************************************************************
** @Description: bubble_sort
** @Author: haodaquan
** @Date:   2018-08-29 22:01
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-08-29 22:01
*************************************************************/
package bubble

func Sort(s []int) []int {
	len := len(s)
	if len < 2 {
		return s
	}
	for i := 1; i < len; i++ {
		for j := 0; j < len-i; j++ {
			if s[j] > s[j+1] {
				s[j+1], s[j] = s[j], s[j+1]
			}
			//fmt.Println(s)
		}
	}
	return s
}
