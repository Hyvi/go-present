package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter,
		r *http.Request) {
		i := 2
		i += 1
		w.Write([]byte("hello, version 1"))
	})
	
	// 用于golangci github action 测试， 
	// fmt.Sprintf("%d", 1) 

	log.Println("start http server ...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
