/************************************************************
** @Description: _17_containsDuplicate
** @Author: george hao
** @Date:   2018-12-27 09:12
** @Last Modified by:  george hao
** @Last Modified time: 2018-12-27 09:12
*************************************************************/
package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 1}

	fmt.Println(containsDuplicate(nums))
}

func containsDuplicate(nums []int) bool {
	l := len(nums)

	if l < 2 {
		return false
	}

	m := make(map[int]int, l)

	for i := 0; i < l; i++ {
		if _, ok := m[nums[i]]; ok {
			return true
		}
		m[nums[i]] = 1
	}

	return false

}
