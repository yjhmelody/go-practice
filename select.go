package main

import (
	"fmt"
)


func main(){
	intChan := make(chan int, 5)	
	for i := 0; i < 5; i++{
		select{
			case intChan <- 1:
			case intChan <- 2:
			case intChan <- 3:
		}
	}
	for i := 0;i < 5; i++ {
		fmt.Println(<-intChan)
	}
}