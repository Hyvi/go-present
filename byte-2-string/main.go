package byte2string

import (
	"fmt"
)

func B2S() {
	var bys [20]byte

	bys[0] = 'h'
	bys[1] = 'e'
	bys[2] = 'i'

	if string(bys[:]) == "hei" {
		fmt.Println("[20]byte('h','e','i') === \"hei\"")
	}

	fmt.Println("1111111")
}
