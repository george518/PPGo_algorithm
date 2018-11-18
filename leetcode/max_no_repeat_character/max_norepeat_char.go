/************************************************************
** @Description: max_no_repeat_character
** @Author: haodaquan
** @Date:   2018-11-17 23:08
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-11-17 23:08
*************************************************************/
package main

func lengthOfLongestSubstring(s string) int {
	b := []byte(s)
	len := len(b)
	if len == 0 {
		return 0
	}
	startIndex := 0
	endIndex := 0
	maxLen := 1

	for i := 1; i < len; i++ {
		flag := 0
		for j := startIndex; j <= endIndex; j++ {
			if b[i] == b[j] {
				startIndex = j + 1
				endIndex += 1
				if (endIndex - startIndex + 1) > maxLen {
					maxLen = endIndex - startIndex + 1
				}
				break
			} else {
				flag = 1
			}
		}
		if flag == 1 {
			endIndex = i
			flag = 0
			if (endIndex - startIndex + 1) > maxLen {
				maxLen = endIndex - startIndex + 1
			}
		}
	}
	return maxLen
}

func main() {
	s := "abcaabcdefghijklmnsi"
	n := lengthOfLongestSubstring(s)
	println(n)

	s = "abc"
	n = lengthOfLongestSubstring(s)
	println(n)

	s = "a"
	n = lengthOfLongestSubstring(s)
	println(n)

	s = "acc"
	n = lengthOfLongestSubstring(s)
	println(n)

	s = "abcdefghijklmnopq"
	n = lengthOfLongestSubstring(s)
	println(n)
}
