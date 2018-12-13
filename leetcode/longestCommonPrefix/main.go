/************************************************************
** @Description: longestCommonPrefix
** @Author: george hao
** @Date:   2018-12-12 17:40
** @Last Modified by:  george hao
** @Last Modified time: 2018-12-12 17:40
*************************************************************/
package main

import "fmt"

func main() {

	strs := []string{"haohd117ddd", "haohd117123", "haohd117ddaquan", "haohd117ao"}
	//strs := []string{""}
	//strs := []string{"aca", "0ba"}
	str := longestCommonPrefix(strs)

	fmt.Println(str)

}

func longestCommonPrefix(strs []string) string {
	//求最短字符串
	size := len(strs)
	if size == 0 {
		return ""
	}
	if size == 1 {
		return strs[0]
	}

	var dp = make([][]rune, 0) //构造一个矩阵

	var minLen = len([]rune(strs[0])) //最小字符串长度
	for _, v := range strs {
		b := []byte(v)
		if minLen >= len(b) {
			minLen = len(b)
		}
		dp = append(dp, []rune(v))
	}

	if minLen == 0 {
		return ""
	}
	maxPreLen := 0

	var isEq = true
	for j := 0; j < minLen; j++ {
		if !isEq {
			break
		}
		for i := 1; i < size; i++ {
			//不相等就跳出
			if dp[i][j] != dp[i-1][j] {
				isEq = false
				break
			}
			//如果对比到最后，i=size-1
			if i == size-1 {
				maxPreLen++
				break
			}
		}
	}
	return string(dp[0][:maxPreLen])
}
