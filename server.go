package main

import (
	"fmt"
	"net"
)

func main() {
	var conn net.Listener
	var err error
	conn, err = net.Listen("tcp", "127.0.0.1:3000")
	if err != nil {
		fmt.Println(err)
	}
	for {
		ln, err := conn.Accept()
		if err != nil {
			fmt.Println(err)
		}
		go func(ln net.Conn) {
			buf := make([]byte, 1000)
			l, err := ln.Read(buf)
			if err != nil {
				fmt.Println(err)
			}
			l, err = ln.Write(buf[:l])
			fmt.Println(string(buf[:l]))
			if err != nil {
				fmt.Println(err)
			}
			ln.Close()
		}(ln)
	}
}
