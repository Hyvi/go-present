package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTicker(time.Nanosecond).C
	start := time.Now()
	for i := 0; i < 100; i++ {
		<-timer
	}
	// 在mac下跑并不是100ns
	fmt.Printf("vim-go, %v\n", time.Now().Sub(start))
}
