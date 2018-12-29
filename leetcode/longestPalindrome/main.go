/************************************************************
** @Description: longestPalindrome
** @Author: george hao
** @Date:   2018-12-18 10:40
** @Last Modified by:  george hao
** @Last Modified time: 2018-12-18 10:40
*************************************************************/
package main

import (
	"fmt"
	"sync"
)

//给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

//示例 1：
//输入: "babad"
//输出: "bab"
//注意: "aba" 也是一个有效答案。

func main() {
	s := "cbbd"

	s2 := longestPalindrome(s)
	fmt.Println(s2)
}

var maxLength int
var longest []byte

var lock sync.Mutex  //互斥锁
var w sync.WaitGroup //等待子线程退出

//多线程版
func longestPalindrome1(s string) string {

	len := len(s)
	if len < 2 {
		return s
	}
	b := []byte(s)

	maxLength = 0
	longest = b[:1]

	for i := 0; i < len; i++ {

		//odd
		w.Add(1)
		go Palindrome(b, len, i, 0)

		//even
		w.Add(1)
		go Palindrome(b, len, i, 1)

	}
	w.Wait()
	return string(longest)
}

func Palindrome1(b []byte, len, i, offset int) {
	left := i
	right := i + offset

	for left >= 0 && right < len && b[left] == b[right] {
		left--
		right++
	}

	lock.Lock()
	if (right - left) >= maxLength {
		maxLength = right - left
		longest = b[left+1 : right]
	}
	lock.Unlock()
	w.Done() //相当于标记关闭一个子线程
}

//正常班
func longestPalindrome2(s string) string {

	len := len(s)
	if len < 2 {
		return s
	}
	b := []byte(s)

	maxLength = 0
	longest = b[:1]

	for i := 0; i < len; i++ {

		//odd
		Palindrome(b, len, i, 0)

		//even
		Palindrome(b, len, i, 1)

	}
	return string(longest)
}

func Palindrome2(b []byte, len, i, offset int) {
	left := i
	right := i + offset

	for left >= 0 && right < len && b[left] == b[right] {
		left--
		right++
	}

	if (right - left) >= maxLength {
		maxLength = right - left
		longest = b[left+1 : right]
	}

}

//正常班
func longestPalindrome(s string) string {

	len := len(s)
	if len < 2 {
		return s
	}
	b := []byte(s)

	longest = b[:1]

	for i := 0; i < len; i++ {
		//odd
		Palindrome(b, len, i, 0)

		//even
		Palindrome(b, len, i, 1)
	}
	return string(longest)
}

func Palindrome(b []byte, length, i, offset int) {
	left := i
	right := i + offset

	for left >= 0 && right < length && b[left] == b[right] {
		left--
		right++
	}

	if len(b[left+1:right]) >= len(longest) {
		longest = b[left+1 : right]
	}
}
