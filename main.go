package main

import (
	"fmt"
)


var unexp = errors.New("some unexported error") // MATCH /error var unexp should have name of the form errFoo/

func main() {
	fmt.Println("Hello, playground")
}
