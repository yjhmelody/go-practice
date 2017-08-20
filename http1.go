package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	db := database{
		"shoes": 50,
		"socks": 5,
	}
	log.Fatal(http.ListenAndServe("localhost:8080", db))
}

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("%2.f", d)
}

type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for k, v := range db {
		fmt.Fprintf(w, "%s: %s\n", k, v)
	}
}
