package main

import(
	"fmt"
	"time"
	"runtime"
)


func main(){
	// 设置P的最大数量
	runtime.GOMAXPROCS(2)
	names := []string{"A", "B", "C", "D"}
	for _, name := range names {
		go func(name string){
			// 暂停该goroutine放入调度器队列
			runtime.Gosched()
			fmt.Println(name)
			fmt.Println("numgoroutinue:", runtime.NumGoroutine())
			// 当前goroutine立即退出
			runtime.Goexit()
		}(name)
	}
	time.Sleep(time.Microsecond)
}