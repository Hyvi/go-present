// https://stackoverflow.com/questions/28432658/does-go-garbage-collect-parts-of-slices/28432812#28432812

// 验证问题
// 1. the unreachable part of the slice won't be garbage-collected,
// 2. only reachable element will be copied when  append and slice is already full.
package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}
