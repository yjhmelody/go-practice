package main

import "fmt"

func main(){
	var str1 = "abc"
	var str2 = "bcd"
	table := LCSlength(str1, str2)

	fmt.Println(table[3][3])
}

func min(x, y int)int {
	if x > y {
		return y
	}
	return x
}

func LCSlength(str1, str2 string) ([][]int){
	len1 := len(str1)
	len2 := len(str2)
	table := make([len1+1][len2+1]int, (len1+1)*(len2+1))
	
	fmt.Println(table[0][1], 3)
	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			// fmt.Println(i, j)
			if str1[i] == str2[j] {
				table[i][j] = table[i-1][j-1] + 1
			} else {
				table[i][j] = min(table[i-1][j], table[i][j-1]) 
			}
		}
	}
	return table
}
