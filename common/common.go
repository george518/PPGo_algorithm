/************************************************************
** @Description: common
** @Author: haodaquan
** @Date:   2018-08-29 22:17
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-08-29 22:17
*************************************************************/
package common

import (
	"math/rand"
	"time"
)

//生成随机数组
func RandArray(n, min, max int) []int {
	arr := make([]int, n)
	if n <= 0 || min > max {
		return arr
	}
	rand.Seed(time.Now().UnixNano()) //利用当前时间的UNIX时间戳初始化rand包
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(max)%(max-min+1) + min
	}
	return arr
}

//检查数据是否是否从小到大排列
func CheckArray(arr *[]int) bool {
	var pre int
	for k, v := range *arr {
		if k == 0 {
			pre = v
			continue
		}
		if pre > v {
			return false
		}
		pre = v
	}
	return true
}
