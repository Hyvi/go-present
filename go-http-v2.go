package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &dussHandler{})

	log.Println("start http server ...")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

type dussHandler struct{}

func (*dussHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("duss, version 2"))
}
