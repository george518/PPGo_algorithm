/************************************************************
** @Description: removeDuplicates
** @Author: george hao
** @Date:   2018-12-24 17:51
** @Last Modified by:  george hao
** @Last Modified time: 2018-12-24 17:51
*************************************************************/
package main

import (
	"fmt"
)

func main() {

	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	n := removeDuplicates(nums)
	fmt.Println(n)

}

func removeDuplicates(nums []int) int {

	l := len(nums)
	if l <= 1 {
		return l
	}

	n := 0
	for i := 0; i < l; i++ {
		if nums[i] != nums[n] {
			n++
			nums[n] = nums[i]
		}
	}
	return n + 1

}
