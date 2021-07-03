// https://stackoverflow.com/questions/28432658/does-go-garbage-collect-parts-of-slices/28432812#28432812

// 验证问题
// 1. the unreachable part of the slice won't be garbage-collected,
// 2. only reachable element will be copied when  append and slice is already full.
package main

import (
	"fmt"
	"reflect"
	"runtime"
	"unsafe"
)

func PopFront(q *[]string) string {
	var r string
	r, *q = (*q)[0], (*q)[1:]
	return r
}

func PopEnd(q *[]string) string {
	var e string
	e, *q = (*q)[len(*q)-1], (*q)[:len(*q)-1]
	return e
}

func PushBack(q *[]string, a string) {
	*q = append(*q, a)
}
func main() {
	fmt.Println("Hello World")
	q := make([]string, 0, 4)

	PushBack(&q, "A")
	PushBack(&q, "B")

	PopEnd(&q)

	fmt.Println(q)
	fmt.Printf("cap:%d, len:%d \n", cap(q), len(q))
	// runtime error: index out of range [1] with length 1
	// fmt.Println(q[1])

    // https://golang.org/pkg/runtime/#GC
	runtime.GC() // GC runs a garbage collection and blocks the caller until the garbage collection is complete. It may also block the entire program.

	// https://stackoverflow.com/questions/36706843/how-to-get-the-underlying-array-of-a-slice-in-go
	// How to get the underlying array of a slice in Go
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&q))
	data := *(*[4]string)(unsafe.Pointer(hdr.Data))

	fmt.Println(data) // "B" 没有被GC

    /***
     * 复制的slice的时候，只复制了slice对应的数据，并没有复制整个底层的数组
     */
    dst_q := make([]string, 3, 4)
    num := copy(dst_q, q)
    dst_hdr := (*reflect.SliceHeader)(unsafe.Pointer(&dst_q))
	dst_data := *(*[4]string)(unsafe.Pointer(dst_hdr.Data))


    fmt.Println(num) // 复制的个数
    fmt.Println(dst_data) // 底层数组并没有 "B",  说明了一切

}
