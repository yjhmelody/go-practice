package main

import (
	"fmt"
	"math"
)

func main() {
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2
	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n", y)

	medals := []string{"gold", "silver", "bronze"}
	for i := len(medals) - 1; i >= 0; i-- {
		fmt.Println(medals[i])
	}

	var a int32 = 1
	var b int64 = 2
	fmt.Println(int(a) + int(b))

	o := int64(0666)
	// 重复使用第一个操作数
	fmt.Printf("%d %[1]o %#[1]o %T\n", o, o)
	var f float32 = 1 << 24
	// 最大
	fmt.Println(f == f+1)

	nan := math.NaN()
	fmt.Println(nan == nan)
	// 1+2i
	var complexX complex128 = complex(1, 2)
	fmt.Println(complexX)
	fmt.Printf("%T\n", 1+2i)

	var boolA int = 1
	if boolA == 1 {
		fmt.Println(boolA)
	}

	var str string = "hello \u4e16\u754c"
	var hello = str[:5]
	var world = str[6:]
	fmt.Println(hello, len(hello))
	fmt.Println(world, len(world))

	const html = `
	<div> 
	hello world 
	</div>
	`
	fmt.Println(html)
	fmt.Println(HasPrefix(str, hello))
	fmt.Println(HasSuffix(str, world))

	const (
		deadbeef = 0xdeadbeef
		a        = uint32(deadbeef)
		b        = float32(deadbeef)
		c        = float64(deadbeef)
		// error
		d = int32(deadbeef)
		e = float32(deadbeef)
		f = uint32(-1)
	)
}

func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}
