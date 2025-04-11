package main
import (
	"fmt"
)

type T struct {
	s string	
}
 
// https://gronskiy.com/posts/2020-04-golang-pointer-vs-value-methods/#fn:1
// type T struct { },  the copying will not happen and all the four addresses will be the same! 
// Pointer type receiver
func (receiver *T) pointerMethod() {
    fmt.Printf("Pointer method called on \t%#v with address %p\n\n", *receiver, receiver)
}

// Value type receiver
func (receiver T) valueMethod() {
    fmt.Printf("Value method called on \t%#v with address %p\n\n", receiver, &receiver)
}
func main() {
	var (
		val     T  = T{s: "111"}
		pointer *T = &val
	)

	fmt.Printf("Value created \t%#v with address %p\n", val, &val)
	fmt.Printf("Pointer created on \t%#v with address %p\n", *pointer, pointer)

	val.valueMethod()
	pointer.pointerMethod()	

	// origin article wrong output

	pointer.valueMethod()   // Implicitly converted to: (*pointer).valueMethod()
	val.pointerMethod()   // Implicitly converted to: (&value).pointerMethod()
}
