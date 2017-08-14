package main

import "fmt"

func main(){
	arr := []int{2, 5, 3, 0, 2, 3, 0, 3}
	arr2 := make([]int, 8)
	for _, v := range arr{
		fmt.Print(v, " ")
	}
	fmt.Println()
	Countingsort(arr, arr2, 5)

	for _, v := range arr2{
		fmt.Print(v, " ")
	}
}

func Countingsort(arr, arr2 []int, max int) bool{
	tmp := make([]int, max+1)
	for _, v := range arr {
		tmp[v] += 1
		// fmt.Println(k)
	}

	// tmp确定某个值左边的值个数
	for i := 1; i <= max; i++{
		tmp[i] += tmp[i-1]
		// fmt.Print(tmp[i-1])
	}
	
	for _, v := range arr {
		//arr[k]的值对应tmp下的个数作为位置
		// 次序不变
		// 下标从0开始所以-1
		arr2[tmp[v]-1] = v
		tmp[v] -= 1
	}
	return true
}