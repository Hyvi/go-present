package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	server := &http.Server{
		Addr:         ":8080",
		WriteTimeout: 2 * time.Second,
	}
	mux := http.NewServeMux()
	mux.Handle("/", &dussHandler{})
	server.Handler = mux
	log.Println("start http server ...")
	log.Fatal(server.ListenAndServe())
}

type dussHandler struct{}

func (*dussHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	time.Sleep(3 * time.Second)
	w.Write([]byte("duss, version 3"))
}
