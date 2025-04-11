// https://github.com/fatih/vim-go-tutorial#guru for Study, 虽然已经过时的文档了，但是对理解基础概念非常有用

// add option to choose between guru and gopls for :GoReferrers
//   https://github.com/fatih/vim-go/pull/2566/files

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	h := make(handler)
	go counter(h)
	if err := http.ListenAndServe(":8000", h); err != nil {
		log.Print(err)
	}
}

func counter(ch chan<- int) {
	for n := 0; ; n++ {
		ch <- n
	}
}

type handler chan int

func (h handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "text/plain")
	fmt.Fprintf(w, "%s: you are visitor #%d", req.URL, <-h)
}
