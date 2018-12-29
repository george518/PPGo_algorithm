/************************************************************
** @Description: _89_rotate
** @Author: george hao
** @Date:   2018-12-25 17:37
** @Last Modified by:  george hao
** @Last Modified time: 2018-12-25 17:37
*************************************************************/
package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	fmt.Println(nums)
}

func rotate(nums []int, k int) {

	if k == 0 {
		return
	}
	l := len(nums)
	if l == 1 {
		return
	}

	n := k % l
	for i := l - n; i < l; i++ {
		step := 0
		for j := i; j > 0; j-- {
			if step < l-n {
				nums[j], nums[j-1] = nums[j-1], nums[j]
				step++
			}
		}
	}
}
