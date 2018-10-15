/************************************************************
** @Description: bubble_sort
** @Author: haodaquan
** @Date:   2018-08-29 22:01
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-08-29 22:01
*************************************************************/
package bubble

//冒泡排序（基础版)
//最直观的冒泡排序
func Sort(arr []int) []int {
	len := len(arr)
	if len < 2 {
		return arr
	}
	for i := 1; i < len; i++ {
		//内层循环相邻进行比较
		for j := 0; j < len-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
			//fmt.Println(arr)
		}
	}
	return arr
}

//优化版2：观察可以发现，外层循环一次，内层循环完毕，此时最大值会被排在最后，
// 外层循环n次，则最后n位上的数字也是排序好的，因此不需要排序
func Sort2(arr []int) []int {
	len := len(arr)
	if len < 2 {
		return arr
	}
	for i := 0; i < len-1; i++ {
		//内层循环相邻进行比较
		// len-i 之后的数据是排序好的，因为每次外层循环都是把大数字向后排
		for j := 0; j < len-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
			//fmt.Println(arr)
		}
	}
	return arr
}

// 优化版3 鸡尾酒排序法
// 数组中的数字本是无规律的排放，先找到最小的数字，把他放到第一位，然后找到最大的数字放到最后一位。
// 然后再找到第二小的数字放到第二位，再找到第二大的数字放到倒数第二位。以此类推，直到完成排序。
func Sort3(arr []int) []int {
	len := len(arr)
	if len < 2 {
		return arr
	}

	for i := 0; i < len/2; i++ {
		for j := i; j < len-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
		for j := len - i - 1; j > i; j-- {
			if arr[j] < arr[j-1] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}
	return arr
}

//优化4：如果内部循环没有发生交换行为，说明内部循环数组是排好序的，则跳出外部循环
func Sort4(arr []int) []int {
	len := len(arr)
	if len < 2 {
		return arr
	}

	for i := 0; i < len-1; i++ {
		flag := true
		for j := 0; j < len-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
				flag = false
			}
		}
		if flag {
			break
		}
	}
	return arr
}

// 优化5：如果内部循环没有发生交换行为，
// 说明内部循环数组是排好序的，则跳出外部循环,提前结束
func Sort5(arr []int) []int {
	len := len(arr)
	if len < 2 {
		return arr
	}

	k := len - 1
	lastIndex := 0

	for i := 0; i < len-1; i++ {
		flag := true
		for j := 0; j < k; j++ {
			if arr[j] > arr[j+1] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
				flag = false
				lastIndex = j
			}
		}
		if flag {
			break
		}

		k = lastIndex
	}
	return arr
}
