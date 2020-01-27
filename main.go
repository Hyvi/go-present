package go-present

import (
	"fmt"
	"errors"
)

var unexp = errors.New("some unexported error") // MATCH /error var unexp should have name of the form errFoo/

func main() {
	fmt.Printf("Hello, playground: %v", unexp)
}
