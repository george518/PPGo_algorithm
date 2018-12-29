/************************************************************
** @Description: _36_singleNumber
** @Author: george hao
** @Date:   2018-12-27 16:14
** @Last Modified by:  george hao
** @Last Modified time: 2018-12-27 16:14
*************************************************************/
package main

func main() {

}

func singleNumber(nums []int) int {
	l := len(nums)
	m := 0
	for i := 0; i < l; i++ {
		m ^= nums[i]
	}
	return m
}
