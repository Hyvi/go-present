package main

import ( 
	"os"
	"io"
	"bytes"
	"fmt"
)

type ValueMethodCaller interface {
    valueMethod()
}

type PointerMethodCaller interface {
    pointerMethod()
}


// No indication that struct `T` intends to implement `ValueMethodCaller` interface
type T struct {
    // ...
	s string
}

func (receiver T) valueMethod() {
    fmt.Printf("Value method called on \t%#v with address %p\n\n", receiver, &receiver)
}

func (receiver *T) pointerMethod() {
    fmt.Printf("Pointer method called on \t%#v with address %p\n\n", *receiver, receiver)
}



func callValueMethodOnInterface(v ValueMethodCaller) {
    v.valueMethod()
}

func callPointerMethodOnInterface(p PointerMethodCaller) {
    p.pointerMethod()
}

func main() {
	var (
		val     T  = T{s: "hello interface"}
		pointer *T = &val
	)

	callValueMethodOnInterface(val)
	callPointerMethodOnInterface(pointer)	

	callValueMethodOnInterface(pointer)
	// can not compile
	// callPointerMethodOnInterface(val)  


	// https://go.dev/doc/faq#:~:text=var%20buf%20bytes.-,Buffer,-io.Copy(buf 
	// 为什么这里一定要用指针类型
	var buf *bytes.Buffer 
	_, _ = io.Copy(buf, os.Stdin)

}
