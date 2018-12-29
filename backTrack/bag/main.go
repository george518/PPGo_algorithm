/************************************************************
** @Description: bag
** @Author: george hao
** @Date:   2018-12-29 10:48
** @Last Modified by:  george hao
** @Last Modified time: 2018-12-29 10:48
*************************************************************/
package main

import "fmt"

var maxW int //存储背包中物品总重量的最大值
func main() {
	items := []int{2, 3, 4, 9, 9, 10}
	maxWeightbags(items, 0, 0, 5, 17)

	fmt.Println(maxW)
}

//回溯算法实现

// cw 表示当前已经装进去的物品的重量和，i表示考察到哪个物品了
// w表示背包重量，items 表示每个物品的重量集合，n表示物品个数

// 假设背包可承受重量为100，物品数10，物品重量存储在数组a中，那可以这样调用函数
//f(0,0,a,10,100)

func maxWeightbags(items []int, i, cw, n, w int) {
	if cw == w || i == n { //cw=w表示装满了，i=n表示已经考察完所有的物品
		if cw > maxW {
			maxW = cw
		}
		return
	}
	maxWeightbags(items, i+1, cw, n, w)
	if cw+items[i] <= w { //超过背包可承受中量的时候，就不要再装了
		maxWeightbags(items, i+1, cw+items[i], n, w)
	}
}
