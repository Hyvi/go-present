package main

import (
	"log"
	"net/http"

	byte2string "github.com/Hyvi/go-present/byte-2-string"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter,
		r *http.Request) {
		w.Write([]byte("hello, version 1"))
	})

	byte2string.B2S()
	log.Println("start http server ...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
