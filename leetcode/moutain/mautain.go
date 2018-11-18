/************************************************************
** @Description: moutain
** @Author: haodaquan
** @Date:   2018-11-18 16:30
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-11-18 16:30
*************************************************************/
package main

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 4, 3, 1}
	//a := []int{0, 9, 0}
	i := peakIndexInMountainArray(a)
	println(i)
}

func peakIndexInMountainArray(A []int) int {
	len := len(A)

	i := len / 2
	//fmt.Println(i)
	if A[i] > A[i-1] && A[i] > A[i+1] {
		return i
	}

	//上坡
	if A[i] > A[i-1] && A[i] < A[i+1] {
		//i = (i+len)/2 + 1
		for j := i + 1; j < len-1; j++ {
			if A[j] > A[j-1] && A[j] > A[j+1] {
				return j
			}
		}
	}

	//下坡
	if A[i] < A[i-1] && A[i] > A[i+1] {
		for j := 1; j < i; j++ {
			if A[j] > A[j-1] && A[j] > A[j+1] {
				return j
			}
		}
	}

	return 0
}
