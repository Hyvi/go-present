package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	server := &http.Server{
		Addr:         ":8080",
		WriteTimeout: 2 * time.Second,
	}
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	mux := http.NewServeMux()
	mux.Handle("/", &dussHandler{})
	server.Handler = mux

	go func() {
		<-quit
		if err := server.Close(); err != nil {
			log.Fatal("Close Server: ", err)
		}
	}()

	log.Println("start http server ...")
	err := server.ListenAndServe()
	if err != nil {
		if err == http.ErrServerClosed {
			log.Println("Server closed under request")
		} else {
			log.Fatal("Server cloesed unexpected")
		}
	}
	log.Println("Server exit")
}

type dussHandler struct{}

func (*dussHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("duss, version 3"))
}
