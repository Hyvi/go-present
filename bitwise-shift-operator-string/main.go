package main

import (
	"fmt"
)

const s = "Go101.org"

var a byte = 1 << len(s) / 128 
var b byte = 1 << len(s[:]) / 128

func main() {
	fmt.Println(a, b) // 4 0	
}
