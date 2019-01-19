package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", dussHandlerFunc)

	log.Println("start http server ...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func dussHandlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("duss, version 1"))
}
