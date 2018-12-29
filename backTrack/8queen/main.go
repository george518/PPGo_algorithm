/************************************************************
** @Description: _queen
** @Author: george hao
** @Date:   2018-12-29 16:10
** @Last Modified by:  george hao
** @Last Modified time: 2018-12-29 16:10
*************************************************************/
package main

import "fmt"

var Result []int

func main() {
	Result = make([]int, 8)

	cal8Queens(0)

	fmt.Println(Result)

	PrintQueens()
}

func cal8Queens(row int) {

	if row == 8 {
		return
	}
	for column := 0; column < 8; column++ {
		if isOk(row, column) {
			Result[row] = column
			cal8Queens(row + 1)
		}
	}
}

func isOk(row, column int) bool {
	leftUp := column - 1
	rightUp := column + 1

	for i := row - 1; i >= 0; i-- { //当前行向上考察每一行

		if Result[i] == column { //第i行是否有皇后？
			if row == 5 {
				fmt.Println(i, Result[i], column, leftUp)
			}
			return false
		}

		if leftUp >= 0 { //左上对角线有皇后吗？
			if Result[i] == leftUp {

				return false
			}

		}

		if rightUp < 8 { //右上角有皇后吗？
			if Result[i] == rightUp {
				return false
			}
		}

		leftUp--
		rightUp++
	}

	return true
}

func PrintQueens() {
	for row := 0; row < 8; row++ {
		for column := 0; column < 8; column++ {
			if Result[row] == column {
				fmt.Print(" 1 ")
			} else {
				fmt.Print(" 0 ")
			}
		}

		fmt.Println()
	}
}
