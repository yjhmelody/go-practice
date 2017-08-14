package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){
	var mutex sync.Mutex
	mutex.Lock()
	for i := 1; i <= 3; i++{
		go func(i int) {
			fmt.Printf("Lock the lock. g%d\n", i)
			mutex.Lock()
			fmt.Printf("The lock is locked. g%d\n", i)
		}(i)
		
		time.Sleep(time.Second)
		fmt.Println("Unlock the lock.")
		mutex.Unlock()
		fmt.Println("The lock is unlocked.")
		time.Sleep(time.Second)
	}
}