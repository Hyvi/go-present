package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter,
		r *http.Request) {
		w.Write([]byte("hello, version 1"))
	})

	log.Println("start http server ...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
