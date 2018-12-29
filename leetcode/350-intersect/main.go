/************************************************************
** @Description: _50_intersect
** @Author: george hao
** @Date:   2018-12-27 16:28
** @Last Modified by:  george hao
** @Last Modified time: 2018-12-27 16:28
*************************************************************/
package main

import "fmt"

func main() {
	n1 := []int{4, 9, 5}
	n2 := []int{9, 4, 9, 8, 4}

	fmt.Println(intersect(n1, n2))
}

func intersect(nums1 []int, nums2 []int) []int {

	nums := make(map[int]int, 0)
	for _, v := range nums1 {
		nums[v]++
	}

	num := make([]int, 0)
	for _, v := range nums2 {
		if _, ok := nums[v]; ok && nums[v] > 0 {
			num = append(num, v)
			nums[v]--
		}
	}
	return num
}
