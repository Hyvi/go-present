package main

import "fmt"

// https://stackoverflow.com/questions/19310700/what-is-a-rune#:~:text=(-,Got%20a%20feeling,-that%20the%20above

func main() {
	s := "hello你好"

	fmt.Printf("len of \"%s\": %d\n", s, len(s)) // len is 11

	// byte is an alias type of uint8, in go

	fmt.Printf("type of s[0]: %T, value of s[0]: %v\n", s[0], s[0])

	// When converting the string to []rune, it found 7 utf8 chars, thus 7 runes.
	rs := []rune(s)

	fmt.Printf("len of \"%v\": %d; type: %T\n", rs, len(rs), rs)
}
