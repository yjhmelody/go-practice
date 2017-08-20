package main

import "fmt"

func main() {
	var arr = []int{2, 8, 7, 1, 3, 5, 6, 4}
	for _, v := range arr {
		fmt.Print(v, " ")
	}

	ok := Quicksort(arr, 0, 7)
	fmt.Println(ok)

	for _, v := range arr {
		fmt.Print(v, " ")
	}

}

func Quicksort(arr []int, left, right int) (ok bool) {
	if len(arr) <= right || left < 0 {
		return false
	}
	quicksort(arr, left, right)
	return true
}

func quicksort(arr []int, left, right int) {
	if left < right {
		mid := partition(arr, left, right)
		quicksort(arr, left, mid-1)
		quicksort(arr, mid+1, right)
	}
}

func partition(arr []int, left, right int) int {
	var x = arr[right]
	var i = left - 1
	for j := left; j < right; j++ {
		if arr[j] < x {
			i++
			var tmp = arr[j]
			arr[j] = arr[i]
			arr[i] = tmp
		}
	}
	var tmp = arr[i+1]
	arr[i+1] = arr[right]
	arr[right] = tmp
	// 返回分割点
	return i + 1
}
