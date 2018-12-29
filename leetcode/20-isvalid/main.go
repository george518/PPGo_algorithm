/************************************************************
** @Description: _0_isvalid
** @Author: george hao
** @Date:   2018-12-24 10:29
** @Last Modified by:  george hao
** @Last Modified time: 2018-12-24 10:29
*************************************************************/
package main

import (
	"fmt"
	"strings"
)

//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
//
//有效字符串需满足：
//
//1、左括号必须用相同类型的右括号闭合。
//2、左括号必须以正确的顺序闭合。

//注意空字符串可被认为是有效字符串。

func main() {
	s := "(])"

	fmt.Println(isValid(s))

}

func isValid(s string) bool {

	if s == "" {
		return true
	}
	//用于存放括号的对应关系
	strMap := make(map[string]string, 0)
	strMap = map[string]string{"(": ")", "[": "]", "{": "}"}
	leftType := "([{"
	rightType := ")]}"

	l := len(s)
	//用于存放括号的栈 //slice
	strStack := make([]string, 0)
	sBytes := []byte(s)

	for i := 0; i < l; i++ {
		if strings.Contains(leftType, string(sBytes[i])) {
			strStack = append(strStack, string(sBytes[i]))
		} else if strings.Contains(rightType, string(sBytes[i])) {

			//最后一个
			lastIndex := len(strStack) - 1
			if lastIndex < 0 {
				return false
			}

			//如果相等，删除
			if strMap[strStack[lastIndex]] == string(sBytes[i]) {
				//fmt.Println(strStack)
				strStack = strStack[:lastIndex]
			} else {
				return false
			}
		} else {
			return false
		}
	}

	if len(strStack) == 0 {
		return true
	}

	return false
}
